# 변수 사용하기

```
var i int
var s string
var age int = 10
var name string = "Taewan Kim"
```

- 형식
  - var {name} {type}
  - {type} 생략가능
  - 대입 값으로 결정

## 짧은 선언
- ```:=```을 사용하여 var 생략 가능

```
age := 10
name := "taewan kim"
```

## 7.2 변수 여러개 선언하기

- var {variable name 1}, {variable name 1} {type} = {init_value}, {init_value}
- var {variable name 1}, {variable name 1} = {init_value}, {init_value}

```
var x, y, int = 30, 50
var age, name = 10, 'Taewan kim'
```

- {variable name 1}, {variable name 1} := {init_value}, {init_value}
```
a, b, c, d := 1, 3.4, "Hello World", false
```

```
var(
  x, y int = 3, 4
  age, name = 10, "taewan kim"
  )
```

- 사용하지 않는 변수와 패키지 처리
  - ```_```를 사용

```
package main

import "fmt"
import _ "runtime"

func main(){
  a := 1
  b := 2
  _ := b
}

```
