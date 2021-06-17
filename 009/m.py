def isPythagoreanTriple(a, b, c):
    if not a < b < c: return False
    if a ** 2 + b ** 2 == c ** 2: return True
    return False
    
def getTripletsOf(n):
    triplets = []
    for I in range(n + 1):
        for J in range(n + 1 - 500):
            for K in range(n + 1 - 750):
                if I + J + K == n:
                    triplets.append([I, J, K])
    return triplets

# ans = []
# for triplet in getTripletsOf(1000):
#     if isPythagoreanTriple(*triplet):
#         ans.append(triplet)

# ans_ = 1
# for I in ans:
#     ans_ *= I
#     
# print(ans_)

def generatePythagoreanTriplet(m, n):
    a = (m ** 2) - (n ** 2)
    b = 2 * (m * n)
    c = (m ** 2) + (n ** 2)
    return (a, b, c)
 
ans = []
for I in range(1, 1001):
    for J in range(I + 1, 1001):
        a, b, c = generatePythagoreanTriplet(J, I)
        if a + b + c == 1000:
            ans.append((a, b, c))
print(ans)
print(ans[0][0] * ans[0][1] * ans[0][2])
