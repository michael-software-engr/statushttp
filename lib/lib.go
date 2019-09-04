package lib

import (
    "fmt"
)

// ErrTypeAssert ...
func ErrTypeAssert(
    actual interface{},
    desiredTypeVar interface{},
) (error) {
    return fmt.Errorf(
        "... type of interface value, %T, should be of type %T",
        actual,
        desiredTypeVar,
    )
}
