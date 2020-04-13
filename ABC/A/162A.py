N = int(input())

a = N // 100
b = (N - 100*a) // 10
c = N % 10

if (a==7) or (b==7) or (c==7):
  print('Yes')
else:
  print('No')
