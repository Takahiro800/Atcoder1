N, M = map(int, input().split())
A = list(map(int, input().split()))
A.sort(reverse=True)

n = M - 1
K = sum(A)

if A[n] < K / (4 *M):
  print('No')
else:
  print('Yes')