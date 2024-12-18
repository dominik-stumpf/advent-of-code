use std::collections::HashSet;

fn get_alphabet_place(c: char) -> Option<u32> {
    if c.is_ascii_lowercase() {
        Some(c as u32 - 'a' as u32 + 1)
    } else if c.is_ascii_uppercase() {
        Some(c as u32 - 'A' as u32 + 27)
    } else {
        None
    }
}

fn first_common_letter(strings: &[&str]) -> Option<char> {
    let mut common_set: HashSet<char> = strings[0].chars().collect();

    for s in strings.iter().skip(1) {
        let current_set: HashSet<char> = s.chars().collect();
        common_set = common_set.intersection(&current_set).copied().collect();
    }

    common_set.into_iter().next()
}

fn main() {
    //let input = r"vJrwpWtwJgWrhcsFMMfFFhFp
    //jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
    //PmmdzqPrVvPwwTWBwg
    //wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
    //ttgJtRGJQctTZtZT
    //CrZsJsPPZsGzwwsLwLmpwMDw"
    //    .trim();
    let input = include_str!("./input_03_part_1");

    let binding = input.lines().map(str::trim).collect::<Vec<&str>>();
    let result: u32 = binding
        .chunks_exact(3)
        .map(|chunk| {
            let common_letter = first_common_letter(chunk).expect("common letter between strings");
            get_alphabet_place(common_letter).expect("letter to be in the alphabet")
        })
        .sum();

    dbg!(result);
}
