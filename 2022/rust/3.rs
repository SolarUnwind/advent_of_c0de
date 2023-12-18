use std::fs;
use std::collections::HashSet;

fn char_to_priority(input : char) -> u32 {
    let lower_case : u32 = 96;
    let upper_case : u32 = 38;

    if input.is_uppercase() {
        input as u32 - upper_case
    }
    else {
        input as u32 - lower_case
    }
}


fn main() {
    let input = fs::read_to_string("../3.txt").unwrap();
    let lines : Vec<&str> = input.split("\n").collect();

    let mut score_part1 : u32 = 0;
    let mut score_part2 : u32 = 0;

    // Part 1
    for line in &lines {
        let (left, right) = line.trim().split_at(line.len() / 2);
        let half : HashSet<char> = left.chars().collect();

        for c in half {
            if right.contains(c) {
                score_part1 += char_to_priority(c);
            }
        }
    }

    // Part 2
    let mut i  = 0;
    let size = lines.len();
    while i < size - 1 {
        let first : HashSet<char> = lines[i].chars().collect();
        for c in first {
            if lines[i+1].contains(c) && lines[i+2].contains(c) {
                score_part2 += char_to_priority(c);
            }
        }
        i += 3;
    }
    
    println!("Score 1: {}", score_part1);
    println!("Score 2: {}", score_part2);
    
}