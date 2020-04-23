N = int(input())

A = list(map(int, input().split()))
a = [0] * N

for i in range(N-1):
  k = A[i]
  a[k-1] += 1

for l in range(N):
  print(a[l])
