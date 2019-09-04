package checker

import (
    "log"
)

// Run ...
//   TODO: not concurrent, probably delete later
func Run(inputs []string) ([]TargetType, []string, error) {
    emptyTargets := []TargetType{}
    emptyStatusMsgs := []string{}

    sanitizedUrls, statusMsgs, err := SanitizeURLs(inputs)
    if err != nil {
        return emptyTargets, emptyStatusMsgs, err
    }

    results := []TargetType{}

    for _, urlPack := range sanitizedUrls {
        name := urlPack[0]
        url := urlPack[1]
        target, err := Exec(TargetType{
            Name: name,
            URL: url,
        })
        if err != nil {
            return results, statusMsgs, err
        }

        results = append(results, target)

        // DEBUG
        log.Printf("%#v\n", target)
    }

    return results, statusMsgs, nil
}
