use std::{collections::HashMap, error::Error, fs};

pub fn part_1() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day10_input.txt")?;
    let lines = raw.split('\n');

    let closing_score = HashMap::from([(')', 3), (']', 57), ('}', 1197), ('>', 25137)]);
    let closing = HashMap::from([(')', '('), (']', '['), ('}', '{'), ('>', '<')]);

    let mut total = 0;
    for line in lines {
        let mut stack = Vec::<char>::new();

        for c in line.chars() {
            match c {
                ')' | ']' | '}' | '>' => {
                    if *stack.last().unwrap() != *closing.get(&c).unwrap() {
                        total += closing_score.get(&c).unwrap();
                        break;
                    } else {
                        stack.pop();
                    }
                }
                _ => stack.push(c),
            }
        }
    }

    println!("{}", total);

    Ok(())
}

pub fn part_2() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day10_input.txt")?;
    let lines = raw.split('\n');

    let char_score = HashMap::from([('(', 1), ('[', 2), ('{', 3), ('<', 4)]);
    let closing = HashMap::from([(')', '('), (']', '['), ('}', '{'), ('>', '<')]);

    let mut scores: Vec<u128> = Vec::new();
    'line: for line in lines {
        let mut cur_score: u128 = 0;
        let mut stack = Vec::<char>::new();

        for c in line.chars() {
            match c {
                ')' | ']' | '}' | '>' => {
                    if *stack.last().unwrap() != *closing.get(&c).unwrap() {
                        continue 'line;
                    } else {
                        stack.pop();
                    }
                }
                _ => stack.push(c),
            }
        }

        println!("{:?}", stack);

        stack.reverse();
        for c in stack {
            cur_score *= 5;
            cur_score += *char_score.get(&c).unwrap();
        }

        scores.push(cur_score);
    }

    scores.sort();
    println!("{}", scores[scores.len() / 2]);

    Ok(())
}
