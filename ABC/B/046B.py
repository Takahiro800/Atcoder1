N,K = map(int, input().split())
ans = 1

for i in range(N-1):
  ans *= (K-1)

print(ans * K)