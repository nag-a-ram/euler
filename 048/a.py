"""
Copyright 2021 Alee Azam
--------- ---- ---- ----

I am not proud that it's brute force, nonetheless it's pretty darn fast.

real	0m0.044s
user	0m0.039s
sys	0m0.004s
"""

total = 0

for i in range(1, 1001):
    total += i ** i

print(str(total)[-10:])