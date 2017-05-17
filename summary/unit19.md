# switch 분기문 사용하기

```
switch 변수 {
  case 값1:
    //코드
  case 값2:
    //코드
  case 값3:
    //코드
  default:
    //코드
}
```

- break 필요 없음
- 값은 숫자와 문자 모두 가능
- switch에서 case 실행후 다음 case문으로 넘어가는 키워드
  - fallthrough
- case문에 여러 값 지정 가능

```
i := 3

switch i{
  case 2, 4, 6:
    fmt.Println("짝수")
  case 1, 3, 5:
    fmt.Println("홀수")
}

```

- case 절에 조건식이 위치할 수 있음

```
ackage main

import "fmt"

func main(){
	var a = 3
	switch { //case에 조건문 사용시 변수 지정 없음
	case a >= 5 && a< 10:
		fmt.Println("5~9")
	case a >=0 && a<5:
		fmt.Println("0~4")
	}


}

```

- 패지지
  - math/rand
    - Seed
    - intn
      - 범위: 0부터 지정값까지
  - time
    - Now
    - UnixNano

```

```
