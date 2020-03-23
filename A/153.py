M = input().split()

H = int(M[0])
A = int(M[1])

quo = H % A

if quo == 0:
  ans = H // A
else:
  ans = (H // A) + 1

print(ans)