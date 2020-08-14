package main

/**
 *
 * @param s string字符串
 * @return bool布尔型
 */
func IsValidExp( s string ) bool {
	// write code here
	m := map[string]string{"(":")","[":"]","{":"}"}
	l := make([]string,0)


	for _,v := range s{
		if _,ok := m[string(v)];ok{
			l = append(l,string(v))
		}else{

			if len(l) == 0{
				return false
			}else if m[l[len(l)-1]] != string(v){
				return false
			}else if m[l[len(l)-1]] == string(v){
				l = l[:len(l)-1]

			}

		}
	}

	if len(l) > 0{
		return false
	}
	return true
}
