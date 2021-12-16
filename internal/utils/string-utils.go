package stringUtils

import (
    "errors"
    "math"
)

func ChunkSplit(s string, l int) ([]string, error) {
    if l <= 0 {
        return nil, errors.New("length should be > 0")
    }

    i := 0
    result := make([]string, int(math.Ceil(float64(len(s)) / float64(l))))
    for len(s) > 0 {
        var length int
        if len(s) > l {
            length = l
        } else {
            length = len(s)
        }
        temp := s[:length]
        result[i] = temp
        i += 1
        s = s[length:]
    }

    return result, nil
}
