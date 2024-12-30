func canConstruct(ransomNote string, magazine string) bool {
    vals := make(map[rune]int)

    for  m := range magazine{
        vals[m] = vals[m] + 1
    }

    for r := range ransomNote {
        if vals[r] == 0{
            return false
        }
        vals[r] = vals[r] - 1
    }
    return true
}