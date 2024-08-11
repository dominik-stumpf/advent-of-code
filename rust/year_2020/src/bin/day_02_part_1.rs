use std::ops::RangeInclusive;

struct PasswordPolicy {
    byte: u8,
    range: RangeInclusive<usize>,
}

impl PasswordPolicy {
    fn is_valid(&self, password: &str) -> bool {
        self.range.contains(
            &password
                .as_bytes()
                .iter()
                .copied()
                .filter(|b| *b == self.byte)
                .count(),
        )
    }
}

fn parse_line(s: &str) -> (PasswordPolicy, &str) {
    let (policy, password) = {
        let mut tokens = s.split(':');
        (tokens.next().unwrap(), tokens.next().unwrap().trim())
    };
    let (range, byte) = {
        let mut tokens = policy.split(' ');
        (tokens.next().unwrap(), tokens.next().unwrap().as_bytes()[0])
    };
    let (min, max) = {
        let mut tokens = range.split('-');
        (
            tokens.next().unwrap().parse::<usize>().unwrap(),
            tokens.next().unwrap().parse::<usize>().unwrap(),
        )
    };

    (
        PasswordPolicy {
            byte,
            range: min..=max,
        },
        password,
    )
}

fn main() {
    let count = include_str!("./input_02_part_1")
        .lines()
        .map(parse_line)
        .filter(|(policy, password)| policy.is_valid(password))
        .count();

    dbg!(count);
}
