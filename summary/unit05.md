# 05. Hello, World! 시작하기

```
~/temp/demosw/study_golang_language/workspace/ch03> tree hello_world
hello_world
├── bin
├── pkg
└── src
   └── hello
       └── hello.go
```

```go
package main

import "fmt"

func main(){
  fmt.Println("Hello, World!")
}
```

```bash
~/temp/demosw/study_golang_language/workspace/ch03/hello_world/src/hello> go build hello.go
~/temp/demosw/study_golang_language/workspace/ch03/hello_world/src/hello   master ●  tree ../../../hello_world
../../../hello_world
├── bin
├── pkg
└── src
    └── hello
        ├── hello
        └── hello.go

4 directories, 2 files
~/temp/demosw/study_golang_language/workspace/ch03/hello_world/src/hello> go build hello.go
~/temp/demosw/study_golang_language/workspace/ch03/hello_world/src/hello> ./hello
Hello, World!
```
