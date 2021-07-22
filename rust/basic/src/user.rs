pub struct Person {
  pub talkative: bool,
  pub dumb: bool,
  pub age: i32,
  pub name: String,
}

pub fn assign_person(p: Person) -> bool {
  if p.dumb == true  {
    return false;
  } else if p.talkative != true {
    return false;
  }

  return true;
}