Part 1 was very straightforward. I think the "catch" would repeating the pattern of the directions, which would have been trivial with Python's itertools. But it was very simple to replicate it.

Now, part 2...
Traversing all nodes at once seemed trivial. But it never ended! I looked for bugs, tried to implement it differently, but nothing. I pondered the problem for a while and realised that rather than walk all paths until all of them found an end node, I could find how many steps each path took until it found an end node (technically, all end nodes found until there was a cycle in the path), and then calculate the least common multiplier between them. But I still wasn't arriving at a result.

Then, just to be sure it wasn't a silly bug or me not knowing enough Go, I tried in Python. First thing I was impressed with is how much shorter and straightforward the solution to part 1 was! But part 2 still wasn't coming. But part 2 was _still_ taking forever. Eventually I realised that my method of calculating the lcm was too slow. I google a bit, found a very quick one, and voila!

At the point, I considered rewriting the solution in Go. But I've had enough of this problem, time to move on.
