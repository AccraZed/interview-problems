use std::{
    cmp::{max, min},
    collections::HashMap,
    error::Error,
    fmt, fs,
};

#[derive(fmt::Debug)]
struct Line {
    x1: i32,
    x2: i32,
    y1: i32,
    y2: i32,
}

impl Line {
    pub fn new(points: &[i32]) -> Result<Line, Box<dyn Error>> {
        let x1 = points[0];
        let y1 = points[1];
        let x2 = points[2];
        let y2 = points[3];
        Ok(Line { x1, x2, y1, y2 })
    }

    pub fn straight_path(&self) -> Result<Vec<(i32, i32)>, Box<dyn Error>> {
        let mut coords: Vec<(i32, i32)> = Vec::new();

        let x_dif = (self.x2 - self.x1).abs();
        let y_dif = (self.y2 - self.y1).abs();
        let x_min = min(self.x1, self.x2);
        let y_min = min(self.y1, self.y2);

        for y in y_min..y_min + y_dif + 1 {
            for x in x_min..x_min + x_dif + 1 {
                coords.push((x, y));
            }
        }

        Ok(coords)
    }

    pub fn path(&self) -> Result<Vec<(i32, i32)>, Box<dyn Error>> {
        if self.x1 == self.x2 || self.y1 == self.y2 {
            return Ok(self.straight_path()?);
        }

        let mut coords: Vec<(i32, i32)> = Vec::new();
        let dif = (self.x2 - self.x1).abs();
        let x_min = min(self.x1, self.x2);
        let y_min = min(self.y1, self.y2);
        let y_max = max(self.y1, self.y2);

        if (self.x2 - self.x1) * (self.y2 - self.y1) > 0 {
            for i in 0..dif + 1 {
                coords.push((x_min + i, y_min + i));
            }
        } else {
            for i in 0..dif + 1 {
                coords.push((x_min + i, y_max - i));
            }
        }

        Ok(coords)
    }
}

pub fn part_1() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day5_input.txt")?;
    let lines = raw
        .split('\n')
        // Parse each line // 102,578 -> 363,317
        .map(|line| {
            let quad = line
                .split(" -> ")
                // Parse each point // 102,578
                .map(|point| {
                    point
                        .split(',')
                        // Parse each coord // 102
                        .map(|coord| coord.parse::<i32>().unwrap())
                        .collect::<Vec<i32>>()
                })
                .flatten()
                .collect::<Vec<i32>>();

            return Line::new(quad.as_slice()).unwrap();
        });

    let straight_lines = lines
        .filter(|line| line.x1 == line.x2 || line.y1 == line.y2)
        .collect::<Vec<Line>>();

    println!("{:#?}", straight_lines);

    let mut coord_counts: HashMap<(i32, i32), i32> = HashMap::new();
    for line in straight_lines {
        line.straight_path()?
            .iter()
            .for_each(|coord| *coord_counts.entry(*coord).or_insert(0) += 1);
    }

    let mut count = 0;
    for freq in coord_counts.values() {
        if *freq > 1 {
            count += 1;
        }
    }

    println!("{}", count);

    Ok(())
}

pub fn part_2() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day5_input.txt")?;
    let lines = raw
        .split('\n')
        // Parse each line // 102,578 -> 363,317
        .map(|line| {
            let quad = line
                .split(" -> ")
                // Parse each point // 102,578
                .map(|point| {
                    point
                        .split(',')
                        // Parse each coord // 102
                        .map(|coord| coord.parse::<i32>().unwrap())
                        .collect::<Vec<i32>>()
                })
                .flatten()
                .collect::<Vec<i32>>();

            return Line::new(quad.as_slice()).unwrap();
        });

    let mut coord_counts: HashMap<(i32, i32), i32> = HashMap::new();
    for line in lines {
        line.path()?
            .iter()
            .for_each(|coord| *coord_counts.entry(*coord).or_insert(0) += 1);
    }

    let mut count = 0;
    for freq in coord_counts.values() {
        if *freq > 1 {
            count += 1;
        }
    }

    println!("{}", count);

    Ok(())
}
