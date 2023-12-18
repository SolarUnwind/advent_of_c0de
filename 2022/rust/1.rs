use std::fs;


fn main() {
    let input = fs::read_to_string("../1.txt").unwrap();
    let lines : Vec<&str> = input.split("\n").collect();

    let mut cur_value : i32 = 0;
    let mut totals : Vec<i32> = Vec::new();
    let mut cur_index : i32 = 0;
    for line in lines {
        if (line.len() > 0) {
            cur_value += line.parse::<i32>().unwrap();
        }
        else {
            totals.push(cur_value);
            cur_index += 1;
            cur_value = 0;
        }
    }
    totals.sort();
    println!("Max Value: {}",  totals[totals.len() - 1]);
    let top3 = totals[totals.len() - 1] + totals[totals.len() - 2] + totals[totals.len() - 3];
    println!("Top 3 Value: {}",  top3);
}