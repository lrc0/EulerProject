```
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

最长柯拉茨序列

对一个正整数，当n为奇数，乘3再加1；n为偶数，除以2.

对13按以上规则操作，我们得到序列

13-40-20-10-5-16-8-4-2-1

以1为序列的终点，序列包含10个数。我们猜想每个起始数最终会回到1.

求序列长度最长的1000000以下的起始数。

（提示：序列中的数可能会大于1000000）

```