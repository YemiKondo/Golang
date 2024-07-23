package main

import (
    "fmt"
    "strconv"
)

func isPalindrome(num int) bool {
    str := strconv.Itoa(num)
    for i := 0; i < len(str)/2; i++ {
        if str[i] != str[len(str)-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    if isPalindrome(num) {
        fmt.Println(num, "is a palindrome")
    } else {
        fmt.Println(num, "is not a palindrome")
    }
}