N,K,Q = map(int, input().split())
ary = [0] * (N+1)

for i in range(Q):
  A = int(input())
  ary[A] += 1

i = 1
for i in range(1, N+1):
  print('Yes' if ((K + ary[i] - Q) > 0) else 'No')
