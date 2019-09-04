package main

import (
    "fmt"
    "time"

    "bitbucket.org/server-monitor/checkrlight/lib"
    "bitbucket.org/server-monitor/checkrlight/lib/conc"
    "bitbucket.org/server-monitor/checkrlight/app/features/checker"
)

func main() {
    urls := []string{
        // // Should time out
        // "google.com:81",

        "no-such-host-example",

        // // Should time out
        // "yahoo.com:81",

        "nmap.org",
        // "golang.org",
        "yahoo.com",
        "yahoo.com",
        // "duckduckgo.com",
        "nmap.org",
    }

    startTime := time.Now()
    tasks, fi, err := conc.Run(
      &checker.ConcFactoryType{URLs: urls},
      1000,
    )
    if err != nil {
        panic(err)
    }
    factory, ok := fi.(*checker.ConcFactoryType)
    if !ok {
        panic(lib.ErrTypeAssert(factory, &checker.ConcFactoryType{}))
    }
    if factory.Err != nil {
        panic(factory.Err)
    }
    endTime := time.Now()
    elapsed := endTime.Sub(startTime)
    factory.StatusMessages = append(factory.StatusMessages, fmt.Sprintf("Elapsed time: %s", elapsed))

    for _, task := range tasks {
        task.Print()
        fmt.Println()
    }

    if len(factory.WarningMessages) > 0 {
        fmt.Println("Warning messages...")
        for _, msg := range factory.WarningMessages {
            fmt.Println(msg)
        }
        fmt.Println()
    }

    if len(factory.StatusMessages) > 0 {
        fmt.Println("Status messages...")
        for _, msg := range factory.StatusMessages {
            fmt.Println(msg)
        }
        fmt.Println()
    }
}
