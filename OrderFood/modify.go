package main

import "fmt"

func modifyOrder() {
	for {
		var isOrderOk string
		fmt.Println("Do you want to change your order")
		fmt.Scan(&isOrderOk)
		if isOrderOk != "y" {
			return
		}
		var serialNo uint
		var modifyType uint

		fmt.Println("Please enter the respective no. to proceed further: ")
		fmt.Println("Press '1' to update item quantity.")
		fmt.Println("Press '2' to delete an item from the order list.")
		fmt.Println("Press '3' to add item(s) in the order list.")
		fmt.Scan(&modifyType)

		switch modifyType {
		case 1:
			printMenu()
			fmt.Println("Please select")
			fmt.Scan(&serialNo)
			updateQuantity(serialNo)

		case 2:
			printMenu()
			fmt.Println("Please select")
			fmt.Scan(&serialNo)
			delFromOrder(serialNo)

		case 3:
			insertIntoOrder()

		default:
			return
		}
		displayGeneratingBill()
		generateBill()
	}
}

func updateQuantity(serialNo uint) {
	var newQuantity uint

	for _, targteItem := range menu {
		if serialNo == targteItem.itemNo {
			itemName := targteItem.itemName
			oldQuantity := customerOrder[itemName]

			fmt.Printf("Current quantity of %v is %v.\n", itemName, oldQuantity)
			fmt.Printf("Now, enter the updated quantity of %v to be ordered: \n", itemName)
			fmt.Scan(&newQuantity)

			if newQuantity == 0 {
				delFromOrder(serialNo)
				return
			}
			fmt.Printf("")

			customerOrder[targteItem.itemName] = newQuantity
			fmt.Printf("Updated the quantity of %v from %v to %v.\n", itemName, oldQuantity, newQuantity)

			subTotalBill -= float64(oldQuantity) * targteItem.itemPrice
			subTotalBill += float64(newQuantity) * targteItem.itemPrice

			break
		}
	}
}

func delFromOrder(serialNo uint) {
	for _, targetItem := range menu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName
			subTotalBill -= float64(customerOrder[itemName]) * targetItem.itemPrice
			delete(customerOrder, itemName)
			fmt.Printf("Deleted %v from the order list.\n", itemName)
			break
		}
	}
}

func insertIntoOrder() {
	orderItems()
}
