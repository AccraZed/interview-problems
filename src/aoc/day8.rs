use std::{error::Error, fs};

pub fn part_1() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day8_input.txt")?;
    let lines = raw.split('\n');

    let mut count = 0;
    for line in lines {
        let mut data = line.split(" | ");
        let legend = data
            .next()
            .unwrap()
            .split(' ')
            .map(|digit| {
                let mut chars = digit.chars().collect::<Vec<char>>();
                chars.sort();
                return String::from_iter(chars);
            })
            .collect::<Vec<String>>();
        let output = data.next().unwrap().split(' ').collect::<Vec<&str>>();

        for digit in output {
            match digit.len() {
                2 | 3 | 4 | 7 => count += 1,
                _ => continue,
            }
        }
    }

    println!("{}", count);

    Ok(())
}
