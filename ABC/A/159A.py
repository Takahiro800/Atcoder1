t = input().split()

N = int(t[0])
M = int(t[1])

A = (N * (N -1)) // 2
B = (M * (M -1)) // 2


ans = A + B
print(ans)