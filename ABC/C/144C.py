import math
N = int(input())
ans = 0

a = 1
while a <= math.sqrt(N):
  if N % a == 0:
    b = N // a
    len = (a + b) - 2

    if (ans == 0) or (ans > len):
      ans = len
  a += 1

print(ans)

#まだ途中エラーが出ているので明日取り組む
