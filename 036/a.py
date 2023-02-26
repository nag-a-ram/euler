"""
Written by @agzg

real    0m0.701s
user    0m0.684s
sys     0m0.017s
"""

total = 0

for i in range(1, 1000001):
    b10, b2 = str(i), str(bin(i))[2:]

    if b10 == b10[::-1] and b2 == b2[::-1]:
        total += i

print(total)
