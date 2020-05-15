N = int(input())
V = list(map(int, input().split()))

V.sort()
ans = 0

for count,i in enumerate(V):
  if count == 0 or count == 1:
    ans += i / (2**(N-1))
  else:
    ans += i / (2**(N-count))
else:
  print(ans)