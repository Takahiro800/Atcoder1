N = int(input())
A = list(map(int, input().split()))
B = [0] * (N+1)


for i in range(N):
  B[A[i]] = i+1
else:
  B.pop(0)
  B = [str(b) for b in B]


B = " ".join(B)
print(B)