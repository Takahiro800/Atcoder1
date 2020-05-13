A = [list(map(int,input().split())) for i in range(3)]
N = int(input())
B = [int(input()) for i in range(N)]

c = [[0]*3 for _ in range(3)]
for b in B:
    for i in range(3):
        for j in range(3):
            if A[i][j] == b:
                c[i][j] = 1


for i in range(3):
  if c[i][0] and c[i][1] and c[i][2]:
    print('Yes')
    exit()
  if c[0][i] and c[1][i] and c[2][i]:
    print('Yes')
    exit()
if c[0][0] and c[1][1] and c[2][2]:
  print('Yes')
  exit()
if c[0][2] and c[1][1] and c[2][0]:
  print('Yes')
  exit()

print('No')