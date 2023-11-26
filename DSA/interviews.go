package DSA

func PalindromNumber(num int ) bool {
 	result,res:=num,0
	for num>0 {
		res*=10+num%10
		num/=10
	}

	if res==result{
		return true 
	}

return  false }
