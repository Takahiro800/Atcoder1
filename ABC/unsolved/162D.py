N = int(input())
S = input()

r = S.count('R')
g = S.count('G')
b = S.count('B')

x = 0
y = 1
ans = 0

while x < N-2:
  y = 0
  k = N - (x+y)

  while y < N-1:
    if x> y or y > k:
      continue
    elif y - x == k -y:
      continue
    elif S[x] == S[y] or S[y] == S[k] or S[k] == S[x]:
      continue
    else:
      ans += 1

    y += 1

print(ans)