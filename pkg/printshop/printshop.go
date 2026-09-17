// Package printshop contains calculations and report formatting for a print shop.
package printshop

import "fmt"

// PageCost calculates the cost of printing the specified number of pages.
func PageCost(pages int, pricePerPage float64) (float64, error) {
	if pages <= 0 {
		return 0, fmt.Errorf("pages must be greater than zero")
	}
	if pricePerPage <= 0 {
		return 0, fmt.Errorf("price per page must be greater than zero")
	}

	return float64(pages) * pricePerPage, nil
}

// OrderCost calculates the total cost of printing several copies of a document.
func OrderCost(pages, copies int, pricePerPage float64) (float64, error) {
	if copies <= 0 {
		return 0, fmt.Errorf("copies must be greater than zero")
	}

	oneCopyCost, err := PageCost(pages, pricePerPage)
	if err != nil {
		return 0, fmt.Errorf("calculate one copy cost: %w", err)
	}

	return oneCopyCost * float64(copies), nil
}

// ApplyBulkDiscount reduces cost by the specified percentage through a pointer.
func ApplyBulkDiscount(cost *float64, percent float64) error {
	if cost == nil {
		return fmt.Errorf("cost pointer must not be nil")
	}
	if *cost < 0 {
		return fmt.Errorf("cost must not be negative")
	}
	if percent < 0 || percent > 100 {
		return fmt.Errorf("discount percent must be between 0 and 100")
	}

	*cost *= 1 - percent/100
	return nil
}

// FormatPrintReport formats a printable report for an order.
func FormatPrintReport(orderID string, pages, copies int, cost float64) (string, error) {
	if orderID == "" {
		return "", fmt.Errorf("order ID must not be empty")
	}
	if pages <= 0 {
		return "", fmt.Errorf("pages must be greater than zero")
	}
	if copies <= 0 {
		return "", fmt.Errorf("copies must be greater than zero")
	}
	if cost < 0 {
		return "", fmt.Errorf("cost must not be negative")
	}

	return fmt.Sprintf("Order %s: %d pages x %d copies, total cost: %.2f", orderID, pages, copies, cost), nil
}
