package checker

import (
    "fmt"
    "strings"
    "regexp"
)

// SanitizeURLs ...
func SanitizeURLs(inputUrls []string) ([][]string, []string, error) {
    const (
        limitNumberOfURLs = 10
    )

    warningMessages := []string{}

    urls, msgs := removeDuplicates(inputUrls)
    warningMessages = appendWarningMessages(warningMessages, msgs...)

    urls, msg := chopExtra(urls, limitNumberOfURLs)
    warningMessages = appendWarningMessages(warningMessages, msg)

    urlPacks := insertScheme(urls)

    return urlPacks, warningMessages, nil
}

func chopExtra(inputs []string, limit int) ([]string, string) {
    if len(inputs) > limit {
        msg := fmt.Sprintf("Number of unique URLs, %d, exceeded, only %d allowed", len(inputs), limit)
        return inputs[:limit], msg
    }

    return inputs, ""
}

func removeDuplicates(urls []string) ([]string, []string) {
    urlTable := map[string]int{}
    uniq := []string{}

    msgs := []string{}
    for _, url := range urls {
        if count, ok := urlTable[url]; ok {
            if count < 2 {
                msg := fmt.Sprintf("Duplicate URL found: %s", url)
                msgs = append(msgs, msg)
            }
            urlTable[url]++
        } else {
            urlTable[url]++
            uniq = append(uniq, url)
        }
    }

    return uniq, msgs
}

func appendWarningMessages(warningMessages []string, msgs... string) []string {
    if len(msgs) == 0 {
        return warningMessages
    }

    for _, msg := range msgs {
        if strings.TrimSpace(msg) == "" {
            continue
        }

        warningMessages = append(warningMessages, msg)
    }

    return warningMessages
}

func insertScheme(inputUrls []string) ([][]string) {
    var scheme = regexp.MustCompile(`\Ahttps?:\/\/`)
    urls := [][]string{}

    for _, url := range inputUrls {
        if scheme.MatchString(url) {
            urls = append(urls, []string{scheme.ReplaceAllString(url, ""), url})

            // // https://stackoverflow.com/questions/16703501/replace-one-occurrence-with-regexp
            // // Another approach to this although ReplaceAllString should be robust since
            // //   the regexp is anchored.
            // // I don't know why regexp library does not have a "ReplaceOne" function.
            // found := scheme.FindString(url)
            // name := url
            // if found != "" {
            //     name = (strings.Replace(name, found, "", 1))
            // }
        } else {
            urls = append(
                urls,
                []string{url, strings.Join([]string{"https://", url}, "")},
            )
        }
    }

    return urls
}
