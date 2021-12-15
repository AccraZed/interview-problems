use std::{error::Error, fs};

pub fn part_1() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day9_input.txt")?;
    let hm: Vec<Vec<i32>> = raw
        .split('\n')
        .map(|i| {
            i.chars()
                .map(|j| j.to_digit(10).unwrap() as i32)
                .collect::<Vec<i32>>()
        })
        .collect();

    let mut total = 0;
    for (i, row) in hm.iter().enumerate() {
        for (j, val) in row.iter().enumerate() {
            let high_above = i == 0 || *val < hm[i - 1][j];
            let high_below = i + 1 == hm.len() || *val < hm[i + 1][j];
            let high_left = j == 0 || *val < hm[i][j - 1];
            let high_right = j + 1 == hm[i].len() || *val < hm[i][j + 1];

            if high_above && high_below && high_left && high_right {
                total += val + 1
            }
        }
    }

    println!("{}", total);

    Ok(())
}

pub fn part_2() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day9_input.txt")?;
    // Init height map
    let mut height_map: Vec<Vec<i32>> = raw
        .split('\n')
        .map(|i| {
            i.chars()
                .map(|j| j.to_digit(10).unwrap() as i32)
                .collect::<Vec<i32>>()
        })
        .collect();

    // Number of Islands, in a nutshell
    let mut sizes = Vec::<i32>::new();
    for (i, row) in height_map.clone().iter().enumerate() {
        for (j, val) in row.iter().enumerate() {
            if *val < 9 {
                sizes.push(mark_basin(i as i32, j as i32, &mut height_map)?);
            }
        }
    }

    sizes.sort_by(|a, b| b.cmp(a)); // Reverse sort
    println!("{}", sizes[..3].iter().fold(1, |acc, size| acc * size));

    Ok(())
}

fn mark_basin(i: i32, j: i32, height_map: &mut Vec<Vec<i32>>) -> Result<i32, Box<dyn Error>> {
    if !in_bounds(i, j, height_map) {
        return Ok(0);
    }
    if height_map[i as usize][j as usize] == 9 {
        return Ok(0);
    }

    height_map[i as usize][j as usize] = 9;

    let mut total = 1;
    total += mark_basin(i + 1, j, height_map)?;
    total += mark_basin(i - 1, j, height_map)?;
    total += mark_basin(i, j + 1, height_map)?;
    total += mark_basin(i, j - 1, height_map)?;

    Ok(total)
}

fn in_bounds(i: i32, j: i32, height_map: &Vec<Vec<i32>>) -> bool {
    if i < 0 || j < 0 || i >= height_map.len() as i32 || j >= height_map[0].len() as i32 {
        return false;
    }

    true
}
