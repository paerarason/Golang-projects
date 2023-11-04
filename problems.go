package main

func TwoSum(arr []int,target int ) (int, int) {
    heapsort(&arr,len(arr))
	left,right:=0,len(arr)-1
	for i:=0;i<len(arr);i++{
        if arr[left]+arr[right]==target{
		   return arr[left],arr[right]
		} else if arr[left]+arr[right]<target && left<right{
			left++
		} else if arr[left]+arr[right]>target && left<right {
			right--
		} else if left>right{
			break
		}
	} 
	return -1,-1
}