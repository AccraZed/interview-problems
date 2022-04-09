use std::{error::Error, fs};

struct OctopusGrid {
    grid: [[u32; 10]; 10],
    flashes: u32,
}

struct Coord {
    i: usize,
    j: usize,
}

impl OctopusGrid {
    fn new(filename: String) -> Result<OctopusGrid, Box<dyn Error>> {
        let raw = fs::read_to_string(filename)?;

        let mut grid: [[u32; 10]; 10] = [[0; 10]; 10];
        raw.split_whitespace().enumerate().for_each(|(i, row)| {
            row.chars()
                .enumerate()
                .for_each(|(j, val)| grid[i][j] = val.to_digit(10).unwrap())
        });
        let flashes = 0;

        Ok(OctopusGrid { grid, flashes })
    }

    fn simulate_step(&mut self) -> Result<(), Box<dyn Error>> {
        let mut pops: Vec<Coord> = Vec::new();

        for (i, row) in self.grid.iter_mut().enumerate() {
            for (j, val) in row.iter_mut().enumerate() {
                *val += 1;
                if *val > 9 {
                    pops.push(Coord { i, j })
                }
            }
        }

        let first = pops.last().unwrap();
        self.grid[first.i][first.j] = 0;
        self.flashes += 1;

        let dirs: [[i32; 2]; 4] = [[1, 0], [-1, 0], [0, 1], [0, -1]];
        while !pops.is_empty() {
            let coord = pops.pop().unwrap();

            for dir in dirs {
                if self.grid[(coord.i as i32 + dir[0]) as usize][(coord.j as i32 + dir[1]) as usize]
                    != 0
                {
                    self.grid[coord.i][coord.j] += 1;
                    self.flashes += 1;
                }
            }
        }

        Ok(())
    }
}

const STEPS: i32 = 100;
pub fn part_1() -> Result<(), Box<dyn Error>> {
    let mut octo_grid = OctopusGrid::new(String::from("src/aoc/day11_input.txt"))?;

    for _ in 0..STEPS {
        octo_grid.simulate_step()?;
    }
    Ok(())
}
