use std::{error::Error, fs};

const NUM_CYCLES: i32 = 256;
const PROD_RATE: i32 = 9;

pub fn part_1_2() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day6_input.txt")?;
    let fish_ages = raw.split(",");
    let mut age_counts: [u128; PROD_RATE as usize] = [0; PROD_RATE as usize];

    fish_ages.for_each(|age| age_counts[age.parse::<usize>().unwrap()] += 1);

    for i in 0..NUM_CYCLES {
        age_counts[((i + 7) % PROD_RATE) as usize] += age_counts[(i % PROD_RATE) as usize];

        println!("Day {}: {:?}", i, age_counts);
    }

    println!("{}", age_counts.iter().fold(0, |sum, cnt| sum + cnt));

    Ok(())
}
