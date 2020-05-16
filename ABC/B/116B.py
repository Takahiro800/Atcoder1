s = int(input())

A = []
ans = 1

def seq(a):
  if a % 2 == 0:
    a = a // 2
  else:
    a = 3 * a + 1
  return a

while s not in A:
  A.append(s)
  s = seq(s)
  ans += 1
else:
  print(ans)