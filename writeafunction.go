to check wethere given year is a leap year task 1
---------------------------------------------
package main

import "fmt"

func main() {

  var input int
  fmt.Println("Enter a year:")
  fmt.Scanf("%d", &input)
  if (input % 4 == 0) {
    if (input % 100 == 0) {
      if (input % 400 == 0) {
    fmt.Println(input , " its a leap year")
    } else {
    fmt.Println(input , " its not a leap year")
  }
  } else {
    fmt.Println(input , " its a leap year")
}
} else {
  fmt.Println(input , " its not a leap year")
}
}

lets consider this year as an example
2019
after executing the program we will get 
1.enter a year

enter it

2019

shows it is not a leap year,bcz not divided by 4.
