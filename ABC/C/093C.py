A,B,C = map(int,input().split())
list = [A, B, C]
list.sort(reverse = True)

a = list[0]
b = list[1]
c = list[2]
ans = (a -b) + (b-c)//2

if (b - c) % 2 == 0:
  print(ans)
else:
  ans += 2
  print(ans)
