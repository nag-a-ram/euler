// Project Euler
// Problem 14
// Not Timed Per Say, But Approx 1.15 Secs
// Written By: @Abdul-Muiz-Iqbal

fn main() {
    let mut best = 0; // Best Is The Longest Chain THat A Number Makes
    for i in 1..1_000_001 { // Iterate From 1 To 1 Million
        let new = collatz_sequence(i); // Get The Chain At i
        if new > best { // Simple Checking For Bigger Chain
            best = new;
            println!("{}", i); // Getting The Starting Number With Current Highest Chain
        }
    }
}

fn collatz_sequence(start: u64) -> u64 {
    let mut c = 1; // Chain Counter
    let mut num = start; // Helper Variable So I Don't Have To Mess With Ownership
    while num != 1 { // While The Chain Has Not Ended
        if num % 2 != 0 { // If Number Is Odd
            num = (3 * num) + 1; // 3 Multiplied By Num Plus 1
        }else { // Otherwise Num Is Even
            num = num / 2; // Halve Number
        }
        c += 1; // Increment Chain Counter
    }
    c // Return Chain Count
}
