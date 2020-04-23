A,B = map(int, input().split())
k = 0
ans = 1
while ans < B:
  ans = ans - 1 + A
  k += 1

print(k)