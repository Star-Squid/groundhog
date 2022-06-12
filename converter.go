package main

import (
	"fmt"
)

func main() {
  currencies := map[string]float32{
    "JPY": 0.2,
    "EUR": 0.5,
  }

  var dollaramount float32
  var currency string

  fmt.Println("What's the dollar amount?")
  fmt.Scan(&dollaramount)
    if dollaramount == 0 {
      fmt.Println("Error, not a valid number")
    } else {

      fmt.Println("What's the currency?")
      fmt.Scan(&currency)

        if dollaramount == 0 {
       fmt.Println("Error, not a valid currency")
        } else {

          rate,isValid := currencies[currency]
          if !isValid {
            fmt.Println("Currency not on the list")
          } else {
            fmt.Println(dollaramount*rate)
          }

        }
    }
}
