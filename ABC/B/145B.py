N = int(input())
n = N // 2

S = input()
s = S[0:n]

if S == s + s:
  print('Yes')
else:
  print('No')