package main

import "fmt"

func partition(arr []int,start ,end int)(p int)  {

	if start > end {
		return
	}

	i,j,temp := start,end ,arr[start]

	for i<j  {

		for i < j && arr[j] >temp  {
			j--
		}

		for i < j && arr[i] <= temp  {
			i++
		}

		if i < j {
			arr[i] ,arr[j] = arr[j],arr[i]
		}

	}

	arr[i],arr[start] = arr[start],arr[i]

	return i

}

func partition2(arr []int,start,end int)( int){
	left := start
	temp := arr[start]
	right := end

	for left <right{
		for temp >= arr[left]{
			left ++
		}
		for temp < arr[right]{
			right --
		}
		if left >= right{
			break
		}
		arr[left],arr[right] = arr[right],arr[left]
		left ++
		right --
	}
	fmt.Println(arr)
	arr[start],arr[left] = arr[left],arr[start]

	return left
}

func main()  {
	l := []int{5,1,2,3,6,4}
	fmt.Println(partition2(l,0,len(l)-1),l)
}