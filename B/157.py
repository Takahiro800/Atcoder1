m = [input() for i in range(3)]
m_1 = m[0].split()
m_2 = m[1].split()
m_3 = m[2].split()

m_1.extend(m_2)
m_1.extend(m_3)
m = m_1.insert(0, 0)
s = [0] * 10

N = int(input())
bin = [input() for i in range(N)]

for item in bin:
  if m.count(item) == 1:
