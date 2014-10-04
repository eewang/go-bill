package main

import (
  "net/http"
  "io"
  "fmt"
  "strconv"
)

const (
  indexHTML = `
  <html>
    <body>
      <h1>Hello World</h1>
      <form name="bill" action="inputBill" method="post">
        Restaurant: <input type="text" name="restaurant"></input><br>
        Item 1 Name: <input type="text" name="item_name"></input>
        Item 1 Cost: <input type="number" step="0.01" name="item_value"></input><br>
        <input type="submit" value="Submit">
      </form>
    </body>
  </html>
  `
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html;charset=utf-8")
  io.WriteString(w, indexHTML)
}

func billHandler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Printf("%v\n", r.Method)
  fmt.Printf("%v\n", r.URL)

  itemValue, _ := strconv.ParseFloat(r.Form["item_value"][0], 32)

  item := Item{name: r.Form["item_name"][0], value: float32(itemValue)}
  bill := Bill{restaurant: r.Form["restaurant"][0]}

  bill.addItem(&item)

  fmt.Printf("The bill subtotal is: %v\n", bill.subtotal())
  fmt.Printf("The bill tax is: %v\n", bill.tax())
  fmt.Printf("The bill tip is: %v\n", bill.tip())
  fmt.Printf("The bill total is: %v\n", bill.total())
}