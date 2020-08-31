S = input()
tmp = []
ans = "yes"

for i in range(len(S)):
  if (S[i] in tmp):
    ans = "no"
    break
  else:
    tmp.append(S[i])
print(ans)