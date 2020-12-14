import functools as ft
L = int(input())

def permutation(n, k):
    if k > n:
        return 0
    elif 0 < k <= n:
        return ft.reduce(lambda x, y:x * y, [n - v for v in range(k)])
    else:
        return 1


def factorial(n):
    return permutation(n, n - 1)


def combination(n, k):
    return int(permutation(n, k) / factorial(k))

ans = combination(L-1, 11)
print(ans)