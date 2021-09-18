# TNTN Algorithm

This is a paper/document explaining the TNTN Algorithm.  
Current revision: v1.0  
  
## Why?

Imagine you're working with a programming language whereby you can manually
allocate memory. And for some absurd reason, you have a 'top' number, it's the
length of a big array. You need to eventually convert all numbers from 1...'top'
to string and write it into a buffer, but because you want to have a performant
memory allocation and avoid growing a string, you'd like to know the length of
characters that you will need to allocate just for the numbers.
  
Example:  
f(9) = 9 # "123456789", length of string is 9  
f(10) = 11 # "12345678910", length of string is 11  

I currently don't have any knowledge of any solution(other than actually
creating the string), if there's even a word for this kind of calculation, thus
I will call it the Top Number To Number(TNTN) algorithm as the only input will
be a top number whereby you will expect a number as result. Note: a variation on
this naming could include string or string length.  


## How?

Yes, please read on, that's what's this document is for :)


## Algorithm

Before reading this onwards please note, that every usage of `lenDigits()` will
mean the amount of digits a number has. So 1 will return 1 and 500 will return
3. We will use the `floor(log10(x)) + 1` function for this which is now denoted
as the lenDigits() function. This lenDigits() function is the best function to
accept any finite number. This len function can be implemented in any different
way, it's not limited to `floor(log10(x)) + 1`.

So, a simple solution to this problem will simply iterate over every numbers
from 1...x and add for every `i` number the amount of digits to a value which
will later be used as the return value. Such an approach is very inefficient for
bigger numbers like 1_000_000, whereby it will cost `O(1000000)`.

This document proposes a new approach to this problem which will only cost
`O(lenDigits(x)-1)`. This will mean if we again take the example of 1_000_000 the
current `O(1000000)` will be reduced to `O(6)`.

This approach requires looking at numbers differently. It's a more natural
approach, but it's the key component of why this algorithm can perform this
fast. If we look at any 4 digit number, we generally don't care about the "last"
3 digits(ex. 993 in 4993) for the first part of this algorithm. The 993 is
only useful to the last part to calculate the amount of length of numbers with
4 digits. No matter what 4 digits you have the last 3 digits will only
represent the amount of numbers you have with 4 digits, as the numbers with
1,2,3 digits will be the same and constant across any number with 4 or more
digits. So they always will be at some point between 1...x represent the value
of 9, 90, 900 etc.

So knowing this, we can use a simple recursive calculation to calculate this
amount of numbers with digits for this and always assume it will be 9 at some
point.

This code is written in pseudo-code
```go
    // First part of the calculation.
    digitsInNumber = floor(log10(x)) + 1
    resultValue = 0
    evolvingValue = 9

    for i = 0; i < digitsInNumber-1; i++ {
        resultValue += (evolvingValue * (i + 1))
        // Does a multiplication by * 10 so the "value" of 9 will be changed.
        // As it will move 1 to the left and be able to move the "real value".
        evolvingValue *= 10
    }
    // Second part of the calculation.
    // ....
```

`resultValue` now holds the length you need to write all numbers from 1 to the
highest number with lenDigits(x)-1 digits.

The last part, the amount of numbers you have with `lenDigits(x)` digits. The
algorithm will first calculate `10^(lenDigits(x)-1)-1`, the result will be
subtracted from `x`(result is denoted as y) whereby you're now left with the
amount of numbers that have lenDigits(x) digits. you can efficiently calculate
the length by doing `evolvingValue * y`.

This code is written in pseudo-code:
```go
    // First part of the calculation.
    // ....
    // Second part of the calculation.
    digitsInNumber = floor(log10(x)) + 1
    baseNumber = 10^(digitsInNumber-1) - 1
    amountOfNumbersWithMaxDigit = topNumber - baseNumber
    resultValue += amountOfNumbersWithMaxDigit * digitsInNumber
    // Return resultValue
```

by adding that result to `resultValue` the algorithm has successfully calculated
the length you will need for a string that has 1...x after-each-other.

## Footnote

The reference implementation can be found at [the github](https://github.com/gusted/tntn).  
Please check the `IdiomaticTNTN` function.

This paper is released under the [CC BY 4.0 License](https://creativecommons.org/licenses/by/4.0/). 
