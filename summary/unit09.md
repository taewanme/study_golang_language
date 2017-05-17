#  09. 문자열 사용하기

- 문자열 길이 함수: len()

```
package main

import "fmt"
import "unicode/utf8"

func main(){
    	var s1 = "hello"
	var s2 = "김태완"
	var s3 = "김 태완"

	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(len(s3))

	fmt.Println("---------------")

	fmt.Println(utf8.RuneCountInString(s1))
	fmt.Println(utf8.RuneCountInString(s2))
	fmt.Println(utf8.RuneCountInString(s3))
}

```

- utf8의 문자열 길이
  - import "unicode/utf8"
  - utf8.RuneCountInString()


## 9.3 문자열 수정

- 문자열 수정은 컴파일 에러
