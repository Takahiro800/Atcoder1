K = int(input())
S = input()
s = len(S)

if s <= K:
  print(S)
else:
  l = S[:K] + '...'
  print(l)