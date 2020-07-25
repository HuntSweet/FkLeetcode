package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string  {
	if s == ""{
		return ""
	}
	strs := strings.Split(s,"")
	strsMap := map[string]string{}
	for i := 0;i<len(strs);i++{
		strsMap[strs[i]] += strs[i]
	}
	strList := []string{}
	for _,v := range strsMap{
		strList = append(strList,v)
	}
	sort.Strings(strList)
	sort.Slice(strList, func(i, j int) bool {
		return len(strList[i]) < len(strList[j])
	})

	return strings.Join(strList,"")
}

func main()  {
	var str string
	fmt.Scan(&str)
	fmt.Println(str)
	ss := frequencySort(str)
	fmt.Println(ss)
}
