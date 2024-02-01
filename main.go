package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var subTotalBill float64
var customerOrder = make(map[string]uint, 0)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Hello ", customerName)
	fmt.Printf("%72s\n", " Welcome to Little Lemon Restaurant! ")
	fmt.Println()
}

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("Thank you %v for visiting Little Lemon!\n", customerName)
}

var menu = []Menu{
	{1, "Bruschetta", 14.99},
	{2, "Greek Salad", 12.99},
	{3, "Lemon Dessert", 10.99},
	{4, "Humus", 7.95},
	{5, "Falafal", 8.50},
	{6, "Shawarma", 10.39},
	{7, "Moussaka", 15.99},
	{8, "Kunafa", 6.95},
	{9, "Baklava", 7.95},
	{10, "Revani", 5.99},
}

func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "Serial No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}

func orderItems() {
	printMenu()
	var itemNumber uint
	var noOfPlates uint
	for {
		fmt.Println()
		fmt.Println("Enter '0' to EXIT")
		fmt.Print("Enter the Serial No. of the item to order: ")
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
		fmt.Printf("How many %v do you want: ", choiceName)
		fmt.Scan(&noOfPlates)
		if noOfPlates == 0 {
			continue
		} else {
			for index := range menu {
				if index+1 == int(itemNumber) {
					customerOrder[choiceName] += noOfPlates
					subTotalBill += itemPrice * float64(noOfPlates)
					break
				}
			}
			fmt.Printf("\nYou just ordered %v %v which amounts to ₹%v.\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
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
		fmt.Printf(" %-12v %v\n", customerOrder[i], i)
	}
	fmt.Printf("%s\n", strings.Repeat("-", 32))
}

func displayGeneratingBill() {
	fmt.Println()
	billDisplayText := "************************************* Generating Bill *************************************"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 15)
	}
}

func generateBill() {
	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s %-20s\n", "Item Name", "Price($)", "Quantity", "Total Price($)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	printOrderData()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total (excluding Tax): $" {
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
	for _, element := range "Here is your final bill:" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println()
	fmt.Printf("\n%52s\n", "Little Lemon Restaurant")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s\n", strings.Repeat("*", 91))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%70s\n", "431 El Camino Real, Santa Clara, California 95050")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%70s\n", "Contact: +1 (666) 222-555 | contact@littlelemon.com")
	// fmt.Printf("%60s\n\n", "Email: contact@littlelemon.com")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s", strings.Repeat("-", 42))
	fmt.Printf("%s", "INVOICE")
	fmt.Printf("%s\n", strings.Repeat("-", 42))
	time.Sleep(time.Millisecond * 200)
	rand.Seed(time.Now().Unix())
	fmt.Printf(" Receipt No: %d\n", rand.Intn(550)+1)
	fmt.Printf(" Date: %v\n", time.Now().Local().Format("06-Jan-02"))
	fmt.Printf(" Time: %v", time.Now().Local().Format("3:4 PM"))
	time.Sleep(time.Millisecond * 200)
	generateBill()
	tax := 18 * subTotalBill / (100)
	grandTotal := subTotalBill + tax
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: $%.2f\n", "Tax", tax)
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: $%.2f\n", "Total", grandTotal)
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
}

func modifyOrder() {
	for {
		var isOrderOK string
		fmt.Println("Do you want to modify your order? [y/n]")
		fmt.Scan(&isOrderOK)
		if isOrderOK != "y" {
			return
		}
		var serialNo uint
		var modifyType uint
		fmt.Println("Please enter the receipt no. to proceed: ")
		fmt.Println("Press '1' to update item quantity.")
		fmt.Println("Press '2' to delete an item from your order.")
		fmt.Println("Press '3' to add item(s) to your order.")
		fmt.Scan(&modifyType)
		switch modifyType {
		case 1:
			printMenu()
			fmt.Println("Please enter the Serial No. of the item to be updated: ")
			fmt.Scan(&serialNo)
			updateQuantity(serialNo)
		case 2:
			printMenu()
			fmt.Println("Please enter the Serial No. of the item to be deleted: ")
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
	for _, targetItem := range menu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName
			oldQuantity := customerOrder[itemName]
			fmt.Printf("Current quantity of %v is %v.\n", itemName, oldQuantity)
			fmt.Printf("Now, enter the updated quantity of %v to be ordered: \n", itemName)
			fmt.Scan(&newQuantity)
			if newQuantity == 0 {
				delFromOrder(serialNo)
				return
			}
			fmt.Printf("")
			customerOrder[targetItem.itemName] = newQuantity
			fmt.Printf("Updated the quantity of %v from %v to %v.\n", itemName, oldQuantity, newQuantity)
			subTotalBill -= float64(oldQuantity) * float64(targetItem.itemPrice)
			subTotalBill += float64(newQuantity) * float64(targetItem.itemPrice)
			break
		}
	}

}

func delFromOrder(serialNo uint) {
	for _, targetItem := range menu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName
			subTotalBill -= float64(customerOrder[itemName]) * float64(targetItem.itemPrice)
			delete(customerOrder, itemName)
			fmt.Printf("%v removed from your order.\n", itemName)
			break
		}
	}
}

func insertIntoOrder() {
	orderItems()
}

func main() {
	var customerName string
	fmt.Println("Please provide your name: ")
	fmt.Scan(&customerName)
	greet(customerName)
	orderItems()
	displayGeneratingBill()
	generateBill()
	modifyOrder()
	printFinalBill()
	sayTata(customerName)
}
