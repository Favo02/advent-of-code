use std::io::{self, BufRead};

fn main() {
  let stdin = io::stdin();
  let mut lines = Vec::new();

  for line in stdin.lock().lines() {
    lines.push(line.unwrap());
  }

  let mut part1 = 0;
  let mut part2 = 0;

  for line in lines {
    let (f1, l1) = first_and_last(&line);

    let line2 = replace_digits(line);
    let (f2, l2) = first_and_last(&line2);

    part1 += f1 * 10 + l1;
    part2 += f2 * 10 + l2;
  }

  println!("Part 1: {part1}");
  println!("Part 2: {part2}");
}

fn first_and_last(line : &String) -> (u32, u32) {
  let mut first = 0;
  let mut last = 0;
  let mut found_first = false;

  for c in line.chars() {
    if c.is_digit(10) {
      if !found_first {
        first = c.to_digit(10).unwrap();
        found_first = true;
      }
      last = c.to_digit(10).unwrap();
    }
  }

  (first, last)
}

fn replace_digits(mut line : String) -> String {
  line = line.replace("one", "o1e");
  line = line.replace("two", "t2o");
  line = line.replace("three", "t3e");
  line = line.replace("four", "f4r");
  line = line.replace("five", "f5e");
  line = line.replace("six", "s6x");
  line = line.replace("seven", "s7n");
  line = line.replace("eight", "e8t");
  line.replace("nine", "n9e")
}
