N,x = map(int, input().split())
A = list(map(int, input().split()))

A.sort()
ans = 0

for i in range(N):
  if x > A[i]:
    if i+1 == N:
      print(N-1)
      exit()
    else:
      ans += 1
      x -= A[i]
  elif x == A[i]:
    ans += 1
    print(ans)
    exit()
  else:
    print(ans)
    exit()