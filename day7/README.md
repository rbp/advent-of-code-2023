Part 1 was straightforward. I was surprised to find that 1) there are several different ways to sort a slice in Go (sort.\*, slices.Sort{,Func}, cmp). And 2) that none of them uses a key, they take a "cmp"-like function. I have it in my head to implement Python's PowerSort in Go and make it available as a package. Would be a nice project.

The thing I'm the least satisfied about this code is the naming...

Part 2 wasn't hard per se, but it made the code much less readable... Go doesn't have named or default parameters, so the "obvious" code structure didn't work. I could (should?) perhaps create different types, and refactor the code so that the Joker special case can be handled more gracefully. But I'm almost 10 days behind on the advent, so I'll move on.
