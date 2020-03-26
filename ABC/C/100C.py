N = int(input())
list = input().split()

i = 0
ans = 0

for a in range(N):
  while int(list[a]) % 2 == 0:
    ans += 1
    list[a] = int(list[a]) // 2

print(ans)