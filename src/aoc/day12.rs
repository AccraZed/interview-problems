use std::{collections::HashMap, error::Error, fs};

pub fn part_1() -> Result<(), Box<dyn Error>> {
    let mut cave_system: HashMap<&str, &mut Vec<&str>> = HashMap::new();

    let raw = fs::read_to_string("src/aoc/day12_input.txt")?;
    raw.split('\n').map(|path| {
        let mut p = path.split('-');
        let from = p.next().unwrap();
        let to = p.next().unwrap();
        let entry = cave_system.get(from);

        if entry == None {
            let mut list = Vec::new();
            list.push(to);
            cave_system.insert(from, &mut list);
        }
    });

    Ok(())
}
