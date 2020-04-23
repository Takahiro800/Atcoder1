N,M,C = map(int,input().split())
B = input().split()
count = 0

for i in range(N):
  A = input().split()
  sum = C
  for k in range(M):
    sum += int(A[k]) * int(B[k])
  if sum > 0:
    count += 1

print(count)