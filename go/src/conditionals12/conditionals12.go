package main

import "fmt"

func main() {
  if lessonLearned := 2; lessonLearned == 5 {
    fmt.Println("Great job! You can continue on to the next exercise.")
  } else if lessonLearned == 4 {
    fmt.Println("Good job! Give it a try to perfection anyway")
  } else if lessonLearned == 3 {
    fmt.Println("Saved by the bell but we believe that you can do better")
  } else {
    fmt.Println("I think that you should explore a career as a Youtuber.")
  }
  // Create more conditions below!
}