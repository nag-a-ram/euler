"""
Copyright 2021 Alee Azam

real    0m0.034s
user    0m0.029s
sys     0m0.005s
"""

total = 0
w, l, i = 3, 2, 1

while i <= 1001 ** 2:
	total += i
	i += l
	
	if i == w ** 2:
		w += 2
		l += 2
		
print(total)
