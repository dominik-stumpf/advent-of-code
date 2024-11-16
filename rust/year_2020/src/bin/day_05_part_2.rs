use bitvec::prelude::*;

#[derive(Clone, Copy, Default, Debug, PartialEq)]
struct Seat(u16);

impl Seat {
    fn parse(input: &str) -> Self {
        let mut res: Seat = Default::default();

        let bits = BitSlice::<Lsb0, _>::from_element_mut(&mut res.0);
        for (i, &b) in input.as_bytes().iter().rev().enumerate() {
            bits.set(
                i,
                match b {
                    b'F' | b'L' => false,
                    b'B' | b'R' => true,
                    _ => panic!("unexpected letter: {}", b as char),
                },
            );
        }

        res
    }
}

#[test]
fn test_seat_id() {
    assert_eq!(Seat::parse("BFFFBBFRRR"), Seat(567));
    assert_eq!(Seat::parse("FFFBBBFRRR"), Seat(119));
    assert_eq!(Seat::parse("BBFFBBFRLL"), Seat(820));
}

fn main() {
    let max_id = itertools::max(
        include_str!("./input_05_part_1")
            .lines()
            .map(Seat::parse)
            .map(|seat| seat.0),
    );
    println!("The maximum seat ID is {:?}", max_id);
}
