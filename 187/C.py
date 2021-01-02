N = int(input())

S = [""]*N

for i in range(N):
  x = input()
  if "!"+x in S:
    print(x)
    exit()
  elif x[1:] in S:
    print(x[1:])
    exit()
  S[i] = x

print("satisfiable")