A,B,C,D,E = map(int, input().split())
sum = A+B+C+D+E

m = min(A+D, B+C)
ans = sum - m

print(ans)
