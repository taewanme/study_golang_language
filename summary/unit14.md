# 14. 패키지 사용하기

```
package main

import "fmt"

func main(){
  fmt.Println("Hello, World")
}
```

```
package main

import "fmt"
import "runtime"

func main() {
  fmt.Println("cpu count: ", runtime.NumCPU() )
}
```

- Global import
```
package main

import (
	. "fmt"
	"runtime"
)

func main(){
	Println("CPU:", runtime.NumCPU())
}
```

- 별칭
  - import f "fmt"
  - import _ "fmt" //임포트 무시, 컴파일 용
