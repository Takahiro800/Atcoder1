N = int(input())

i = 0
A = input().split()
ans = 'APPROVED'

while i < N:
  a = A[i]
  if int(a) % 2 == 0:
    if not ((int(a) % 5 == 0) or (int(a) % 3 ==0)):
      ans = 'DENIED'

  if ans == 'DENIED':
    break

  i += 1

print(ans)