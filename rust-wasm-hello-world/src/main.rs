use std::io::{self, BufRead};

fn main() {
    println!("Rust main function! (Arch={})", std::env::consts::ARCH);

    // Call the Go function
    println!("Calling go_hello function");
    unsafe {
        go_hello();
    };

    println!("Type some text:");
    let stdin = io::stdin();
    let line = stdin.lock().lines().next().unwrap().unwrap();
    println!("You entered: {}", line);
}

extern "C" {
    fn go_hello();
}

#[no_mangle]
pub extern "C" fn rust_hello() {
    println!("Hello from Rust!");
}
