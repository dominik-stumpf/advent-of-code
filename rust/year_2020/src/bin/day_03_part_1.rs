// fn main() {
//     let input = include_str!("./input_03_part_1");
//     let advancement = 3;
//     let mut row = advancement;
//     let mut tree_count = 0;

//     for line in input.trim().lines().skip(1) {
//         if let Some(point) = line.chars().nth(row % line.len()) {
//             if point == '#' {
//                 tree_count += 1;
//             }
//         }
//         row += advancement;
//     }

//     dbg!(tree_count);
// }

use std::fmt;

#[derive(Debug, Clone, Copy, PartialEq)]
struct Vec2 {
    x: i64,
    y: i64,
}

impl From<(i64, i64)> for Vec2 {
    fn from((x, y): (i64, i64)) -> Self {
        Self { x, y }
    }
}

#[derive(Default, PartialEq, Clone, Copy)]
enum Tile {
    #[default]
    Open,
    Tree,
}

impl fmt::Debug for Tile {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let c = match self {
            Tile::Open => '.',
            Tile::Tree => '#',
        };
        write!(f, "{c}")
    }
}

struct Map {
    size: Vec2,
    tiles: Vec<Tile>,
}

impl Map {
    fn new(size: Vec2) -> Self {
        let num_tiles = size.x * size.y;
        Self {
            size,
            tiles: (0..num_tiles)
                .into_iter()
                .map(|_| Default::default())
                .collect(),
        }
    }

    fn set(&mut self, pos: Vec2, tile: Tile) {
        if let Some(index) = self.index(pos) {
            self.tiles[index] = tile
        }
    }

    fn get(&self, pos: Vec2) -> Tile {
        self.index(pos).map(|i| self.tiles[i]).unwrap_or_default()
    }

    /// Wrap the x coordinate so the map extends infinitely to the
    /// left and the right. Returns `None` for coordinates above 0
    /// or below `self.size.y`.
    fn normalize_pos(&self, pos: Vec2) -> Option<Vec2> {
        if pos.y < 0 || pos.y >= self.size.y {
            None
        } else {
            let x = pos.x % self.size.x;
            // wrap around for left side (negative X coordinates)
            let x = if x < 0 { self.size.x + x } else { x };
            Some((x, pos.y).into())
        }
    }

    fn index(&self, pos: Vec2) -> Option<usize> {
        self.normalize_pos(pos)
            .map(|pos| (pos.x + pos.y * self.size.x) as _)
    }

    fn parse(input: &[u8]) -> Self {
        let mut columns = 0;
        let mut rows = 1;
        for &c in input.iter() {
            if c == b'\n' {
                rows += 1;
                columns = 0;
            } else {
                columns += 1;
            }
        }

        let mut iter = input.iter().copied();
        let mut map = Self::new((columns, rows).into());
        for row in 0..map.size.y {
            for col in 0..map.size.x {
                let tile = match iter.next() {
                    Some(b'.') => Tile::Open,
                    Some(b'#') => Tile::Tree,
                    c => panic!("Expected '.' or '#', but got: {:?}", c),
                };
                map.set((col, row).into(), tile);
            }
            iter.next();
        }
        map
    }
}

impl fmt::Debug for Map {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for row in 0..self.size.y {
            for col in 0..self.size.x {
                write!(f, "{:?}", self.get((col, row).into()))?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn main() {
    let map = Map::parse(include_bytes!("./input_03_part_1"));
    let itinerary = (0..map.size.y).into_iter().map(|y| Vec2::from((y * 3, y)));
    let num_trees = itinerary.filter(|&pos| map.get(pos) == Tile::Tree).count();
    println!("We encountered {} trees", num_trees);
}

// TESTS

#[test]
fn test_normalize_pos() {
    let m = Map::new((2, 2).into());
    assert_eq!(m.normalize_pos((0, 0).into()), Some((0, 0).into()));
    assert_eq!(m.normalize_pos((1, 0).into()), Some((1, 0).into()));
    assert_eq!(m.normalize_pos((2, 0).into()), Some((0, 0).into()));
    assert_eq!(m.normalize_pos((-1, 0).into()), Some((1, 0).into()));
    assert_eq!(m.normalize_pos((-2, 0).into()), Some((0, 0).into()));
    assert_eq!(m.normalize_pos((0, -1).into()), None);
    assert_eq!(m.normalize_pos((0, 2).into()), None);
}

#[test]
fn test_vec2_from() {
    let coords: (i64, i64) = (5, 10);
    let vector = Vec2::from(coords);
    assert_eq!(vector.x, 5);
    assert_eq!(vector.y, 10);
    assert_eq!(vector, Vec2 { x: 5, y: 10 });
}

#[test]
fn test_index() {
    let m = Map::new((3, 5).into());
    assert_eq!(m.index((0, 0).into()), Some(0));
    assert_eq!(m.index((2, 0).into()), Some(2));
    assert_eq!(m.index((0, 1).into()), Some(3));
    assert_eq!(m.index((2, 1).into()), Some(5));
    assert_eq!(m.index((2, 93).into()), None);
}
