best = 0
def createPalindrome(digits):
    global best
    num = 0
    while len(str(num)) <= digits:
        num2 = 0
        while len(str(num2)) <= digits:
            if (x := num2 * num) > best:
                if strReverse(str(x)) == str(x):
                    best = x
            num2 += 1
        num += 1
    return best
    
def strReverse(txt):
    temp = [char for char in txt]
    temp.reverse()
    newStr = ""
    for char in temp:
        newStr += char
    return newStr
print(createPalindrome(3))
