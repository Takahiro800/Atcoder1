N = int(input())

list = input().split()
list_int = [int(s) for s in list]
list_int.sort(reverse=True)

A = 0
B = 0

for x in range(N):
  if x % 2 == 0:
    A += list_int[x]
  else:
    B += list_int[x]

ans = A - B
print(ans)
