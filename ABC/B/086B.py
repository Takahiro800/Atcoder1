import math
a,b = map(str, input().split())

product = int(a + b)
s = math.sqrt(product)

if s.is_integer():
  print('Yes')
else:
  print('No')
