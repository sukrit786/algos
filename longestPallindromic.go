func longestPalindrome(s string) string {
    if s == "" || len(s) < 1 {
        return ""
    }
    
    start := 0
    end := 0
    for i := 0; i < len(s); i++ {
        len1 := Expand(s, i, i)
        len2 := Expand(s, i, i+1)
        length := Max(len1, len2)
        
        if length > end-start {
            start = i - (length-1) / 2
            end = i + (length / 2)
        }
    }
    
    return s[start:end+1]
    
}

func Max(i, j int) int {
    if i > j {
        return i
    }
    return j
}

func Expand(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    
    return right - left - 1
}