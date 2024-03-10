# Part 2 Solution

# Basic Way

## First Thought, doesn't work though

The idea is to reuse the function from part 1. To have that, for every number word, i will replace the first occurrence of it, from the begin, and from the end.
Cost?:

- Go builtin Replace Func cost? - For now call it O(R)
- O(L) is the time to iterate the string
- O(N)x O(R)xO(L)

Pretty huge, not so happy

### Why it doesn't work

- Replace only first occurrence:
  - seven5seven -> doesn't work
- The string can overlap, so when iterate through the word number array, the order of word can effect the result
  - twone, with word number array order like {one,two} => tw1, but the right result is 2ne

## Solution that work

- instead of replacing first occurrence, replace every. And instead of replacing like one -> 1, replace one -> one1one, so the order of replace doesn't effect the string (twone -> twone1one, and the next iteration it can still replace two)

# Advance way that I found on the internet
