package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"lab4-variant11/pkg/printshop"
)

func main() {
	fmt.Println("PageCost calculates the cost of printing the specified number of pages.")
	fmt.Println("OrderCost calculates the total cost for several copies.")
	fmt.Println("ApplyBulkDiscount changes the cost through a pointer and applies a percentage discount.")
	fmt.Println("FormatPrintReport creates a text report for the order.")
	fmt.Println()

	orderID := uuid.New().String()
	pages := 120
	copies := 5
	pricePerPage := 2.50

	pageCost, err := printshop.PageCost(pages, pricePerPage)
	if err != nil {
		fmt.Printf("PageCost error: %v\n", err)
		return
	}
	fmt.Printf("Cost of one copy: %.2f\n", pageCost)

	orderCost, err := printshop.OrderCost(pages, copies, pricePerPage)
	if err != nil {
		fmt.Printf("OrderCost error: %v\n", err)
		return
	}
	fmt.Printf("Cost of the order before discount: %.2f\n", orderCost)

	discountPercent := 10.0
	err = printshop.ApplyBulkDiscount(&orderCost, discountPercent)
	if err != nil {
		fmt.Printf("ApplyBulkDiscount error: %v\n", err)
		return
	}
	fmt.Printf("Cost after %.2f%% discount: %.2f\n", discountPercent, orderCost)

	report, err := printshop.FormatPrintReport(orderID, pages, copies, orderCost)
	if err != nil {
		fmt.Printf("FormatPrintReport error: %v\n", err)
		return
	}

	color.Cyan("Print report:")
	fmt.Printf("%s\n", report)
}
