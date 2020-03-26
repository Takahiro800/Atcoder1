M = input().split()

a = int(M[0])
b = int(M[1])

x = 10 * b
ans = -1


while (x < 10 * (b + 1)) and (ans == -1):
  if int(x * (1 / 10)) != b:
    x += 1
  else:
    ans = x

ans = -1

if x < (a + 1) * (25 / 2):
  while (x < (a + 1) * (25 / 2)) and (ans == -1 ):
    if int(x * (2 / 25)) == a:
      ans = x
    else:
      x += 1

  if int(x * (1 / 10)) == b:
    ans = x
  else:
    ans = -1

print(ans)
