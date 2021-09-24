
// func lengthOfLongestSubstring(s string) int {
//     m := make(map[byte]int);
//     ml,start:=0,1
//     for i:=0;i<len(s);i++ {
//         if m[s[i]] != 0 {
//             start = maxInt(start,m[s[i]]+1)
//         }
//         m[s[i]] = i+1;
//         ml = maxInt(ml,i-start+2)
        
//     }
//     fmt.Print(ml,m)
//     return ml
// }
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func lengthOfLongestSubstring(s string) int {
	maxLen, start := 0, 0
	table := [128]int{}
	for i, _ := range table {
		table[i] = -1
	}
	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		maxLen = maxInt(maxLen, i - start + 1)
	}
	return maxLen
}