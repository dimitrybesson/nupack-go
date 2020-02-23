package main

import (
  "math"
  "fmt"
)

type MarkupCalc struct {
  prodTypeMarkups map[string]float64
}

func NewMarkupCalc() *MarkupCalc {
  ptm := map[string]float64 {
    "food": 13.0,
    "drugs": 7.5,
    "electronics": 2.0,
  }
  return &MarkupCalc{
    prodTypeMarkups: ptm,
  }
}

func (mc *MarkupCalc) GetPrice(price float64, employees int, productType string) float64 {
  total := price + mc.getMarkup(price, employees, productType)
  return (total * 100) / 100
}

func (mc *MarkupCalc) getBaseMarkup(price float64) float64 {
  return math.Round(price * 5) / 100
}

func (mc *MarkupCalc) getEmployeeMarkup(price float64, employees int) float64 {
  empMarkup := float64(employees) * 1.2
  return math.Round(price * empMarkup) / 100
}

func (mc *MarkupCalc) getProductMarkup(price float64, productType string) float64 {
  prodMarkup, ok := mc.prodTypeMarkups[productType]
  if !ok {
    return 0
  }
  return math.Round(price * prodMarkup) / 100
}

func (mc *MarkupCalc) getMarkup(price float64, employees int, productType string) float64 {
  baseMarkup := mc.getBaseMarkup(price)
  basePrice := price + baseMarkup
  empMarkup := mc.getEmployeeMarkup(basePrice, employees)
  prodMarkup := mc.getProductMarkup(basePrice, productType)
  markupSum := baseMarkup + empMarkup + prodMarkup

  return (markupSum * 100) / 100
}

func main() {
  c := NewMarkupCalc()

  fmt.Printf("Input: $%0.2f, %d person, %s\n", 1299.99, 3, "food")
  fmt.Printf("Output: $%0.2f\n", c.GetPrice(1299.99, 3, "food"))

  fmt.Printf("Input: $%0.2f, %d person, %s\n", 5432.00, 1, "drugs")
  fmt.Printf("Output: $%0.2f\n", c.GetPrice(5432.00, 1, "drugs"))

  fmt.Printf("Input: $%0.2f, %d person, %s\n", 12456.95, 4, "books")
  fmt.Printf("Output: $%0.2f\n", c.GetPrice(12456.95, 4, "books"))

}
