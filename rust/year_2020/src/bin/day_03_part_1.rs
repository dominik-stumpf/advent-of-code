fn main() {
    let input = include_str!("./input_03_part_1");
    let advancement = 3;
    let mut row = advancement;
    let mut tree_count = 0;

    for line in input.trim().lines().skip(1) {
        if let Some(point) = line.chars().nth(row % line.len()) {
            if point == '#' {
                tree_count += 1;
            }
        }
        row += advancement;
    }

    dbg!(tree_count);
}
