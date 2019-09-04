package checker

import (
    "fmt"
    "log"

    "bitbucket.org/server-monitor/checkrlight/lib/conc"
)

// Process ...
func (target *TargetType) Process() {
    result, err := Exec(TargetType{
        Name: target.URL,
        URL: target.URL,
    })

    if err != nil {
        target.Err = err
        target.StatusMessage = err.Error()
        return
    }

    target.Passed = result.Passed
    target.StatusMessage = result.StatusMessage

    // DEBUG
    log.Printf("%#v\n", target)
}

// Print ...
//   TODO: seems extraneous, probably delete later
func (target *TargetType) Print() {
  if target.Err != nil {
    fmt.Printf("ERROR: %#v\n", target)
  }

  fmt.Printf("%#v\n", target)
}

// ConcFactoryType ...
type ConcFactoryType struct {
  URLs []string
  Err error
  WarningMessages []string
  StatusMessages []string
}

// BuildTasks ...
func (factory *ConcFactoryType) BuildTasks(inputChannel chan<- conc.TaskInterface) {
    sanitizedUrls, warningMessages, err := SanitizeURLs(factory.URLs)
    if err != nil {
        factory.Err = err
        return
    }

    factory.WarningMessages = warningMessages

    for ix, urlPack := range sanitizedUrls {
        name := urlPack[0]
        url := urlPack[1]
        inputChannel <- &TargetType{Name: name, URL: url, Ix: ix}
    }
}
