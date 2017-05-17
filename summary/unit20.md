# 20. 간단한 예제로 문법 익숙해지기

- 1~100까지 출력
- 3의 배수는 Fizz 출력
- 5의 배수는 Buzz 출력
- 3과 5의 공배수는 FizzBuzz출력

```
package main

import "fmt"

func main(){
	for i:=1; i<=100;i++ {
		switch{
		case (i%3 == 0) && (i%5 ==0):
			fmt.Print("FizzBuzz", ", ")
		case (i%3 == 0):
			fmt.Print("Fizz", ", ")
		case (i%5 == 0):
			fmt.Print("Buzz", ", ")
		default:
			fmt.Print(i,", ")
		}
	}
}
```
