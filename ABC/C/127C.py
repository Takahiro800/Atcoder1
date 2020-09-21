N,M = map(int, input().split())

n = 0
m = 999999

for i in range(M):
  a,b = map(int, input().split())
  n = max(a, n)
  m = min(b , m)

print(max(m-n+1, 0))
