def fib():
	a, b = 1, 1
	while 1:
		yield a
		a, b = b, a+b
	
for i, n in enumerate(fib()):
	if len(str(n)) == 1000:
		print(i+1)
		break

