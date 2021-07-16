pub struct Person {
  talkative: bool,
  dumb: bool,
  age: i32,
  name: String,
}

pub fn assign_person(p: Person) -> bool {
  if p.dumb == true  {
    return false;
  } else if p.talkative != true {
    return false;
  }

  return true;
}