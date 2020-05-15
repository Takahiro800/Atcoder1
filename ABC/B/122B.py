S = input()
N = len(S)
a,b = 0,0
l = [0]

for i in range(N):
  if S[i] == 'A' or S[i] == 'C' or S[i] == 'G' or S[i] == 'T':
  # S[i] in 'ACGT'
    a += 1
    b = 1
  else:
    if b:
      l.append(a)
      b = 0
      a = 0
else:
  if a:
    l.append(a)
  print(max(l))