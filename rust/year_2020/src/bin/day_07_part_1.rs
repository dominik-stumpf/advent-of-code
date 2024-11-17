use itertools::Itertools;
use multimap::MultiMap;

/// (adjective, color), i.e. ("dark", "orange")
type BagSpec<'a> = (&'a str, &'a str);

/// K can contain V.0 of V.1
type Rules<'a> = MultiMap<BagSpec<'a>, (usize, BagSpec<'a>)>;

fn parse_rules(input: &str) -> Rules<'_> {
    let mut rules: Rules = MultiMap::default();

    peg::parser! {
        pub(crate) grammar parser() for str {
            pub(crate) rule root(r: &mut Rules<'input>)
                = (line(r) "." whitespace()*)* ![_]

            rule line(r: &mut Rules<'input>)
                = spec:bag_spec() " contain " rules:rules() {
                    if let Some(rules) = rules {
                        for rule in rules {
                            r.insert(spec, rule);
                        }
                    }
                }

            rule bag_spec() -> BagSpec<'input>
                = adjective:name() " " color:name() " bag" "s"? { (adjective, color) }

            rule rules() -> Option<Vec<(usize, BagSpec<'input>)>>
                = rules:rule1()+ { Some(rules) }
                / "no other bags" { None }

            /// Rule followed by an optional comma and space
            rule rule1() -> (usize, BagSpec<'input>)
                = r:rule0() ", "? { r }

            /// A single rule
            rule rule0() -> (usize, BagSpec<'input>)
                = quantity:number() " " spec:bag_spec() { (quantity, spec) }

            rule number() -> usize
                = e:$(['0'..='9']+) { e.parse().unwrap() }

            /// A sequence of non-whitespace characters
            rule name() -> &'input str
                = $((!whitespace()[_])*)

            /// Spaces, tabs, CR and LF
            rule whitespace()
                = [' ' | '\t' | '\r' | '\n']
        }
    }

    parser::root(input, &mut rules).unwrap();
    rules
}

fn reverse_graph<'a>(graph: &Rules<'a>) -> Rules<'a> {
    graph
        .iter_all()
        .flat_map(|(&node, neighbors)| {
            neighbors
                .iter()
                .map(move |&(quantity, neighbor)| (neighbor, (quantity, node)))
        })
        .collect()
}

fn walk_subgraph2<'iter, 'elems: 'iter>(
    graph: &'iter Rules<'elems>,
    root: &(&'iter str, &'iter str),
) -> Box<dyn Iterator<Item = (&'elems str, &'elems str)> + 'iter> {
    Box::new(
        graph
            .get_vec(root)
            .into_iter()
            .flatten()
            .flat_map(move |&(_, neighbor)| {
                std::iter::once(neighbor).chain(walk_subgraph2(graph, &neighbor))
            }),
    )
}

fn main() {
    let rules = parse_rules(include_str!("./input_07_part_1"));
    let rev_rules = reverse_graph(&rules);

    let needle = ("shiny", "gold");
    let answer = walk_subgraph2(&rev_rules, &needle).unique().count();
    println!("{answer} colors can contain {needle:?} bags");
}
