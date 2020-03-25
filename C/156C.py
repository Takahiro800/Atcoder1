N = int(input())

X = input().split()
#import pdb; pdb.set_trace()
i = 0
sum = 0

while i < N:
  x = int(X[i])
  sum += x
  i += 1

P = round(sum / N)

dis = 0
i = 0


while i < N:
  z = int(X[i])
  dis += (P - z) ** 2
  i += 1

print(dis)