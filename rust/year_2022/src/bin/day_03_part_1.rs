use std::collections::HashSet;

struct Rucksack<'a> {
    compartments: [Compartment<'a>; 2],
}

impl<'a> TryFrom<&'a str> for Rucksack<'a> {
    type Error = color_eyre::Report;

    fn try_from(value: &'a str) -> Result<Self, Self::Error> {
        let (first, second) = value.split_at(value.len() / 2);

        Ok(Rucksack {
            compartments: [Compartment(first), Compartment(second)],
        })
    }
}

impl Rucksack<'_> {
    fn get_priority(self) -> u32 {
        let letter = self
            .get_shared_letter()
            .expect("no shared letter between compartments");
        let res =
            Rucksack::get_alphabet_place(letter).expect("shared letter could not be recognized");
        dbg!(letter, res);
        res
    }

    fn get_alphabet_place(c: char) -> Option<u32> {
        if c.is_ascii_lowercase() {
            Some(c as u32 - 'a' as u32 + 1)
        } else if c.is_ascii_uppercase() {
            Some(c as u32 - 'A' as u32 + 27)
        } else {
            None
        }
    }

    fn get_shared_letter(self) -> Option<char> {
        let s1 = self.compartments[0].0;
        let s2 = self.compartments[1].0;
        let set1: HashSet<char> = s1.chars().collect();

        s2.chars().find(|&c| set1.contains(&c))
    }
}

struct Compartment<'a>(&'a str);

fn main() {
    //    let input = r"vJrwpWtwJgWrhcsFMMfFFhFp
    //jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
    //PmmdzqPrVvPwwTWBwg
    //wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
    //ttgJtRGJQctTZtZT
    //CrZsJsPPZsGzwwsLwLmpwMDw"
    //        .trim();
    let input = include_str!("./input_03_part_1");

    let result: u32 = input
        .lines()
        .map(Rucksack::try_from)
        .filter_map(Result::ok)
        .map(Rucksack::get_priority)
        .sum();

    println!("pickle rick {result}");
}
