use std::fs;
use std::collections::HashSet;

fn main() {
    let input = fs::read_to_string("../4.txt").unwrap();
    let lines : Vec<&str> = input.split("\n").collect();

    let mut score_part1 : u32 = 0;
    let mut score_part2 : u32 = 0;

    // Part 1
    for line in &lines {
        let pair : Vec<&str> = line.trim().split(",").collect();
        let first : Vec<&str> = pair[0].split("-").collect();
        let second : Vec<&str> = pair[1].split("-").collect();

        let first_start = first[0].parse::<i32>().unwrap();
        let first_end = first[1].parse::<i32>().unwrap();
        let second_start = second[0].parse::<i32>().unwrap();
        let second_end = second[1].parse::<i32>().unwrap();
        
        if (first_start <= second_start && first_end >= second_end) || ( second_start <= first_start && second_end >= first_end ) {
            score_part1 += 1;
        }

        if (first_start <= second_end && second_start <= first_end) {
            score_part2 += 1;
        }
    }

    
    println!("Score 1: {}", score_part1);
    println!("Score 2: {}", score_part2);
    
}