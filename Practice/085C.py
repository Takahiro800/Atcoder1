m = input().split()
N = int(m[0])
Y = int(m[1])
ans = [0] * 3

for a in range(N+1):
  for b in range(N - a + 1):
    c = N - (a + b)
    money = (a * 10000) + (b * 5000) + (c * 1000)

    if money == Y:
      ans[0] = a
      ans[1] = b
      ans[2] = c
      break
  else:
    continue

  break

if ans == [0, 0, 0]:
  ans = [-1, -1, -1]

x = ans[0]
y = ans[1]
z = ans[2]

print(str(x) + ' ' + str(y) + ' ' + str(z))
