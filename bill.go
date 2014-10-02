package main

import (
  "fmt"
)

const taxRate float32 = 0.075
const tipRate float32 = 0.20

type Item struct {
  name string
  value float32
  purchasers []*Diner
}

type Discount struct {
  cost float32
  value float32
  minimum float32
}

type Diner struct {
  name string
}

type Bill struct {
  discount Discount
  items []*Item
  diners []*Diner
}

func (b *Bill) itemCount() int {
  return len(b.items)
}

func (b *Bill) addItem(i *Item) {
  b.items = append(b.items, i)
}

func (b *Bill) subtotal() float32 {
  var sum float32 = 0

  for _, item := range b.items {
    sum = sum + item.value
  }

  return sum
}

func (b *Bill) total() float32 {
  return b.subtotal() + b.tax() + b.tip() - b.appliedDiscount()
}

func (b *Bill) appliedDiscount() float32 {
  if (b.subtotal() < b.discount.minimum) {
    return 0
  } else {
    return b.discount.value
  }
}

func (b *Bill) tax() float32 {
  return b.subtotal() * taxRate
}

func (b *Bill) tip() float32 {
  return b.subtotal() * tipRate
}

func (b *Bill) discountApplied() bool {
  return b.appliedDiscount() > 0
}

func (b *Bill) dinerSubtotal(d *Diner) float32 {
  var sum float32 = 0
  for _, item := range b.items {
    if contains(item.purchasers, d) {
      purchaserCount := len(item.purchasers)
      sum = sum + (item.value / float32(purchaserCount))
    }
  }
  return sum
}

func contains(list []*Diner, d *Diner) bool {
  for _, i := range list {
    if i == d {
      return true
    }
  }
  return false
}

func main() {
  eugene := Diner{name: "Eugene"}
  tom := Diner{name: "Thomas"}
  item_1 := Item{name: "Chicken", value: 85.99, purchasers: []*Diner{&eugene}}
  item_2 := Item{name: "Falafel", value: 7.95, purchasers: []*Diner{&eugene, &tom}}
  item_3 := Item{name: "Salad", value: 3.95, purchasers: []*Diner{&eugene}}

  discount := Discount{cost: 6.00, value: 50.00, minimum: 100.00}
  bill := Bill{discount: discount}

  bill.addItem(&item_1)
  bill.addItem(&item_2)
  bill.addItem(&item_3)

  fmt.Printf("%v owes %v\n", eugene.name, bill.dinerSubtotal(&eugene))

  fmt.Printf("Bill discount costs %v, and the value of the discount is %v\n", bill.discount.cost, bill.discount.value)
  fmt.Printf("The bill has %v items\n", bill.itemCount())
  fmt.Printf("Subtotal is %v\n", bill.subtotal())
  fmt.Printf("Tax is %v\n", bill.tax())
  fmt.Printf("Tip is %v\n", bill.tip())
  fmt.Printf("Total is %v\n", bill.total())
  fmt.Printf("Discount was applied: %v\n", bill.discountApplied())
}
