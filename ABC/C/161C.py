N, K = map(int, input().split())

t = N % K

while t >= abs(t -K):
  t = abs(K - t)
  if t == 0:
    break

print(t)