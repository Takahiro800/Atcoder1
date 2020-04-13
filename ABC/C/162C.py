import math
from functools import reduce

def gcd(*numbers):
    return reduce(math.gcd, numbers)

def gcd_list(numbers):
    return reduce(math.gcd, numbers)

K = int(input())

a = 1
b = 1
c= 1
ans = 0

while a <= K:
  b = 1
  while b <= K:
    c = 1
    while c <=K:
      if gcd(a, b) == 1:
        ans += 1
      else:
        ans += gcd(a, b, c)
      c += 1
    b += 1
  a += 1
print(ans)