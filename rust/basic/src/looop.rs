pub fn looooooooop() {
  let mut i = 0;
  loop {
    if i == 10 {
      break
    }
    println!("looooop - {}", i);
    i += 1;
  }
}

pub fn fooor() {
  let score = [80,60,40,47,70,59];
  let mut passed = 0;
  for item in score.iter() {
    if item > &60 {
      passed += 1;
    }
  }
  println!("{} score passed", passed);
}