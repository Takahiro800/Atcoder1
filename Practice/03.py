m = input()
m = list(m)

ans = 0
k = 0

while k < 3:
  if m[k] == '1':
    ans += 1

  k += 1

print(ans)