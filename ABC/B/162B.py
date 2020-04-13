N = int(input())

Sum = (N * (N+1)) / 2

a = N // 3
b = N // 5
c = N // 15

A = a * (3 + 3*a) / 2
B = b * (5 + 5*b) / 2
C = (c * 15 * (c+1)) / 2

ans = int(Sum + C - (A+B))
print(ans)