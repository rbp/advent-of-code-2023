Part 1 was straightforward, but I was afraid of searching times, so I implemented (part of) a binary search tree. I actually started making it self-balancing, but it was getting complex, I was on an armchair with distractions all around me, so I thought I'd try to my luck and risk getting the worst O(n) case. But in the end it was very quick.

I didn't have to do much to get part 2. Just parse ranges instead of seed numbers, fix a couple of minor bugs and iterate through all the seeds.

I'd heard that day 5 part 2 was tricky, but my naive solution (after the arguably complex one for part 1) was fast enough, about 3 minutes (I'd heard of 35 minutes to a couple of hours). I tried adding goroutines, but was surprised to see it didn't make a difference. I'm still trying to figure out why.

Incidentally, I have a coding assistant (Sourcegraph's Cody) enabled, and though I'm usually rather careful about accepting any of its suggestions, this time I allowed it to complete a few trivial for-loop lines without paying much attention. That cost me _2_ off-by-ones and a lot of time debugging.
