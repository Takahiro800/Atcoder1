S = input()

a = S.count('0')
b = S.count('1')

ans = min(a,b) * 2
print(ans)
