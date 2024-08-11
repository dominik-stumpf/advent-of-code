use itertools::Itertools;

fn main() -> anyhow::Result<()> {
    let result: usize = include_str!("input_01_part_1")
        .lines()
        .map(str::parse::<usize>)
        .collect::<Result<Vec<_>, _>>()?
        .into_iter()
        .combinations(3)
        .find(|c| c.iter().sum::<usize>() == 2020)
        .into_iter()
        .flatten()
        .product();

    dbg!(result);

    Ok(())
}
