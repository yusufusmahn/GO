package main

import "fmt"

func main() {
    totalMiles := 0
    totalGallons := 0

    for {
        fmt.Print("Enter miles driven (-1 to quit): ")
        var miles int
        fmt.Scan(&miles)

        if miles == -1 {
            break
        }

        fmt.Print("Enter gallons used: ")
        var gallons int
        fmt.Scan(&gallons)

        if gallons <= 0 {
            fmt.Println("Gallons must be positive. Try again.")
            continue
        }

        mpg := float64(miles) / float64(gallons)
        fmt.Printf("MPG for this trip: %.2f\n", mpg)

        totalMiles += miles
        totalGallons += gallons

        if totalGallons > 0 {
            combinedMPG := float64(totalMiles) / float64(totalGallons)
            fmt.Printf("Combined MPG for all trips: %.2f\n", combinedMPG)
        }
    }

    if totalGallons > 0 {
        combinedMPG := float64(totalMiles) / float64(totalGallons)
        fmt.Printf("Final combined MPG: %.2f\n", combinedMPG)
    } else {
        fmt.Println("No trips entered.")
    }
}