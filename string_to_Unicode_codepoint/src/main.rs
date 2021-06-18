fn main() {
    let args: Vec<String> = std::env::args().collect();
    let given_string: String = String::from(args[1].clone());

    let mut counter:u128 = 0;
    for character in given_string.chars() {
        println!("{}: {} \t<---->\t {:x}", counter, character, character as u32);
        counter += 1;
    }
}
