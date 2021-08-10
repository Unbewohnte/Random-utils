use rand::Rng;

// returns pseudo-randomly generated password from chars
fn generate_password(chars: String, length: u32) -> String {
    let mut rng = rand::thread_rng();
    
    let mut password: String = String::with_capacity(length as usize);
    for _ in 0..length {
        // wtf
        let random_index:usize = rng.gen_range(0..chars.len() - 1);
        
        let random_char:char = chars.chars().nth(random_index).unwrap();
        
        password += random_char.to_string().as_str();
    }

    return password;
}

fn main() {
    // get command line arguments
    let args: Vec<String> = std::env::args().collect();    

    let mut pass_length: u8 = 8; // default password length

    // check for given arguments
    if args.len() > 1 {
        // has some arguments -> get the [1] one and check for an error 
        let pass_length_res = args[1].parse::<u8>();  
    
        match pass_length_res {
            Ok(n) => {
                // check for 0
                if n != 0 {
                    // not 0 -> safely switch default pass_length value with a new one
                    pass_length = n;                   
                }
            },
            Err(e) => println!("Please, specify an u8 (1..255) integer as the first argument next time: {}",e),
        }
    }  

    // characters, from which we will construct a password
    let chars: String = String::from("ABCDEFGHIJKLMNOPQRSUVWXYZabcdefghijklmnopqrstuvwxyz1234567890{}[]()!@#$%^&*~");

    // generate password
    let generated_password: String = generate_password(chars, pass_length as u32);

    // and print it
    println!("{}",generated_password);
}
