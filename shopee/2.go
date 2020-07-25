package main

func find(nums1 []int,nums2 []int) (max int) {
	l1 := len(nums1)
	l2 := len(nums2)
	max = 0
	for i:=l1-1;i>=0;i--{
		for j:=0;j<l2;j++{
			if nums1[i] <= nums2[j]{
				max = compare(i+1+l2-j,max)
			}
		}
	}
	return max
}

func compare(x,y int) int {
	if x > y{
		return x
	}

	return y
}

func main()  {
	var a,b int

}