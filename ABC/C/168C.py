from math import pi,cos, sqrt

A,B,H,M = map(int, input().split())

time = H * 60 + M
a = time % 720
b = time % 60

a = a * 360 / 720
c = cos((a-b)*2*pi)

L = A**2 + B**2 - 2*A*B*c
ans = sqrt(L)
print((a-b)*2*pi)
print(ans)