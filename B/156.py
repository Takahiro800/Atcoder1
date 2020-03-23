s = input().split()
N = int(s[0])
K = int(s[1])

i = 0

while N > K ** i:
  i += 1

ans = i
print(ans)