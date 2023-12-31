package DSA
import (
	"sort" 
     "errors"
)
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
func BUYStocks(){
   


}

// second largest number 
func SecondLarget(arr []int) int {
	largest,second:=-1,0
   for _,i:=range arr {
	   if largest<i{
		 largest,second=i,largest
	   }
   }
   return second 
}





func climbStairs(n int) int {
    if n==0{
        return 1
    }else if n<0{
       return 0
    }
    return climbStairs(n-1)+climbStairs(n-2)
}



func Divide(a,b float32)(float32,error){
	if b==0{
	   return 0,errors.New("cannot be divided")
	}
	return a/b, nil
}


func longestCommonSubsequence(text1 string, text2 string) int {
	dp:=make([][]int,len(text1)+1)
	for i := range dp {
	 dp[i]= make([]int, len(text2)+1)
 }
	for i:=0;i<=len(text1);i++{
		for j:=0;j<=len(text2);j++{
			  if i==0 || j==0{
				  dp[i][j]=0
			  }else if text1[i-1]==text2[j-1]{
				  dp[i][j]=dp[i-1][j-1]+1
			  }else{
				  dp[i][j]=max(dp[i][j-1],dp[i-1][j])
			  }
		}
	}
	return dp[len(text1)][len(text2)]
 }