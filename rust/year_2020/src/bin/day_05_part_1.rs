use bitvec::prelude::*;

#[derive(Clone, Copy, Default, Debug, PartialEq, Eq, PartialOrd, Ord)]
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
    let mut ids: Vec<_> = include_str!("./input_05_part_1")
        .lines()
        .map(Seat::parse)
        .collect();
    ids.sort();

    let mut last_id: Option<Seat> = None;
    for id in ids {
        if let Some(last_id) = last_id {
            let gap = id.0 - last_id.0;
            if gap > 1 {
                println!("Our seat ID is {}", last_id.0 + 1);
                return;
            }
        }
        last_id = Some(id);
    }
}
