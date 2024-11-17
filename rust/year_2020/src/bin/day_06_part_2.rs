use std::collections::HashMap;

#[derive(Debug)]
struct Group(usize);

impl Group {
    fn parse(raw: &str) -> Self {
        let mut map: HashMap<u8, u8> = HashMap::new();
        let person_count = raw.lines().collect::<Vec<&str>>().len();

        raw.lines().for_each(|person| {
            person.bytes().for_each(|question| {
                map.entry(question).and_modify(|v| *v += 1).or_insert(1);
            });
        });

        Group(
            map.values()
                .filter(|&&v| v as usize == person_count)
                .count(),
        )
    }
}

fn main() {
    let raw_groups = include_str!("./input_06_part_1");
    let groups = raw_groups
        .split("\n\n")
        .map(Group::parse)
        .collect::<Vec<Group>>();

    let sum = groups.iter().map(|group| group.0).sum::<usize>();

    dbg!(sum);
}
