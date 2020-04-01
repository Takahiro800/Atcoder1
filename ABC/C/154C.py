N = int(input())
A = list(map(int, input().split()))

A.sort()
ans = 'YES'

for i in range(N-1):
  if A[i] == A[i+1]:
    ans = 'NO'
    break

print(ans)
