use std::fs;
use std::collections::HashMap;
use std::path::PathBuf;
use std::path::Path;
use std::iter::FromIterator;
use std::cmp;

fn main() {
    let input = fs::read_to_string("../7.txt").unwrap();
    let lines : Vec<&str> = input.split("\n").collect();
    let mut paths = HashMap::new();

    let mut path  = Vec::new();

    for line in &lines {
        let parsed_line : Vec<&str> = line.trim().split_whitespace().collect();

        if parsed_line[1] == "cd" {
            if parsed_line[2] == ".." {
                path.pop();
            }
            else {
                path.push(parsed_line[2]);
            }
        }
        else if parsed_line[1] == "ls" {
            continue;
        }
        else if parsed_line[0] == "dir" {
            continue;
        }
        else {

            let size = parsed_line[0].parse::<u32>().unwrap();
            let n = path.len();
            
            for i in 0..n {
                let cur_path = PathBuf::from_iter(&path[..=i]);
                *paths.entry(cur_path).or_insert(0) += size;
            }
        }     
    }

    let mut score_part1 : u32 = 0;
    let mut score_part2 : u32 = 4294967295;
    let limit : u32 = 100000;
    let max : u32 = 70000000 - 30000000;
    let used : &u32 = paths.get(&Path::new("/").to_path_buf()).unwrap();
    let need_to_free : u32 = used - max;
    println!("Need to free: {}", need_to_free);

    for (key, val) in paths.iter() {
        if val <= &limit {
            score_part1 += val;
        }
        if val >= &need_to_free {
            score_part2 = cmp::min(val.clone(), score_part2);
        }
    }

    println!("Score 1: {}", score_part1);
    println!("Score 2: {}", score_part2);
}