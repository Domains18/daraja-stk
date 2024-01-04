package OrderFood

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func displayGeneratingBill() {
	fmt.Println()
	billDisplayText := "--------------------------------Generating Bill------------------------------------------"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 15)
	}
}

func generateBill() {
	fmt.Println()
	fmt.Printf("%s+\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s %-20s\n", "Item Name", "Price($)", "Quantity", "Total Price($)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	printOrderData()

	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total: $" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("%.2f\n", subTotalBill)
}

func printOrderData() {
	for key := range customerOrder {
		for _, element := range menu {
			if key == element.itemName {
				totalCostOfItem := float64(customerOrder[key]) * element.itemPrice
				fmt.Printf(" %-30s %-20.2f %-20v %-20.2f\n", key, element.itemPrice, customerOrder[key], totalCostOfItem)
			}
		}
	}
	fmt.Println()
}

func printFinalBill() {
	for _, element := range "Here is your final Bill: " {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println()
	fmt.Printf("\n%52s\n", "Hello There")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s\n", strings.Repeat("*", 91))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%86s\n", "This is my hotel")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%50s\n", "Tel: your number")
	fmt.Printf("%60s\n\n", "Email: awesomedevs@gmail.com")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s", strings.Repeat("-", 42))
	fmt.Printf("%s", "INVOICE")
	fmt.Printf("%s\n", strings.Repeat("-", 42))
	time.Sleep(time.Millisecond * 200)

	rand.Seed(time.Now().Unix()) //necessary to produce random integers
	fmt.Printf(" Ticket No: %d\n", rand.Intn(550)+1)

	fmt.Printf(" Date: %v\n", time.Now().Local().Format("06-Jan-02")) //display date
	fmt.Printf(" Time: %v", time.Now().Local().Format("3:4 PM"))      //display time
	time.Sleep(time.Millisecond * 200)

	generateBill() //prints details of the bill,like, item name, price, quantity and total price and sub total amount.

	tax := 18 * subTotalBill / (100)
	grandTotal := subTotalBill + tax

	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "GST", tax) //display tax amount
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "Grand Total", grandTotal) //display final bill
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
}
