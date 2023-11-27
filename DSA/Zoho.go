package DSA

func MergeTwoArray(arr1 []int ,arr2 [] int ) []int {
  var  result [] int 
 i,j:=0,0
  for i<len(arr1) && j<len(arr2) {
	  if arr1[i]<arr2[j]{
		result=append(result,arr1[i])
		i++
	  }else if arr1[i]==arr2[j]{
		result=append(result,arr1[i])
		result=append(result,arr2[j])
		i++
		j++
		}else{
		result=append(result,arr2[j])
		j++
	  }
  }

  return result
}