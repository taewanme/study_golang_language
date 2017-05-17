# 18.goto 사용하기

```
goto  LABEL

LABEL:
```

- 중복 코드
```
package main

import "fmt"

func main(){
  a := 1

  if a == 1 {
    fmt.Println("Error 1")
    return
  }

  if a == 2 {
    fmt.Println("Error 2")
    return
  }

  if a == 3 {
    fmt.Println("Error 3")
    return
  }

  fmt.Println(a)
  fmt.Println("Success")
  return
}
```

- 개선코드
```
package main

import "fmt"

func main(){
  a := 1

  if a == 1 {
    goto ERROR1
  }

  if a == 2 {
    goto ERROR2
  }

  if a == 3 {
    goto ERROR3
  }

  fmt.Println(a)
  fmt.Println("Success")
  return

  ERROR1:
    fmt.Pritnln("Error 1")
    return

  ERROR2:
    fmt.Pritnln("Error 2")
    return


}
```
