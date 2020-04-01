import pdb

N = int(input())
i = 0
list = [0] * N

while i < N:
  list[i] = input()
  i += 1

list.sort()
k = 1
count_list = [0]*N
count_list[0] = 1
ans_list = [0] * N
ans_list[0] = list[0]
l = 0



while k < N:
  if list[k] != list[k-1]:
    l +=1
    ans_list[l] = list[k]

  count_list[l] += 1
  k += 1

s = max(count_list)

for i in range(N):
  if count_list[i] == s:
    print(ans_list[i])