package conc

import (
    "sync"
    "time"
)

// Run ...
func Run(fi FactoryInterface, workersCount int) ([]TaskInterface, FactoryInterface, error) {
    var wg sync.WaitGroup

    // START: input processing
    inChan := make(chan TaskInterface)
    wg.Add(1)
    go buildTasks(&wg, fi, inChan)
    // END: input processing

    // START: actual processing
    outChan := make(chan TaskInterface)
    for ix := 0; ix < workersCount; ix++ {
        wg.Add(1)
        go process(&wg, inChan, outChan)
    }
    // END: actual processing

    // START: wait
    go wait(&wg, outChan)
    // END: wait

    tasks := []TaskInterface{}
    for task := range outChan {
        tasks = append(tasks, task)
    }

    return tasks, fi, nil
}

func buildTasks(
    wg *sync.WaitGroup,
    fi FactoryInterface,
    inChan chan<- TaskInterface,
) {
    fi.BuildTasks(inChan)
    close(inChan)
    wg.Done()
}

func process(
    wg *sync.WaitGroup,
    inChan <-chan TaskInterface,
    outChan chan<- TaskInterface,
) {
    defer wg.Done()

    delay := 1
    currentTask := 0
    for task := range inChan {
        if currentTask > 0 {
            time.Sleep(time.Duration(delay) * time.Second)
        }

        task.Process()
        outChan <- task

        currentTask++
    }
}

func wait(wg *sync.WaitGroup, outChan chan TaskInterface) {
    wg.Wait()
    close(outChan)
}
