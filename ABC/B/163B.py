N,M = input().split()
N = int(N)
M = int(M)

lis = map(int, input().split())
sum = sum(lis)

if N >= sum:
  print(N - sum)
else:
  print(-1)