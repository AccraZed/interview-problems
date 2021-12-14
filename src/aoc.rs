use std::error::Error;

pub mod day4;
pub mod day5;
pub mod day6;
pub mod day7;
pub mod day8;

pub fn run_aoc() -> Result<(), Box<dyn Error>> {
    day8::part_1()?;
    
    Ok(())
}
