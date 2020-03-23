S = input()
S = list(S)

ans = "Yes"
N = len(S)

i = 0
while i < N:
  k = N - i
#  print(S[i]),
# print(S[k-1])
  if S[i] != S[k -1]:
    ans = ' No'
  i += 1

i = 0
l = (N - 1) // 2

while i < l:
  k = l -i

  if S[i] != S[k - 1]:
    ans = 'No'
  i += 1

i = (N + 3) // 2
while i < N:
  m = N - i
  if S[i] != S[m -1]:
    ans = 'No'
  i += 1

print(ans)
