import collections

N = int(input())
lis = input().split()
i = 0

print(lis)

while i < N:
  #lis = lis.pop(i)
  count_dict = collections.Counter(lis)

  print(count_dict.most_common(N))

  M = count_dict.most_common(N)
  ans = 0

  for k in count_dict.values():
    x = (k * (k-1)) // 2
    ans += x
  print(ans)
