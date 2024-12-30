func rotateString(s string, goal string) bool {
    if len(s) == len(goal) && strings.Contains(s+s, goal){
        return true
    }
    return false
}