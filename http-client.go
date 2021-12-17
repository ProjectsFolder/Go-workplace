package main

import (
    "fmt"
    "go.uber.org/dig"
    "log"
    "os"
    "strconv"
    "workplace/internal/injector"
    "workplace/internal/services"
)

func main() {
    err := injector.GetContainer().Invoke(func(
        inject struct {
           dig.In
           Billing services.BillingProviderInterface
           Logger *log.Logger `name:"telegramLogger"`
        },
    ) {
        contractId, _ := strconv.Atoi(os.Args[1])
        contract, err := inject.Billing.GetContract(contractId)
        if err != nil {
            inject.Logger.Println("billing get-contract exception:", err.Error())
        }
        if contract != nil {
            log.Println(fmt.Sprintf(
                "id: %d; title: %s; balance: %f; house_id: %d",
                contract.Id,
                contract.Title,
                contract.Balance,
                contract.HouseId,
            ),
            )
        }
    })
    if err != nil {
        log.Println(err)
    }

    fmt.Scanf("\n")
}
