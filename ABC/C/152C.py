N = int(input())
P = list(map(int, input().split()))

ans = 0

for k in range(N):
  # k＝０を避けるため
  s = P[0:k+1]
  min_k = min(s)
  if min_k == P[k]:
    ans += 1

print(ans)