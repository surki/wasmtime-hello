const print = @import("std").debug.print;

export fn hello_world() void {
    print("Hello, {s}!\n", .{"World"});
}

pub fn main() void {
    hello_world();
}