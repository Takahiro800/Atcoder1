s = input().split()

N = int(s[0])
R = int(s[1])

if N < 10:
  ans = R + 100 *(10 - N)
else:
  ans = R

print(ans)