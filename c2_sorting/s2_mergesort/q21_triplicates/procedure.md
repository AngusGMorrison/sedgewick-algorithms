Sort each array so that all are in alphabetical order.
Let i, j, k = list1[0], list2[0], list3[0].
For i, j, k < N
  Keep track of which list is the leader (it's pointer is furthest ahead in the alphabet).
  Move the second pointer until it equals the leader or becomes the new leader.
    * If the second pointer equals the name pointed to by leader, move the third pointer until it too equals the leader or becomes the new leader.
    * If the third pointer also equals the leader's name, return the name.
    * If a new leader arises, repeat. 