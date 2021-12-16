package stringUtils

func StringChunk(s string, l int) []string {
    var result []string
    for len(s) > 0 {
        var length int
        if len(s) > l {
            length = l
        } else {
            length = len(s)
        }
        temp := s[:length]
        result = append(result, temp)
        s = s[length:]
    }

    return result
}
