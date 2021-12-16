package main

import (
    "fmt"
    "os"
    "strconv"
    "workplace/internal/injector"
    "workplace/internal/services"
)

func main() {
    injector.GetContainer().Invoke(func(billing services.BillingProviderInterface, telegram *services.Telegram) {
       contractId, _ := strconv.Atoi(os.Args[1])
       contract, err := billing.GetContract(contractId)
       if err != nil {
           telegram.Log("billing get-contract exception:", err.Error())
       }
       if contract != nil {
           fmt.Println(fmt.Sprintf(
               "id: %d; title: %s; balance: %f; house_id: %d",
               contract.Id,
               contract.Title,
               contract.Balance,
               contract.HouseId,
           ),
           )
       }
    })

    fmt.Scanf("\n")
}
