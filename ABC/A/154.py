M = [input() for i in range(3)]
M[0] = M[0].split()
s = str(M[0][0])
t = str(M[0][1])

M[1] = M[1].split()
a = int(M[1][0])
b = int(M[1][1])

U = str(M[2])


if s == U:
  a -= 1
else:
  b -= 1

print(str(a) + ' ' + str(b))
