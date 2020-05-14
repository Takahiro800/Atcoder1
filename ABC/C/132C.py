N = int(input())

d = list(map(int, input().split()))
list.sort(d)
n = N // 2

a = d[n]
b = d[n-1]

print(a - b)