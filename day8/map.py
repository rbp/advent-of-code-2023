from collections import defaultdict
from itertools import cycle
from functools import reduce
import sys


# gcd/lcm code taken from https://stackoverflow.com/questions/147515/least-common-multiple-for-3-or-more-numbers
def gcd(a, b):
    """Return greatest common divisor using Euclid's Algorithm."""
    while b:      
        a, b = b, a % b
    return a

def lcm(a, b):
    """Return lowest common multiple."""
    return a * b // gcd(a, b)

def lcmm(*args):
    """Return lcm of args."""   
    return reduce(lcm, args)

def main():
    lines = open(sys.argv[1], 'r').readlines()
    directions = lines[0].strip()
    nodes = {}
    for line in lines[2:]:
        n, rest = line.strip().split(' = ')
        l, r = rest.lstrip('(').rstrip(')').split(', ')
        nodes[n] = (l, r)
    steps = 0

    node = 'AAA'
    turns = cycle(directions)
    while node != 'ZZZ':
        if node not in nodes:
            steps = 0
            break
        node = nodes[node][0 if next(turns) == 'L' else 1]
        steps += 1
    print(f'Part 1: {steps}')

    next_nodes = [n for n in nodes if n.endswith('A')]  
    steps_per_node = defaultdict(list)
    for starting_node in next_nodes:
        seen = set()
        turns = cycle(range(len(directions)))
        next_turn_i = next(turns)
        node = starting_node
        steps = 0
        while (node, next_turn_i) not in seen:
            seen.add((node, next_turn_i))
            next_turn = directions[next_turn_i]
            node = nodes[node][0 if next_turn == 'L' else 1]
            steps += 1
            if node.endswith('Z'):
                steps_per_node[starting_node].append(steps)
            next_turn_i = next(turns)
        print(f"Steps to ..Z for {starting_node}: {steps_per_node[starting_node]}")
    
    # ... A little experimentation later...
    # Turns out each starting node only sees one ..Z before a cycle?

    # # Now, to find the least common multiplier amongst the nodes
    steps = tuple(v[0] for v in steps_per_node.values())
    print(f"Part 2: {lcmm(*steps)}")

if __name__ == '__main__':
    main()
