fn overlaps([r1, r2]: &[[usize; 2]; 2]) -> bool {
    if r1[0] <= r2[1] && r1[1] >= r2[0] {
        return true;
    }
    if r2[0] <= r1[1] && r2[1] >= r1[0] {
        return true;
    }
    false
}

fn main() {
    let input = include_str!("./input_04_part_1").trim();
    //    let input = r"2-4,6-8
    //2-3,4-5
    //5-7,7-9
    //2-8,3-7
    //6-6,4-6
    //2-6,4-8"
    //        .trim();

    let result = input
        .lines()
        .map(|line| {
            line.split_once(',')
                .expect("to have comma separator between ranges")
        })
        .map(|(first, second)| [first, second])
        .map(|range| {
            range
                .map(|r| r.split_once('-').expect("ranges to be separated by commas"))
                .map(|(r1, r2)| {
                    [r1, r2].map(|v| v.parse::<usize>().expect("to be valid positive integer"))
                })
        })
        //.inspect(|a| {
        //    dbg!(a, overlaps(a));
        //})
        .filter(overlaps)
        .count();

    dbg!(result);
}
