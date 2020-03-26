N = int(input())

bucket_sort = [0] * 101

for i in range(N):
  d = int(input())

  if bucket_sort[d] != 1:
    bucket_sort[d] = 1

ans = 0
#import pdb; pdb.set_trace()

for k in bucket_sort:
  ans +=  k

print(ans)