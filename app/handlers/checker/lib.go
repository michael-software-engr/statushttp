package checker

import (
    "encoding/json"
    "fmt"

    checkerfeature "bitbucket.org/server-monitor/checkrlight/app/features/checker"
)

func jsonify(
    results []checkerfeature.TargetType,
    err error,
    warningMessages []string,
    statusMessages []string,
) ([]byte, error) {
    type jType struct {
        Results []checkerfeature.TargetType

        // Apparently, error type is not automatically marshaled.
        // Error error
        Error string

        WarningMessages []string
        StatusMessages []string
    }
    errStr := ""
    if err != nil {
        errStr = err.Error()
    }
    jEnc, jErr := json.Marshal(jType{
            Results: results,
            Error: errStr,
            WarningMessages: warningMessages,
            StatusMessages: statusMessages,
        },
    )

    empty := []byte{}
    if jErr != nil {
        jEnc, j2Err := json.Marshal(jType{
            Results: results,
            Error: jErr.Error(),
            WarningMessages: warningMessages,
            StatusMessages: statusMessages,
        })
        if j2Err != nil {
            return empty, j2Err
        }
        return jEnc, nil
    }

    return jEnc, nil
}

func jEncErr(
    err error,
    warningMessages []string,
    statusMessages []string,
    results []checkerfeature.TargetType,
) ([]byte, error) {
    empty := []byte{}
    jEnc, jErr := jsonify(results, err, warningMessages, statusMessages)
    if jErr != nil {
        return empty, jErr
    }
    return jEnc, nil
}

func jEncErrTypeAssert(
    factory interface{},
    desiredTypeVar interface{},
    warningMessages []string,
    statusMessages []string,
    results []checkerfeature.TargetType,
) ([]byte, error) {
    return jEncErr(
        fmt.Errorf(
            "... type of interface value, %T, should be of type %T",
            factory,
            desiredTypeVar,
        ),
        warningMessages,
        statusMessages,
        results,
    )
}
