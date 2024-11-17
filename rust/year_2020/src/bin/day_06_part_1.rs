use std::collections::HashMap;

#[derive(Debug)]
struct Group(usize);

impl Group {
    fn parse(raw: &str) -> Self {
        let mut map: HashMap<char, u8> = HashMap::new();
        raw.lines().for_each(|person| {
            person.chars().for_each(|question| {
                map.entry(question).and_modify(|v| *v += 1).or_insert(1);
            });
        });

        Group(map.len())
    }
}

fn main() {
    let raw_groups = include_str!("./input_06_part_1");
    let groups = raw_groups
        .split("\n\n")
        .map(Group::parse)
        .collect::<Vec<Group>>();

    let sum = groups.iter().map(|group| group.0).sum::<usize>();

    dbg!(groups, sum);
}

#[test]
fn group_parse() {
    let group = Group::parse("abc");
    assert!(group.0 == 3);
}

#[test]
fn group_count() {
    let example = r#"
abc

a
b
c

ab
ac

a
a
a
a

b
"#
    .trim();

    let results: [usize; 5] = [3, 3, 3, 1, 1];
    let groups = example.split("\n\n").map(Group::parse);

    assert!(groups
        .enumerate()
        .all(|(index, group)| group.0 == results[index]));
}
