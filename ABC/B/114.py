S = input()
ans = 99999999999999999999

for i in range(len(S) - 2):
  s = int(S[i:i+3])
  ans = min(ans, abs(s - 753))

print(ans)