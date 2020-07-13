package sort

//归并排序
func mergeSort(l []int) []int {
	if len(l) <= 1{
		return l
	}

	mid := len(l) / 2
	left := mergeSort(l[:mid])
	right := mergeSort(l[mid:])

	result := merge(left,right)
	return result
}

func merge(l,r []int) (res []int) {
	i := 0
	j := 0

	for i < len(l) && j < len(r){
		if l[i] > l[j]{
			res = append(res,l[j])
			i++
		}else {
			res = append(res,r[j])
			j++
		}
	}

	//不会越界
	res = append(res,l[i:]...)
	res = append(res,r[j:]...)

	return
}