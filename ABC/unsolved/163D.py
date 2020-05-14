import math

N,K = input().split()
N = int(N)
K = int(K)

l = K
ans = 0
while l <= N+1:
  ans += math.factorial(N+1) // (math.factorial(N-l+1) * math.factorial(l))
  l += 1

print(ans)