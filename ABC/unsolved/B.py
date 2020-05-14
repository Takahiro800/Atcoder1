A,B,C,D = map(int, input().split())

t = A // D
a = C // B

if A % D == 0:
  t = A // D
else:
  t = (A // D) +1

if C % B == 0:
  a = C // B
else:
  a = (C // B) + 1

if a > t:
  print('No')
else:
  print('Yes')