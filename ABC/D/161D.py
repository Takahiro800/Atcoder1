K = int(input())
ans = []
i = 1

while len(ans) < K :
  if i < 10:
    ans = ans + [i]
    i += 1

  else:
    s = (str(i))[-1]
    t = int(s) - 1

    if int(s) == 0:
      s
