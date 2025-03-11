package boba

import (
	"fmt"
)

func QueryUsers() string {
	return "Finding all users!"
}

func QueryItems() string {
	return "Finding all items!"
}

func QueryOrders() string {
	return "Finding all orders!"
}
func QueryInvoices() string {
	return "Finding all invoices!"
}
func QueryPayments() string {
	return "Finding all payments!"
}

func (m model) View() string {
	// The header
	s := "\n\nFETCH:\n\n"

	// Iterate over our choices
	for i, choice := range m.options {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		getting := ""
		if _, ok := m.selected[i]; ok {
			switch m.options[i] {
			case "Users":
				result := QueryUsers()
				getting = ("-> " + m.spinner.View() + result)
			case "Items":
				result := QueryItems()
				getting = ("-> " + m.spinner.View() + result)
			case "Orders":
				result := QueryOrders()
				getting = ("-> " + m.spinner.View() + result)
			case "Invoices":
				result := QueryInvoices()
				getting = ("-> " + m.spinner.View() + result)
			case "Payments":
				result := QueryPayments()
				getting = ("-> " + m.spinner.View() + result)
			}
		} else {
			getting = ""
		}

		s += fmt.Sprintf("%s [%s] %s %s\n", cursor, checked, choice, getting)
	}

	// The footer
	s += "\nPress q to quit.\n"

	return s
}
