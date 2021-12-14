package main

import (
    "fmt"
    "github.com/go-co-op/gocron"
    "time"
)

func task(task string) {
    fmt.Println(task)
}

func main() {
    s := gocron.NewScheduler(time.Local)
    s.Cron("*/1 * * * *").Do(task, "task1")
    s.Every(10).Seconds().Do(task, "task2")
    s.StartAsync()
    
    fmt.Scanf("\n")
}
