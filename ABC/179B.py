N = int(input())


max = 0
stock = 0
for i in range(N):
  a,b = map(int, input().split())
  if a == b:
    stock += 1
    max = stock
  else:
    max = stock
    stock = 0

print('Yes' if (max > 2) else "No")
