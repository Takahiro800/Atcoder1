N = int(input())
lis = input().split()

count_list = [0] * (N + 1)

for k in lis:
  count_list[int(k)] += 1

Sum = 0
for k in count_list:
  Sum += ( k * (k-1)) // 2



for i in lis:
  count = count_list[int(i)]
  s = int(count) - 1
  ans = Sum - s
  print(ans)