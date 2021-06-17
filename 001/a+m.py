sum = 0
arr = [num for num in range(1, 1000) if num % 3 == 0 or num % 5 == 0]
for num in arr: sum += num

# a more compact solution proposed by @aliazam
print(
  "the answer is",
  sum([i for i in range(1, 1000) if num%3 == 0 or num % 5 == 0])
)  
