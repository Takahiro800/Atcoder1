S = input()

for i in range(26):
  a = chr(i + ord('a'))
  if a not in S:
    print(a)
    exit()
print('None')