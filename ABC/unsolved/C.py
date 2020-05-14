N = int(input())
ans = 0
l = [] * N


for k in range(N):
  t = input()
  if t not in l:
    ans += 1
    l.append(t)

print(ans)