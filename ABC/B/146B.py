N = int(input())
S = input()

ans = ''
for c in S:
  i = ord(c) - ord('A')
  i = (i+N) % 26
  ans += chr(i + ord('A'))

print(ans)