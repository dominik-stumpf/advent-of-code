use std::ops::Index;

use itertools::Itertools;

fn main() {
    //    let input = r"
    //    [D]
    //[N] [C]
    //[Z] [M] [P]
    // 1   2   3
    //
    //move 1 from 2 to 1
    //move 3 from 1 to 3
    //move 2 from 2 to 1
    //move 1 from 1 to 2
    //"
    //    .trim_matches('\n');
    let input = include_str!("./input_05_part_1").trim_matches('\n');
    let (arrangement, instructions) = input
        .split("\n\n")
        .collect_tuple()
        .expect("to separate arrangement and instruction");

    let arrangement_length = (arrangement.split('\n').next().unwrap().len() + 1) / 4;
    let mut result: Vec<Vec<char>> = vec![vec![]; arrangement_length];

    arrangement
        .lines()
        //.rev()
        //.skip(1)
        .take_while(|line| line.chars().nth(1).unwrap() != '1')
        .for_each(|line| {
            line.chars()
                .skip(1)
                .step_by(4)
                .enumerate()
                .for_each(|(i, char)| {
                    if char != ' ' {
                        result[i].push(char);
                    }
                });
        });

    instructions.lines().for_each(|line| {
        dbg!(&result);
        let elements = line.split_whitespace().skip(1).step_by(2);
        let (mov, from, to) = elements
            .map(|i| {
                i.parse::<usize>()
                    .expect("to have nonnegative integer instructions")
            })
            .collect_tuple()
            .expect("instruction to have elements");

        //let slice = result[from - 1].drain(mov..);
        //result[to - 1].extend(slice);

        let mut temp = result[from - 1].clone();
        //result[to - 1].extend(temp.drain(..mov).rev());
        result[to - 1].splice(..0, temp.drain(..mov).rev());
        result[from - 1] = temp;

        dbg!(mov, from, to);
    });

    let sequence = result.iter().map(|arr| arr[0]).join("");

    dbg!(
        arrangement,
        instructions,
        arrangement_length,
        &result,
        sequence
    );
}
