# 01. Go 언어란?

- go 언어는 구글이 개발한 언어
- from 2007
- Developer
  - 켄톰슨
  - 롭파이크
  - 로버트 그리즈머
  - 켄 톰슨

- go 언어의 특징
  - 정적 타입
  - 강타입
  - 가비지 컬렉션
  - 병행성
  - 멀티코더 지원
  - 모듈화 및 패지키 시스템
  - 빠른 컴파일ㄹ

- Weakly-typed & Strongly-typed
  - 타입의 변환이 가능한 언어
  - 컴파일 시점 혹은 런타입에 정해진 규칙에 따라서 형변환을 지원하는 언어
  - Strongly-typed는 값 자체가 타입

```go
var a int = 1
var b float32 = 1.3
var c float32 = a + b // 컴파일 에러 발생
```
