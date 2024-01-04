package OrderFood

import (
	"fmt"
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


func printOrderData(){
	for key := range customerOrder {
		for _, element := range menu {
			if key == element.itemName{
				totalCostOfItem := float64(customerOrder[key]) * element.itemPrice
				fmt.Printf(" %-30s %-20.2f %-20v %-20.2f\n", key, element.itemPrice, customerOrder[key], totalCostOfItem)
			}
		}
	}
	fmt.Println()
}

