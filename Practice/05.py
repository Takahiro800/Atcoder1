m = input().split()
N = int(m[0])
A = int(m[1])
B = int(m[2])
sum = 0

for n in range(N + 1):
  y = 0
  s = n
  while s > 0:
    y += s % 10
    s = s // 10

  if (y >= A) and (y <= B):
    sum += n

print(sum)
