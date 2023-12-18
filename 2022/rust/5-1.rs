use std::fs;
use std::collections::HashMap;

fn main() {
    let input = fs::read_to_string("../5.txt").unwrap();
    let lines : Vec<&str> = input.split("\n").collect();
    let mut stack : Vec<&str> = Vec::new();
    // Parse supplies
    let mut i  = 0;
    let size = lines.len();
    while i < size {
        stack.push(lines[i]);
        i += 1;

        if lines[i].len() <= 0 {
            break;
        }
    }

    let mut j = 1;
    let tmp : Vec<&str> = stack[i - j].split_whitespace().collect();
    let _size = tmp.len();

    let mut stacks = HashMap::new();

    for k in 0.._size {
        stacks.insert(k, Vec::<char>::new());
    }

    j = 2;
    while j <= i {
        let chars : Vec<char> = stack[i - j].chars().collect();

        for k in 0.._size {

            if chars[1 + 4*k].is_alphabetic() {
                stacks.get_mut(&k).unwrap().push(chars[1 + 4*k]);
            }
        }
        j += 1;   
    }

    // Process commands

    while i < size - 1 {
        i += 1;
        let cmd : Vec<&str> = lines[i].split(" ").collect();
        let n = cmd[1].parse::<usize>().unwrap();
        let from = cmd[3].parse::<usize>().unwrap();
        let to = cmd[5].parse::<usize>().unwrap();
        
        for k in 0..n {
            let item = stacks.get_mut(&(from - 1)).unwrap().pop().unwrap();
            stacks.get_mut(&(to - 1)).unwrap().push(item);
        }

    }

    for k in 0.._size {
        println!("{}: {:?}", k, stacks[&k]);
    }

}