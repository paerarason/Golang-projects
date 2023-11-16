package DSA
import "sort"
//Check if pair with given Sum exists in Array (Two Sum)
func TwoSum(target int, arr[] int) (int , int) {
     sort.Ints(arr)
	 
	 //for loop 
    left,right:=0,len(arr)-1
	 for left<right  {
         if arr[left]+arr[right]<target{
			right--
		 }else if arr[left]+arr[right]>target {
            left++
		 }else{
			return right,left
		 }
	 }
	 return -1,-1
}

//Best Time to Buy and Sell Stock
//

func BUYStocks(){
   


}