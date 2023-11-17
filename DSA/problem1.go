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

//Product of Array except itself

func ProductOfItself( arr[] int) []int {
	right:=arr[0:]
	left:=arr[0:]
	prod:=arr[0:]

	left[0]=1
	right[len(arr)-1]=1
	
	//construct right array 
	for i:=len(arr)-2;i>0;i-- {
		right[i]=arr[i]*right[i+1]   
	}
		//construct left array
	for i:=1;i<len(arr);i++ {
		left[i]=arr[i]*left[i-1]   
	}
	
	// product 
	for i:=0;i<len(arr);i++ {
		prod[i]=left[i]*right[i]   
	}

  return prod
}


//Best Time to Buy and Sell Stock
//

func BUYStocks(){
<<<<<<< HEAD
   


}

//

func climbStairs(n int) int {
    if n==0{
        return 1
    }else if n<0{
       return 0
    }
    return climbStairs(n-1)+climbStairs(n-2)
}