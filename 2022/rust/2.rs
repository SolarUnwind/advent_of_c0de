use std::fs;
use std::collections::HashMap;

fn main() {
    let input = fs::read_to_string("../2.txt").unwrap();
    let lines : Vec<&str> = input.split("\n").collect();

    let mut responses = HashMap::from([("X", 1), ("Y", 2), ("Z", 3)]);
    let mut part1 = HashMap::from(
        [
            ("A", HashMap::from([("X", 3), ("Y", 6), ("Z", 0)])),
            ("B", HashMap::from([("X", 0), ("Y", 3), ("Z", 6)])),
            ("C", HashMap::from([("X", 6), ("Y", 0), ("Z", 3)])),
        ]
    );

    let mut results = HashMap::from([("X", 0), ("Y", 3), ("Z", 6)]);
    let mut part2 = HashMap::from(
        [
            ("A", HashMap::from([("X", 3), ("Y", 1), ("Z", 2)])),
            ("B", HashMap::from([("X", 1), ("Y", 2), ("Z", 3)])),
            ("C", HashMap::from([("X", 2), ("Y", 3), ("Z", 1)])),
        ]
    );

    let mut score_part1 : i32 = 0;
    let mut score_part2 : i32 = 0;
    for line in lines {
        let round : Vec<&str> = line.trim().split(" ").collect();
        if (line.len() > 0) {
            let opponent = round[0];
            let response = round[1];
            score_part1 += responses.get(&response).unwrap();
            score_part2 += results.get(&response).unwrap();
            
            
            let p1 = part1.get(&opponent).unwrap();
            let p2 = part2.get(&opponent).unwrap();
            score_part1 += p1.get(&response).unwrap();
            score_part2 += p2.get(&response).unwrap();
   
        }
    }
    
    println!("Score 1: {}", score_part1);
    println!("Score 2: {}", score_part2);
    
}

