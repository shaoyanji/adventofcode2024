# Notes for advent of code

## Day 1

Chief historian is missing. First 50 places to check to save Christmas.

The list of numbers is long and could be a cipher.

- first, find the difference per pair
- find the smallest number in the left list and pair with the smallest number in the right list.

- use sort and then a diff
- just use a diff for quick total distance
- absolute value is needed at each pair comparison

Part 2 makesup something called the similarity score which is the sum of left numbers multiplied by the instances paired with the right numbers. initial approach was to do a nested forloop but splitting it into two functions made sense. Bug was += is equivalent to x = x + (whatever).

```go
go run main.go
```


