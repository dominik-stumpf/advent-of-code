fn bsp(zero: char, one: char, sequence: &[u8]) -> u8 {
    sequence
        .iter()
        .enumerate()
        .map(|(index, c)| match *c as char {
            c if c == zero => 0,
            c if c == one => 2_u8.pow(sequence.len() as u32 - index as u32 - 1 as u32),
            _ => panic!("unhandled char"),
        })
        .sum()
}

fn calculate_seat_id(sequence: &[u8]) -> u32 {
    let row = bsp('F', 'B', &sequence[0..=6]);
    let column = bsp('L', 'R', &sequence[7..]);

    row as u32 * 8 + column as u32
}

fn main() {
    let highest_id = include_str!("./input_05_part_1")
        .lines()
        .map(|line| calculate_seat_id(line.as_bytes()))
        .max();

    dbg!(highest_id);
}

#[test]
fn test_bsp() {
    let e = b"FBFBBFFRLR";
    let a = bsp('F', 'B', &e[0..=6]);
    assert!(a == 44);
}
