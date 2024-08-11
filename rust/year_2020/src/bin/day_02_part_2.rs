struct PasswordPolicy {
    byte: u8,
    positions: [usize; 2],
}

impl PasswordPolicy {
    fn is_valid(&self, password: &str) -> bool {
        self.positions
            .iter()
            .copied()
            .filter(|&pos| password.as_bytes()[pos - 1] == self.byte)
            .count()
            == 1
    }
}

fn parse_line(s: &str) -> (PasswordPolicy, &str) {
    let (policy, password) = {
        let mut tokens = s.split(':');
        (tokens.next().unwrap(), tokens.next().unwrap().trim())
    };
    let (indices, byte) = {
        let mut tokens = policy.split(' ');
        (tokens.next().unwrap(), tokens.next().unwrap().as_bytes()[0])
    };
    let (first, second) = {
        let mut tokens = indices.split('-');
        (
            tokens.next().unwrap().parse::<usize>().unwrap(),
            tokens.next().unwrap().parse::<usize>().unwrap(),
        )
    };

    (
        PasswordPolicy {
            byte,
            positions: [first, second],
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
