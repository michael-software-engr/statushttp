package conc

// TaskInterface ...
type TaskInterface interface {
    Process()
    Print()
}

// FactoryInterface ...
type FactoryInterface interface {
    BuildTasks(chan<- TaskInterface)
}
