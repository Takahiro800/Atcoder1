m = input().split()
N = int(m[0])
A = int(m[1])
B = int(m[2])

sum = A + B
quo = N // sum
rem = N % sum

if A >= rem:
  ans = quo * A + rem
else:
  ans = (quo + 1) * A

print(ans)