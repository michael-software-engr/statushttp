package checker

import (
    "net/http"
    "encoding/json"
    "log"
    "strings"
    "io/ioutil"
    "time"
    "fmt"

    "bitbucket.org/server-monitor/checkrlight/lib/conc"

    checkerfeature "bitbucket.org/server-monitor/checkrlight/app/features/checker"
)

const (
    // Name ...
    Name = checkerfeature.Name
)

// Handler ...
func Handler(writer http.ResponseWriter, request *http.Request) {
    var handler func(request *http.Request) ([]byte, error)
    log.Print(request)
    switch (request.Method) {
    case "POST":
        handler = post
    default:
        log.Fatalf("... request method %#v not supported", request.Method)
    }

    jEnc, err := handler(request)

    if err != nil {
        log.Fatal(err)
    }
    writer.Write(jEnc)
}

func post(request *http.Request) ([]byte, error) {
    emptyWarningMessages := []string{}
    emptyStatusMessages := []string{}
    emptyResults := []checkerfeature.TargetType{}

    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        return jEncErr(err, emptyWarningMessages, emptyStatusMessages, emptyResults)
    }

    var inputStr string
    err = json.Unmarshal(body, &inputStr)
    if err != nil {
        return jEncErr(err, emptyWarningMessages, emptyStatusMessages, emptyResults)
    }
    inputs := strings.Fields(inputStr)

    startTime := time.Now()

    // Concurrent: START
    tasks, fi, err := conc.Run(
      &checkerfeature.ConcFactoryType{URLs: inputs},
      1000,
    )
    if err != nil {
        return jEncErr(err, emptyWarningMessages, emptyStatusMessages, emptyResults)
    }

    factory, ok := fi.(*checkerfeature.ConcFactoryType)
    if !ok {
        return jEncErrTypeAssert(
            factory,
            &checkerfeature.ConcFactoryType{},
            emptyWarningMessages,
            emptyStatusMessages,
            emptyResults,
        )
    }

    warningMessages := factory.WarningMessages
    statusMessages := factory.StatusMessages
    results := []checkerfeature.TargetType{}

    for _, task := range tasks {
        result, ok := task.(*checkerfeature.TargetType)
        if !ok {
            return jEncErrTypeAssert(
                result,
                &checkerfeature.TargetType{},
                warningMessages,
                statusMessages,
                results,
            )
        }
        results = append(results, *result)
    }
    // Concurrent: END

    endTime := time.Now()

    elapsed := endTime.Sub(startTime)
    statusMessages = append(statusMessages, fmt.Sprintf("Elapsed time: %s", elapsed))
    factory.StatusMessages = statusMessages

    // // TODO: probably delete later
    // // Not concurrent: START
    // results, statusMessages, err := Run(inputs)
    // if err != nil {
    //     return jEncErr(err, emptyWarningMessages, emptyStatusMessages, emptyResults)
    // }
    // // Not concurrent: END

    empty := []byte{}
    jEnc, err := jsonify(results, nil, warningMessages, statusMessages)
    if err != nil {
        return empty, err
    }

    if len(factory.WarningMessages) > 0 {
        log.Print("Warning messages: START...")
        for _, msg := range factory.WarningMessages {
            log.Print(msg)
        }
        log.Print("Warning messages: END")
    }

    if len(factory.StatusMessages) > 0 {
        log.Print("Status messages: START...")
        for _, msg := range factory.StatusMessages {
            log.Print(msg)
        }
        log.Print("Status messages: END")
    }

    return jEnc, nil
}
