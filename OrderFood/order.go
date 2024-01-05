package OrderFood

import (
	"fmt"
	"strings"
)

func orderItems() {
	printMenu()
	var itemNumber uint
	var noOfItems uint

	for {
		fmt.Println()
		fmt.Println("Enter 0 to exit")
		fmt.Println("Enter the item number")
		fmt.Scan(&itemNumber)
		if itemNumber == 0 {
			break
		}
		var choiceName string
		var itemPrice float64

		for index, item := range menu {
			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}
		}
		fmt.Printf("Enter the number of %v you want to order: ", choiceName)
		fmt.Scanf("%v", &noOfItems)
		if noOfItems == 0 {
			continue
		} else {
			for index := range menu {
				if index+1 == int(itemNumber) {
					customerOrder[choiceName] += noOfItems
					subTotalBill += itemPrice * float64(noOfItems)
					break
				}
			}
			fmt.Printf("\nYou just ordered %v %v which amounts to ₹%v.\n", noOfItems, choiceName, itemPrice*float64(noOfItems))
			orderTillNow()
		}
		fmt.Println()
	}
}

func orderTillNow() {
	fmt.Println("\nYour order till now: ")
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Printf(" %-12s %s\n", "Quantity", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 32))

	for i := range customerOrder {
		fmt.Printf(" %-12v %v\n", customerOrder[i], 1)
	}
	fmt.Printf("%s\n", strings.Repeat("-", 32))
}
