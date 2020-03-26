m = input().split()

A = int(m[0])
B = int(m[1])
C = int(m[2])

if A == B:
  if A == C:
    ans = 'No'
  else:
    ans = 'Yes'
else:
  if A == C:
    ans = 'Yes'
  elif B == C:
    ans = 'Yes'
  else:
    ans = 'No'

print(ans)