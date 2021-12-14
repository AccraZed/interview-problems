use std::{error::Error, fs};

const ROWS: usize = 5;
const COLS: usize = 5;

struct BingoBoard {
    board: [[i32; ROWS]; COLS],
    marked: [[bool; ROWS]; COLS],
    rows: [u8; ROWS],
    cols: [u8; COLS],
}

impl BingoBoard {
    pub fn new(numbers: &[i32]) -> Result<BingoBoard, Box<dyn Error>> {
        let mut board = [[0; ROWS]; COLS];
        for i in 0..(ROWS * COLS) {
            board[i / ROWS][i % COLS] = numbers[i];
        }
        let marked = [[false; ROWS]; COLS];
        let rows: [u8; ROWS] = [0; ROWS];
        let cols: [u8; COLS] = [0; COLS];

        Ok(BingoBoard {
            board,
            marked,
            rows,
            cols,
        })
    }

    pub fn mark(&mut self, target: i32) -> Result<bool, Box<dyn Error>> {
        for (i, row) in self.board.iter().enumerate() {
            for (j, num) in row.iter().enumerate() {
                // If number exists
                if *num == target {
                    // Mark number
                    self.rows[i] += 1;
                    self.cols[j] += 1;
                    self.marked[i][j] = true;

                    // Check for complete row
                    if self.rows[i] == ROWS as u8 || self.cols[j] == COLS as u8 {
                        return Ok(true);
                    }
                }
            }
        }

        Ok(false)
    }

    pub fn has_won(&self) -> Result<bool, Box<dyn Error>> {
        for i in 0..5 {
            if self.rows[i] == ROWS as u8 || self.cols[i] == COLS as u8 {
                return Ok(true);
            }
        }

        Ok(false)
    }

    fn score(&self, last_call: i32) -> Result<i32, Box<dyn Error>> {
        let sum = self
            .board
            .iter()
            .flatten()
            .zip(self.marked.iter().flatten())
            .fold(0, |sum, num| if !num.1 { sum + *num.0 } else { sum });

        Ok(sum * last_call)
    }
}

pub fn part1() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day4_input.txt")?;
    let mut split = raw.split_whitespace();

    let bingo_numbers: Vec<i32> = split
        .next()
        .unwrap()
        .split(',')
        .map(|item| item.parse::<i32>().unwrap())
        .collect();

    let boards: &mut Vec<BingoBoard> = &mut split
        .map(|num| num.parse::<i32>().unwrap())
        .collect::<Vec<i32>>()
        .as_slice()
        .chunks_exact(25)
        .map(|nums| BingoBoard::new(nums).unwrap())
        .collect::<Vec<BingoBoard>>();

    for bingo_number in bingo_numbers {
        for board in boards.iter_mut() {
            if board.mark(bingo_number)? {
                println!("{}", board.score(bingo_number)?);
                return Ok(());
            }
        }
    }

    Ok(())
}

pub fn part2() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day4_input.txt")?;
    let mut split = raw.split_whitespace();

    let bingo_numbers: Vec<i32> = split
        .next()
        .unwrap()
        .split(',')
        .map(|item| item.parse::<i32>().unwrap())
        .collect();

    let boards: &mut Vec<BingoBoard> = &mut split
        .map(|num| num.parse::<i32>().unwrap())
        .collect::<Vec<i32>>()
        .as_slice()
        .chunks_exact(25)
        .map(|nums| BingoBoard::new(nums).unwrap())
        .collect::<Vec<BingoBoard>>();

    let _target_wins = boards.len();
    let mut num_wins = 0;
    for bingo_number in bingo_numbers {
        for board in boards.iter_mut() {
            // Skip boards that already won
            if board.has_won()? {
                continue;
            }
            if board.mark(bingo_number)? {
                num_wins += 1;

                println!("Win #{}: {}", num_wins, board.score(bingo_number)?);
            }
        }
    }

    Ok(())
}
