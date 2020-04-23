S = input()
len = len(S)

n = len // 2
ans = 0

for i in range(n):
  k = (i+1) * (-1)
  if S[i] != S[k]:
    ans += 1

print(ans)