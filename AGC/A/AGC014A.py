A,B,C = map(int, input().split())

ans = 0

while A%2 == B%2 == C%2 == 0:
  ans += 1
  A = A // 2
  B = B // 2
  C = C // 2
  #print(A,B,C)
  if A == B == C:
    ans = -1
    break
  elif A%2 == 1 or B%2 == 1 or C%2 == 1:
    a = A
    b = B
    A = B + C
    B = C + a
    C = a + b

print(ans)