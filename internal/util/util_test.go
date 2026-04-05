package util

import (
	"testing"

	"github.com/erwindrsno/Quotation-Builder/internal/client"
)

func TestGenerateQuoteNumber(t *testing.T) {
	client := client.Client{
		Name: "ED",
	}
	companyName := "SPM"

	expected := "SPM/ED/050426/"

	result := GenerateQuoteNumber(client, companyName)

	if result != expected {
		t.Errorf("GenerateQuoteNumber() = %v, want %v", result, expected)
	}
}
