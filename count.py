from typing import Iterable


def avg(d: str | Iterable) -> float:
    if isinstance(d, str):
        d = tuple(map(int, d.split()))
    return sum(d) / len(d)


def print_floats(v: Iterable[float]):
    for i in v:
        print(f'{i:7.2f}', end=' ')
    print()


searches = ('bidirectional', 'manhatten', 'subtask_1col', 'subtask')

for search in searches:
    iters = []
    opens = []
    close = []
    for i in range(1, 8):
        with open(f'{search}_{i}_100.txt') as f:
            lines = f.readlines()
            for line, list_ptr in zip(lines, (iters, opens, close)):
                line = line.split()
                v = avg(tuple(map(int, line[1:])))
                list_ptr.append(v)
    print(search.upper())
    print_floats(iters)
    print_floats(opens)
    print_floats(close)
    print()
