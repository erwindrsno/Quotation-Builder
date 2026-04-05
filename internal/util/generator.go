package util

import (
	"fmt"
	"time"

	"github.com/erwindrsno/Quotation-Builder/internal/client"
)

// TODO: To be fix. Should I use company id or company initials? and client's initials might need to store also.
func GenerateQuoteNumber(c client.Client, companyName string) string {
	now := time.Now()
	// 02 = Day, 01 = Month, 06 = Year (2 digits)
	dateStr := now.Format("020106")

	return fmt.Sprintf("%s/%s/%s/", companyName, c.Name, dateStr)
}
