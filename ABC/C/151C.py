N, K, M = map(int, input().split())
A = list(map(int, input().split()))

sumA = sum(A)

ans = N * M - sumA
if ans > K:
  print(-1)
elif ans < 0:
  print(0)
else:
  print(ans)