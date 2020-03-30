m = input().split()
K = int(m[0])
N = int(m[1])

A = input().split()
a = [0] * N

for i in range(N):
  a[i] = int(A[i])

#import pdb; pdb.set_trace()
Sum = 0
n = N -1
for x in range(N):
  if x == n:
    Sum += (K - a[n]) + a[0]
  else:
    Sum += a[x+1] - a[x]

ans = [0] * N
for k in range(N):
  if k == n:
    ans[k] = (K - a[n]) + a[0]
  else:
    ans[k] = a[k+1] - a[k]

l = max(ans)

min_road = Sum - l
print(min_road)