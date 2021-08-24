"""
Copyright 2021 Alee Azam
--------- ---- ---- ----

real    0m4.631s
user    0m4.625s
sys     0m0.004s
"""

def prime(n):
    if n == 1 or n == 0:
        return False

    for i in range(2, int(n**0.5)+1):
        if n%i == 0:
            return False
    return True

def trunc(n):
    left, right = str(n), str(n)

    while len(left) > 0:
        if not prime(int(left)) or not prime(int(right)):
            return False

        left, right = left[1:], right[:-1]
        
    return True

if __name__ == "__main__":
    n      = 8  # Current number.
    truncs = 0  # Truncatable primes so far.
    total  = 0  # Sum of truncatable primes.

    while truncs < 11:
        if trunc(n):
            total += n
            truncs += 1
        n += 1

    print(total)
