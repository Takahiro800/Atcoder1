X = int(input())

a = X // 500
b = (X % 500) // 5

Max = a*1000 + b*5

print(Max)