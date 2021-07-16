mod neighbor;
mod looop;
mod user;

fn main() {
    // must always remember that variables are immutable by default
    // and you should end everything with semicolons
    let num: i32 = 20;
    println!("{}", num);

    let tuple = (1,2,3,4,5);
    // {:?} is called debug trait (or something, I don't quite understand this yet)
    println!("{:?}", tuple);

    let array = [6,7,8,9,10];
    println!("{:?}", array);

    // printing specific index
    println!("{}", tuple.1);
    // or maybe this would work
    println!("{}", array[1]);
    
    let plus = add(5, 10);
    println!("{}", plus);

    let multi = neighbor::multiply(2, 5);
    println!("{}", multi);

    looop::fooor();
    looop::looooooooop();
}

fn add(x: i32, y: i32) -> i32 {
    x + y
}
