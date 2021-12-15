use std::{collections::HashMap, error::Error, fs, path::Iter};

pub fn part_1() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day8_input.txt")?;
    let lines = raw.split('\n');

    let mut count = 0;
    for line in lines {
        let mut data = line.split(" | ");
        let legend = data
            .next()
            .unwrap()
            .split(' ')
            .map(|digit| {
                let mut chars = digit.chars().collect::<Vec<char>>();
                chars.sort();
                return String::from_iter(chars);
            })
            .collect::<Vec<String>>();
        let output = data.next().unwrap().split(' ').collect::<Vec<&str>>();

        for digit in output {
            match digit.len() {
                2 | 3 | 4 | 7 => count += 1,
                _ => continue,
            }
        }
    }

    println!("{}", count);

    Ok(())
}

pub fn part_2() -> Result<(), Box<dyn Error>> {
    let raw = fs::read_to_string("src/aoc/day8_input.txt")?;
    let lines = raw.split('\n');

    let mut count = 0;

    for line in lines {
        let mut data = line.split(" | ");
        let mut digits: Vec<&str> = data.next().unwrap().split_whitespace().collect();
        digits.sort_by(|a, b| a.len().cmp(&b.len()));
        let digits = digits
            .iter()
            .map(|digit| {
                digit.chars().fold(0, |val, segment| {
                    val | (1 as u8) << (segment as u8 - 'a' as u8)
                })
            })
            .collect::<Vec<u8>>();

        let one = digits[0];
        println!("one: {:b}", one);
        let seven = digits[1];
        println!("seven: {:b}", seven);
        let four = digits[2];
        println!("four: {:b}", four);
        let eight = digits[9];
        println!("eight: {:b}", eight);

        let five_segs = &digits[3..6];
        println!(
            "segs: {:b} {:b} {:b}",
            five_segs[0], five_segs[1], five_segs[2],
        );

        /*
         aaaa
        b    c
        b    c
         dddd.
        e    f
        e    f
         gggg
         */

        let c_f = one;
        let b_d = one ^ four;

        let two = *five_segs
            .iter()
            .filter(|digit| (**digit | b_d != **digit) && (**digit | c_f != **digit))
            .next()
            .unwrap();

        let c = c_f & two;
        let f = c_f ^ c;
        let d = b_d & two;

        let three = *five_segs
            .iter()
            .filter(|digit| **digit & c_f == c_f)
            .next()
            .unwrap();

        let e = (two ^ three) ^ f;

        let decoder = HashMap::from([
            (one, 1),
            (two, 2),
            (three, 3),
            (four, 4),
            (eight ^ c ^ e, 5),
            (eight ^ c, 6),
            (seven, 7),
            (eight, 8),
            (eight ^ e, 9),
            (eight ^ d, 0),
        ]);

        println!("Decoder: {:#?}", decoder);
        let output = data
            .next()
            .unwrap()
            .split(' ')
            .map(|digit| {
                let key: u8 = digit.chars().fold(0, |val, segment| {
                    val | (1 as u8) << (segment as u8 - 'a' as u8)
                });
                println!("{}", key);

                return *decoder.get(&key).unwrap();
            })
            .collect::<Vec<i32>>();

        println!("Output: {:#?}", output);
        let num: i32 = output
            .into_iter()
            .map(|i| i.to_string())
            .collect::<String>()
            .parse()?;

        println!("Num: {}", num);
        count += num;
        println!("Count: {}", count);
    }

    println!("{}", count);

    Ok(())
}
