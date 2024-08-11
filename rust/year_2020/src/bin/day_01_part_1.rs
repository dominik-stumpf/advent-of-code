use itertools::Itertools;

fn main() -> anyhow::Result<()> {
    let (a, b, c) = include_str!("input_01_part_1")
        .lines()
        .map(str::parse::<usize>)
        .collect::<Result<Vec<_>, _>>()?
        .into_iter()
        .tuple_combinations()
        .find(|(a, b, c)| a + b == 2020)
        .unwrap();

    dbg!(a + b);
    dbg!(a * b);

    Ok(())
}
