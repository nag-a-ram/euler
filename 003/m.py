from math import floor as flr

def SieveOfEratosthenes(maxNum):
    primeNums = []
    allNums = []
    for I in range(maxNum - 1):
        allNums.append(True)
    
    for i in range(flr(maxNum**0.5) + 1):
        if (i != 0 and i != 1):
            multNum = 2
            if (allNums[i]):
                while ((j := multNum * i) <= maxNum):
                    allNums[j - 2] = False
                    multNum += 1
    for i, _ in enumerate(allNums):
        if (allNums[i]):
            primeNums.append(i + 2)
    return primeNums

def primeFactorOf(n):
    num = n
    p = 2
    while num > p**2:
        if num % p == 0:
            num /= p
        else:
            p += 1
    return num

print(primeFactorOf(600851475143))
