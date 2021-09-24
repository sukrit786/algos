func sortArray(nums []int) []int {
    nums = heapify(nums,len(nums))
    fmt.Print(nums)
    nums = sortViaHeap(nums,len(nums))
    return nums
}
//bubble
func sortArrayBubble(nums []int) []int {
    for i:=0;i<len(nums);i++ {
        for j:=i+1;j<len(nums);j++ {
            if nums[i] > nums[j] {
               nums[i],nums[j] = nums[j],nums[i] 
            }
        }
    }
    return nums
}
// top to bottom
func heapify(arr[] int,n int)[]int {
    na:=[]int{}
    for i:=0;i<len(arr);i++{                                                                                                               
        na = append(na,arr[i]);
        if(i>0) {
            parent:= (i-1)/2
            child:=i
            for na[child]>na[parent]  {
                // if na[child]>na[parent] {
                    na[parent],na[child] = na[child],na[parent];
                    child = parent;
                    parent = (child-1)/2;
                // }
            }
        }
    }
    return na
}

func heapify2(arr[] int, n int)[]int {
    lnln:=n/2-1;//lastnonleafnode
    for i:=lnln;i>=0;i-- {
        parent := i;
        left := 2*i+1;
        right:=2*i+2;
        fmt.Print(left,right,n,"\n")
        for left<n && right<n && (arr[parent]<arr[left] || arr[parent]<arr[right]) {
            if arr[left]<arr[right] {
                arr[parent],arr[right] = arr[right],arr[parent];
                parent = right
            }else {
                arr[parent],arr[left] = arr[left],arr[parent];
                parent = left
            }
            left = 2*parent+1;
            right = 2*parent+2;
        }
        if left<n && right>=n {
                fmt.Print("in here le",i)
            if arr[left]>arr[parent] {
                arr[parent],arr[left] = arr[left],arr[parent];
            }
        }
        if right<n && left>=n {
                fmt.Print("in here ri",i)
            if arr[right]>arr[parent] {
                arr[parent],arr[right] = arr[right],arr[parent];
            }
        }
    }
    return arr
}


func sortViaHeap(arr []int,n int)[]int {
    templast:=n-1
    for templast>=0 {
        fmt.Print(templast)
        arr[0],arr[templast] = arr[templast],arr[0];
        templast--
        parent:= 0;
        left := 2*parent+1;
        right := 2*parent+2;
        for left<=templast && right<=templast && (arr[parent]<arr[left] || arr[parent]<arr[right]) {
            if arr[left]<arr[right] {
                arr[parent],arr[right] = arr[right],arr[parent];
                parent = right
            }else {
                arr[parent],arr[left] = arr[left],arr[parent];
                parent = left
            }
            left = 2*parent+1;
            right = 2*parent+2;
        }
        if left<=templast && right>templast {
            if arr[left]>arr[parent] {
                fmt.Print("he needed me lefty",left)
                arr[parent],arr[left] = arr[left],arr[parent];
            }
        }
        if right<=templast && left>templast {
            if arr[right]>arr[parent] {
                fmt.Print("he needed me righty")
                arr[parent],arr[right] = arr[right],arr[parent];
            }
        }
    }
    return arr
}

//quick sort
func qs(arr []int,low int,high int) {
    if low < high {
        pi:=partition(arr,low,high)
        qs(arr,low,pi-1);
        qs(arr,pi+1,high);
    }
}
func partition(arr []int,low int,high int)int {
    pi:=arr[high];
    
    i:=low
    
    for j:=low;j<high;j++ {
        if pi > arr[j] {
            arr[j],arr[i] = arr[i],arr[j];
            i++
        }
    }
    arr[i],arr[high] = arr[high],arr[i];
    return i
}