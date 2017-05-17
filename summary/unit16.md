# if 조건안에 함수 사용

```

var b []byte
var err error

b, err = ioutil.ReadFile("./hello.txt")

if err==nil {
    fmt.Printf("%s", b)
}

```

```
if b, err := iotuil.ReadFile("./hello.txt"); err==nil {
    fmt.Printf("%s", b)
}
```
