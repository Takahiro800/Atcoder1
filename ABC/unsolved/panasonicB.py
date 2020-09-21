H,W = map(int, input().split())

all = H * W

if (H == 1) or (W == 1):
  print(1)
elif all % 2 == 1:
  print((all + 1)// 2)
else:
  print(all // 2)