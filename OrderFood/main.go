package OrderFood

import "fmt"

var subTotalBill float64

var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Println("what is your firstName")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill()
	generateBill()
	modifyOrder()
	printFinalBill()
	bye(customerName)
}
