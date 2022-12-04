for left, ok = q1.Pop(); ok {
  if right, ok = q2.Pop(); ok {
    output.Push(min(left, right))
    push max(left, right) back onto its q (could also use peeking to reduce number of ops)
  } else {
    output.Push(left)
  }
}

for right, ok = q2.Pop(); ok {
  output.Push(right)
}