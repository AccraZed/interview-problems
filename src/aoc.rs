#![allow(dead_code)]
#![allow(unused_imports)]

use std::error::Error;

pub mod day10;
pub mod day11;
pub mod day12;
pub mod day4;
pub mod day5;
pub mod day6;
pub mod day7;
pub mod day8;
pub mod day9;

pub fn run_aoc() -> Result<(), Box<dyn Error>> {
    day12::part_1()?;
    Ok(())
}
