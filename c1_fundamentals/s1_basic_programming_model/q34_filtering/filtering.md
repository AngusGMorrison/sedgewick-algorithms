# 1.1.34 Filtering

Which of the following require saving all the values from standard input (in an array, say), and
which could be implemented as a filter using only a fixed number of variables and arrays of fixed
size (not dependent on N)? For each, the input comes from standard input and consists of N real
numbers between 0 and 1.

1. Print the maximum and minimum numbers.
  * Filter. As each number is processed, it is checked against the current minimum and maximum and
    replaces the respective variable if it is smaller or greater than the current value, as
    appropriate.
2. Print the median of the numbers.
  * Save. In order to know which number or numbers represent the middle of the sample, we must know
    the size of the sample. The median therefore can't be calculated until we've captured all
    values.
3. Print the kth smallest value, for k less than 100.
  * Filter. Finding the kth smallest value requires k variables to store and update the
    first-smallest, second-smallest, ..., kth-smallest values as the input is processed.
4. The sum of the squares of the numbers.
  * Filter. We require only one variable to hold the sum. Each number is squared and added to the
    sum as it is received.
5. Print the average of the n numbers.
  * Filter. We can use two variables: one to hold the running total, and one to count the number of
    entries.
6. Print the percentage of numbers greater than the average.
  * Save. To calculate the average requires only two values, but we must have saved each value while
    calculating the average to determine how many are greater than the average value.
7. Print the N numbers in increasing order.
  * Save. The total ordering of the numbers can only be determined once all are known.
8. Print the N numbers in random order.
  * Save. For each number to have an equal probability of appearing in any position, all must be
    known before printing begins.
