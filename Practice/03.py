N = int(input())
A = input().split()

i = 0
ans = 0

even_count = N

while even_count == N:
  even_count = 0
  i = 0

  while i < N:
    a = int(A[i])
    if a % 2 == 0:
      even_count += 1
      A[i] = a // 2
      i += 1
    else:
      break

  if even_count == N:
    ans += 1
  else:
    break

print(ans)