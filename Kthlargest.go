//Heap ones work much better

func heapify(arr[] int, n int)[]int {
    lnln:=n/2-1;//lastnonleafnode
    for i:=lnln;i>=0;i-- {
        parent := i;
        left := 2*i+1;
        right:=2*i+2;
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
            if arr[left]>arr[parent] {
                arr[parent],arr[left] = arr[left],arr[parent];
            }
        }
        if right<n && left>=n {
            if arr[right]>arr[parent] {
                arr[parent],arr[right] = arr[right],arr[parent];
            }
        }
    }
    return arr
}

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


// [6 5 1 3 2 4]
func sortViaHeap(arr []int,n int,k int)int {
    templast:=n-1;
    m:=0
    for templast>=0 {
        arr[0],arr[templast] = arr[templast],arr[0];
        if m==k-1 {
            return arr[templast]; 
        }
        m++
        templast--;

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
                arr[parent],arr[left] = arr[left],arr[parent];
            }
        }
        if right<=templast && left>templast {
            if arr[right]>arr[parent] {
                arr[parent],arr[right] = arr[right],arr[parent];
            }
        }
    }
    return 0
}

func findKthLargest(nums []int, k int) int {
    qs(nums,0,len(nums)-1)
    nums = heapify(nums,len(nums));
    fmt.Print(nums)
    x := sortViaHeap(nums,len(nums),k);
    x:=nums[len(nums)-k]
    return x
}
