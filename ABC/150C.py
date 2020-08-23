import math

N = int(input())
P = list(map(int, input().split()))
p = sorted(P)
Q = list(map(int, input().split()))
q = sorted(Q)
a = 0
b = 0

import itertools
i = 1
for ptn in itertools.permutations(range(1, N+1)):
  if list(ptn)==P:
    a = i
  if list(ptn)==Q:
    b = i
  i += 1
print(abs(a-b))