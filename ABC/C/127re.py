N,M = map(int, input().split())
LR = [tuple(map(int, input().split())) for i in range(M)]
maxl = 0
minr = N+1

for l,r in LR:
  maxl = max(maxl, l)
  minr = min(minr, r)
print(max(0, minr-maxl+1))