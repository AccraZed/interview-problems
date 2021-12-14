use std::process;

mod aoc;

fn main() {
    if let Err(e) = aoc::run_aoc() {
        println!("Error running aoc: {}", e);
        process::exit(1);
    }
}
