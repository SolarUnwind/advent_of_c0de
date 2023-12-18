use std::fs;
use std::collections::HashSet;

fn main() {
    let input = fs::read_to_string("../6.txt").unwrap();
    let signal : Vec<char> = input.chars().collect();
  
    let size = signal.len();
    let mut index = 0;
    
    for it in signal.windows(4) {
        let chunk: HashSet<_> = it.iter().collect();
        if chunk.len() == 4 {
            break;
        }
        index += 1;
    }

    println!("Index : {}", index + 4);
    index = 0;
    for it in signal.windows(14) {
        let chunk: HashSet<_> = it.iter().collect();
        if chunk.len() ==  14 {
            break;
        }
        index += 1;
    }
    println!("Index : {}", index + 14);
}