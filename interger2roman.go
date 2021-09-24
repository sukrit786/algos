var (
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"} // 1,2,3,4,5,6,7,8,9
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"} // 10,20,30,60,50,40,70,80,90
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"} // 100,200,300,400,500, ...
	r3 = []string{"", "M", "MM", "MMM"}                                       // 1000,2000,3000
)

func intToRoman(num int) string {
	// 2048
	// This is efficient in Go. The 4 operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return r3[num%1e4/1e3] + r2[num%1e3/1e2] + r1[num%100/10] + r0[num%10]
}