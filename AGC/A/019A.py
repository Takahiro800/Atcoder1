m = input().split()
Q = int(m[0])
H = int(m[1])
S = int(m[2])
D = int(m[3])

ans = 0
N = int(input())

for d in range((N // 2) + 1 ):
  s_set = N - (2*d)
  for s in range(s_set + 1):
    h_set = N - (2*d + 1*s)
    for h in range(h_set * 2 +1):
      q = int( 4 * (N - (2*d + 1 * s + 0.5 * h)))
      money = d*D + s*S + h*H + q*Q

      if (ans == 0) or (ans > money):
        ans = money

print(ans)