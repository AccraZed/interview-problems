use std::error::Error;

pub mod day4;
pub mod day5;
pub mod day6;

pub fn run_aoc() -> Result<(), Box<dyn Error>> {
    day6::part_1()?;

    Ok(())
}
