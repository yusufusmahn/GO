package main

import "fmt"

func main() {
    for {
        fmt.Print("Enter account number (-1 to quit): ")
        var accountNumber int
        fmt.Scan(&accountNumber)

        if accountNumber == -1 {
            break
        }

        fmt.Print("Enter beginning balance: ")
        var beginBalance int
        fmt.Scan(&beginBalance)

        fmt.Print("Enter total charges: ")
        var charges int
        fmt.Scan(&charges)

        fmt.Print("Enter total credits: ")
        var credits int
        fmt.Scan(&credits)

        fmt.Print("Enter credit limit: ")
        var creditLimit int
        fmt.Scan(&creditLimit)

        if beginBalance < 0 || charges < 0 || credits < 0 || creditLimit < 0 {
            fmt.Println("Inputs cannot be negative. Try again.")
            continue
        }

        newBalance := beginBalance + charges - credits
        fmt.Printf("Account %d: New balance = %d\n", accountNumber, newBalance)

        if newBalance > creditLimit {
            fmt.Println("Credit limit exceeded")
        }
    }

    fmt.Println("Program ended.")
}
