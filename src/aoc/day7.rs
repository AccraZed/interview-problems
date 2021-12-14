use std::cmp::min;
use std::error::Error;
use std::fs;

pub fn part_1() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day7_input.txt")?;
    let mut positions: Vec<i32> = raw
        .split(',')
        .map(|pos| pos.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();
    positions.sort();

    let target = positions[positions.len() / 2];

    let mut total_fuel = 0;
    for pos in positions {
        total_fuel += (target - pos).abs();
    }

    println!("{}", total_fuel);

    Ok(())
}

pub fn part_2() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day7_input.txt")?;
    let mut positions: Vec<i32> = raw
        .split(',')
        .map(|pos| pos.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();
    positions.sort();

    let mut min_fuel = i32::MAX;
    for target_pos in positions[0]..positions[positions.len() - 1] {
        let mut total_fuel = 0;
        for pos in &positions {
            let dist = (target_pos - pos).abs();
            total_fuel += ((dist) * (dist + 1)) / 2;
        }

        min_fuel = min(min_fuel, total_fuel);
    }

    println!("{}", min_fuel);

    Ok(())
}
