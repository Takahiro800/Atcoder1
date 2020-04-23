N,A,B = map(int, input().split())
S = input()

for i in range(N):
  if S[i] == 'c':
    print('No')
  elif S[i] == 'a' and (A+B) > 0:
    print('Yes')
    A -= 1
  elif S[i] == 'b' and (A+B) > 0 and B > 0:
    print('Yes')
    B -= 1
  else:
    print('No')