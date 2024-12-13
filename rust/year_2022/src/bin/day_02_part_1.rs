// #[derive(Debug, PartialEq, PartialOrd, Clone, Copy)]
// enum HandShape {
//     Rock,
//     Paper,
//     Scissors,
// }

// impl HandShape {
//     fn build(shape: char) -> Self {
//         match shape {
//             'A' | 'X' => Self::Rock,
//             'B' | 'Y' => Self::Paper,
//             'C' | 'Z' => Self::Scissors,
//             _ => panic!("hand shape not recognized"),
//         }
//     }
// }

// #[derive(Debug, Default)]
// struct Score(usize);

// impl Score {
//     fn update(&mut self, round: (HandShape, HandShape)) {
//         self.0 += 1 + match round {
//             shape if shape.0 == shape.1 => 3 + (shape.1 as usize),
//             shape if shape.0 < shape.1 => 6 + (shape.1 as usize),
//             shape if shape.0 > shape.1 => shape.1 as usize,
//             _ => panic!("round result could not be determined"),
//         };
//     }
// }

// fn main() {
//     let input = include_str!("./input_02_part_1");
//     // let input = r#"A Y
//     // B X
//     // C Z"#
//     //     .trim();
//     let rounds = input.lines().map(|round| {
//         let shapes: Vec<char> = round.chars().filter(|c| !c.is_whitespace()).collect();
//         (HandShape::build(shapes[0]), HandShape::build(shapes[1]))
//     });
//     let mut score = Score::default();
//     rounds.for_each(|round| {
//         dbg!(&score);
//         score.update(round);
//         dbg!(&score);
//     });

//     println!("{score:?}");
// }

use std::str::FromStr;

#[derive(Debug, Clone, Copy)]
enum Move {
    Rock,
    Paper,
    Scissors,
}

#[derive(Debug, Clone, Copy)]
struct Round {
    theirs: Move,
    ours: Move,
}

impl TryFrom<char> for Move {
    type Error = color_eyre::Report;

    fn try_from(c: char) -> Result<Self, Self::Error> {
        match c {
            'A' | 'X' => Ok(Move::Rock),
            'B' | 'Y' => Ok(Move::Paper),
            'C' | 'Z' => Ok(Move::Scissors),
            _ => Err(color_eyre::eyre::eyre!("not a valid move: {c:?}")),
        }
    }
}

impl FromStr for Round {
    type Err = color_eyre::Report;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut chars = s.chars();
        let (Some(theirs), Some(' '), Some(ours), None) =
            (chars.next(), chars.next(), chars.next(), chars.next())
        else {
            return Err(color_eyre::eyre::eyre!(
                "expected <theirs>SP<ours>EOF, got {s:?}"
            ));
        };

        Ok(Self {
            theirs: theirs.try_into()?,
            ours: ours.try_into()?,
        })
    }
}

#[derive(Debug, Clone, Copy)]
enum Outcome {
    Win,
    Draw,
    Loss,
}

impl Outcome {
    fn inherent_points(self) -> usize {
        match self {
            Outcome::Win => 6,
            Outcome::Draw => 3,
            Outcome::Loss => 0,
        }
    }
}

impl Move {
    fn beats(self, other: Move) -> bool {
        matches!(
            (self, other),
            (Self::Rock, Self::Scissors)
                | (Self::Paper, Self::Rock)
                | (Self::Scissors, Self::Paper)
        )
    }

    fn outcome(self, theirs: Move) -> Outcome {
        if self.beats(theirs) {
            Outcome::Win
        } else if theirs.beats(self) {
            Outcome::Loss
        } else {
            Outcome::Draw
        }
    }

    /// How many points do we get for picking that move?
    fn inherent_points(self) -> usize {
        match self {
            Move::Rock => 1,
            Move::Paper => 2,
            Move::Scissors => 3,
        }
    }
}

impl Round {
    fn outcome(self) -> Outcome {
        self.ours.outcome(self.theirs)
    }

    fn our_score(self) -> usize {
        self.ours.inherent_points() + self.outcome().inherent_points()
    }
}

use itertools::Itertools;

fn main() -> color_eyre::Result<()> {
    color_eyre::install()?;

    let total_score: usize = itertools::process_results(
        include_str!("./input_02_part_1")
            .lines()
            .map(Round::from_str)
            // 👇 this is provided by `Itertools`
            .map_ok(|r| r.our_score()),
        |it| it.sum(),
    )?;
    dbg!(total_score);

    Ok(())
}
