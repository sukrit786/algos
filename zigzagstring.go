
//4ms 6.4MB 
func convert(s string, numRows int) string {
    // arr:=[numRows][len(s)]int{{}};
    n:=len(s)
    if numRows==1 || numRows>=len(s) {
        return s
    }
    arr := make([][]byte, numRows)
    for i := range arr {
        arr[i] = []byte{}
    }
    direction:=-1;
    row:=0
    for i:=0;i<n;i++ {
        arr[row] = append(arr[row],s[i])
        if(row == 0 || row == numRows-1) {
            direction*=-1
        }
        row+=direction
    }
    // str:=""
    
    return string(bytes.Join(arr, []byte{}))
}
//less memory(4.2MB) same time
func convert(s string, numRows int) string {
    n:=len(s)
    cycle:=2*numRows-2
    res:=[]byte{}
    if(numRows==1) {
        return s
    }
    for i:=0;i<numRows;i++ {
        for j:=i;j<n;j+=cycle {
            res = append(res,s[j])
            k:=j+cycle-2*i
            if(i!=0 && i!=numRows-1 && k<len(s)) {
                res = append(res,s[k])
            }
        }
    }  
    return string(res)
}
