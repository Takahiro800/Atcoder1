N = int(input())
K = int(input())
X = list(map(int, input().split()))

line = K // 2
ans = 0

for x in X:
	if x <= line:
		ans += x
	else:
		ans += K - line
ans *= 2
print(ans)