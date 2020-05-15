N,M,X = map(int,input().split())
A = list(map(int, input().split()))
a,b = 0,0

for i in range(M):
  if A[i] < X:
    a += 1
  else:
    b += 1
else:
  print(min(a,b))