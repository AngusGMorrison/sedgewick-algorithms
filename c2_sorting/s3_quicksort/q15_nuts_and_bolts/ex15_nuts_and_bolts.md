1. Create two "arrays", one of nuts and one of bolts. Assume these arrays are shuffled.
2. Using the first bolt as a pivot, partition the array of nuts, putting all nuts smaller than the
   bolt on the left, and all nuts greater than the bolt on the right. At some point in this process,
   we match a nut to the pivot bolt.
3. Use the nut matching the pivot bolt to partition the bolts array. We now know that all bolts on
   the left of this pivot nut are smaller than the original pivot bolt, and all nuts on the right
   are larger than the original pivot bolt.
4. Recursively repeat the process for the two "left-hand" subarrays (one of nuts and one of bolts),
   and the two "right-hand" subarrays, until all nuts have a bolt.