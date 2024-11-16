use std::ops::RangeInclusive;

use peg;

#[derive(Copy, Clone, PartialEq, Debug)]
struct Year(u64);

#[derive(Copy, Clone, PartialEq, Debug)]
enum Length {
    Cm(u64),
    In(u64),
}

#[derive(Copy, Clone, PartialEq, Debug)]
struct Color<'a>(&'a str);

#[derive(Copy, Clone, PartialEq, Debug)]
struct Id<'a>(&'a str);

#[derive(Debug, PartialEq)]
struct Passport<'a> {
    /// byr (Birth Year)
    birth_year: Year,
    /// iyr (Issue Year)
    issue_year: Year,
    /// eyr (Expiration Year)
    expiration_year: Year,
    /// hgt (Height)
    height: Length,
    /// hcl (Hair Color)
    hair_color: Color<'a>,
    /// ecl (Eye Color)
    eye_color: Color<'a>,
    /// pid (Passport ID)
    passport_id: Id<'a>,
    /// cid (Country ID)
    country_id: Option<Id<'a>>,
}

#[derive(PartialEq, Debug, Default)]
struct PassportBuilder<'a> {
    birth_year: Option<Year>,
    issue_year: Option<Year>,
    expiration_year: Option<Year>,
    height: Option<Length>,
    hair_color: Option<Color<'a>>,
    eye_color: Option<Color<'a>>,
    passport_id: Option<Id<'a>>,
    country_id: Option<Id<'a>>,
}

#[derive(thiserror::Error, Debug)]
enum Error {
    #[error("missing field: {0}")]
    MissingField(&'static str),

    #[error("could not parse {0}: {1}")]
    ParseError(String, String),
}

impl<'a> PassportBuilder<'a> {
    fn build(self) -> Result<Passport<'a>, Error> {
        macro_rules! build {
            (
                required => {
                    $($req: ident),* $(,)*
                }$(,)*
                optional => {
                    $($opt: ident),* $(,)*
                }$(,)*
            ) => {
                Ok(Passport {
                    $($req: self.$req.ok_or(Error::MissingField(stringify!($req)))?),*,
                    $($opt: self.$opt),*
                })
            }
        }

        build! {
            required => {
                birth_year,
                issue_year,
                expiration_year,
                height,
                hair_color,
                eye_color,
                passport_id,
            },
            optional => {
                country_id,
            },
        }
    }
}

impl<'a> PassportBuilder<'a> {
    fn parse(input: &'a str) -> Result<Self, Error> {
        let mut b: Self = Default::default();

        peg::parser! {
            grammar parser() for str {

                pub(crate) rule root(b: &mut PassportBuilder<'input>)
                    = (field(b) separator()*)* ![_]

                rule separator()
                    = ['\n' | ' ']

                rule field(b: &mut PassportBuilder<'input>)
                    // years
                    = byr(b) / iyr(b) / eyr(b)
                    // height
                    / hgt(b)
                    // colors
                    / hcl(b) / ecl(b)
                    // IDs
                    / pid(b) / cid(b)

                rule byr(b: &mut PassportBuilder<'input>)
                    = "byr:" year:year(1920..=2002) { b.birth_year = Some(year) }

                rule iyr(b: &mut PassportBuilder<'input>)
                    = "iyr:" year:year(2010..=2020) { b.issue_year = Some(year) }

                rule eyr(b: &mut PassportBuilder<'input>)
                    = "eyr:" year:year(2020..=2030) { b.expiration_year = Some(year) }

                rule hgt(b: &mut PassportBuilder<'input>)
                    = "hgt:" height:length() {?
                        match &height {
                            Length::Cm(v) if !(150..=193).contains(v) => {
                                Err("bad height (cm)")
                            }
                            Length::In(v) if !(59..=76).contains(v) => {
                                Err("bad height (in)")
                            }
                            _ => {
                                b.height = Some(height);
                                Ok(())
                            }
                        }
                        // b.height = Some(height)
                }

                rule pid(b: &mut PassportBuilder<'input>)
                                = "pid:" id:$(['0'..='9']*<9,9>) { b.passport_id = Some(Id(id)) }

                rule cid(b: &mut PassportBuilder<'input>)
                    = "cid:" id:$((!separator()[_])+) { b.country_id = Some(Id(id)) }

                rule hcl(b: &mut PassportBuilder<'input>)
                    = "hcl:" color:hcl0() { b.hair_color = Some(color) }

                rule hcl0() -> Color<'input>
                    = s:$("#" ['0'..='9' | 'a'..='f']*<6,6>) { Color(s) }

                rule ecl(b: &mut PassportBuilder<'input>)
                    = "ecl:" color:ecl0() { b.eye_color = Some(color) }

                rule ecl0() -> Color<'input>
                    = s:$("amb" / "blu" / "brn" / "gry" / "grn" / "hzl" / "oth") { Color(s) }

                rule year(range: RangeInclusive<u64>) -> Year
                    = num:num() {?
                        if range.contains(&num) {
                            Ok(Year(num))
                        } else {
                            Err("year out of range")
                        }
                }

                rule color() -> Color<'input>
                    = s:$((!separator()[_])*) { Color(s) }

                rule length() -> Length
                    = num:num() "cm" { Length::Cm(num) }
                    / num:num() "in" { Length::In(num) }

                rule num() -> u64
                    = s:$(['0'..='9']+) { s.parse().unwrap() }

                rule id() -> Id<'input>
                    = s:$(['0'..='9' | 'a'..='z' | '#']+) { Id(s) }
            }
        }

        parser::root(input, &mut b).map_err(|e| Error::ParseError(input.into(), e.to_string()))?;
        Ok(b)
    }
}

fn main() {
    let passports = include_str!("./input_04_part_1")
        .split("\n\n")
        .map(|raw| PassportBuilder::parse(raw).and_then(PassportBuilder::build));

    let valid_passports = passports.flatten().count();
    dbg!(valid_passports);
}

#[test]
fn test_builder() {
    assert!(PassportBuilder {
        ..Default::default()
    }
    .build()
    .is_err());
    assert!(PassportBuilder {
        birth_year: Some(Year(2014)),
        issue_year: Some(Year(2017)),
        expiration_year: Some(Year(2023)),
        height: Some(Length::Cm(195)),
        hair_color: Some(Color("#ffffff")),
        eye_color: Some(Color("#ee7812")),
        passport_id: Some(Id("00023437")),
        country_id: None,
    }
    .build()
    .is_ok());
}
