import collections

with open("input.txt") as file:
    data = [int(x) for x in file.read().splitlines()]

idxs = collections.deque(range(len(data)))

# "Mix" the array
for i in range(0, len(idxs)):
    idx = idxs.index(i)
    v = data[i]
    idxs.rotate(-idx)
    idxs.popleft()
    idxs.rotate(-v)
    idxs.appendleft(i)

data = [data[i] for i in idxs]
indexZero = data.index(0)
result = [data[(indexZero+1000) % len(data)],
          data[(indexZero+2000) % len(data)],
          data[(indexZero+3000) % len(data)]]
print(sum(result))
