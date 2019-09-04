package checker

import (
    "net"
    "net/http"
    "net/url"
    "time"
)

// Name ...
const Name = "checker"

const (
    defaultTimeOut = 10
)

// TargetType ...
type TargetType struct {
    Ix int // Used to sort as they come in because they'll be out of order as they are processed concurrently
    Name string
    URL string
    Timeout time.Duration
    Passed bool
    Err error // Redundant for now because all messages are put inside StatusMessage
    StatusMessage string
}

// Exec ...
func Exec(target TargetType) (TargetType, error) {
    target, err := setup(target)
    empty := TargetType{}
    if err != nil {
        return empty, err
    }

    resp, err := request(target)
    if err != nil {
        statusMsg, err := errorTriage(err)
        if err != nil {
            return empty, err
        }
        target.StatusMessage = statusMsg
        return target, nil
    }

    if resp.StatusCode == http.StatusOK {
        target.Passed = true
    }

    return target, nil
}

func setup(target TargetType) (TargetType, error) {
    target.Timeout = time.Duration(defaultTimeOut * time.Second)

    return target, nil
}

func request(target TargetType) (*http.Response, error) {
    client := http.Client{
        Timeout: target.Timeout,
    }

    resp, err := client.Get(target.URL)
    empty := &http.Response{}
    if err != nil {
        return empty, err
    }
    defer resp.Body.Close()

    return resp, nil
}

func errorTriage(err error) (string, error) {
    empty := ""
    switch errUnderlying := err.(type) {
    case *url.Error:
        return urlError(errUnderlying, err)
    }

    return empty, err
}

func urlError(urlErr *url.Error, err error) (string, error) {
    empty := ""
    switch urlErrUnderlying := urlErr.Err.(type) {
    case *net.OpError:
        switch urlErrU2 := urlErrUnderlying.Err.(type) {
        case *net.DNSError:
            errStr := urlErrU2.Err
            switch true {
            case errStr == "no such host":
                return errStr, nil
            }
        }
    case net.Error:
        if urlErrUnderlying.Timeout() {
            return urlErrUnderlying.Error(), nil
        }
    }

    return empty, err
}
