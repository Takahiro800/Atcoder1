N = int(input())

""" if N >= 64:
  print(64)
elif N >=32:
  print(32)
elif N >= 16:
  print(16)
elif N >= 8:
  print(8)
elif N >= 4:
  print(4)
elif N >= 2:
  print(2)
else:
  print(1) """

k = len(bin(N)) - 3
print(2**k)