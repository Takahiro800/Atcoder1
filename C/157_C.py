t = input().split()

N = int(t[0])
M = int(t[1])

y = [input() for i in range(M)]

#import pdb; pdb.set_trace()
#まだ未完了
l = [-1] * N
i = 0
ans = 0

if ('1 0' in y) and (N != 1):
  ans = -1

while i < M:
  y[i] = y[i].split()
  a = int(y[i][0])
  b = int(y[i][1])

  if  l[a-1] == -1:
    l[a-1] = b
  elif l[a-1] != b:
    ans = -1
    break
  i += 1

#import pdb; pdb.set_trace()

while -1 in l:
  k  = l.index(-1)
  l[k] = 0

if (l[0] == 0) and (N != 1):
  l[0] = 1

if ans == 0:
  k = N
  while k > 0:
    ans += l[-k] * (10 ** (k-1))
    k -= 1

print(ans)