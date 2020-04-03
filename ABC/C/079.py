S = input()

for a in '+-':
  for b in '+-':
    for c in '+-':
      ans = S[0] + a + S[1] + b + S[2] + c + S[3]
      if eval(ans) == 7:
        print(ans + '=7')
        exit()