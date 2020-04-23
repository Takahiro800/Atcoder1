import math
N = int(input())
ans = math.ceil(N // (1.08))

while math.floor(ans*1.08) < N:
  ans += 1
  if math.floor(ans*1.08) == N:
    print(ans)
  elif math.floor(ans*1.08) > N:
    print(':(')
