def factorial(num):
    out = 1
    for i in range(2, num + 1):
        out *= i
    return out

print("the answer is", sum(
    [int(x) for x in list(str(factorial(100)))]
))

# written by @aliazam (Alaz)
