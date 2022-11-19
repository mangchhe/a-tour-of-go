# Go를 향한 여행

## Import

- 괄호로 import를 그룹 짓고 이를 `"factored" import` 문이라고 합니다.

```go
import (
    "fmt"
    "math"
)
```

## Export되는 이름

- 대문자로 시작하는 이름들만 export되고 그렇지 않다면 패키지 밖에서 접근할 수 없다.

```go
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi)
}
```

## 함수

- 변수 이름 뒤에 타입이 오고 리턴 타입도 마찬가지이다. 코틀린이랑 유사
- 같은 타입이 연속될 경우 마지막 매개변수 타입을 제외하고 생략 가능하다.

```go
func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(subtract(42, 13))
}
```

## 복수개의 결과

- 고는 여러 개를 동시에 리턴할 수 있다. 파이썬이랑 유사

```go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

## 이름이 주어진 반환값

- 리턴 값에 이름을 정의할 수 있고 변수처럼 사용된다.
- 인자가 없는 return은 이름이 주어진 리턴 값을 반환하는데 이를 `"naked" return`이라고 한다.

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

## 변수 선언

- `var` 키워드를 이용하여 변수를 선언하고 마지막은 타입을 정의한다.
- 명시적인 초깃값이 없을 경우 `zero value`가 주어진다.
  - 숫자 : 0
  - 불린 : false
  - 문자열 : "" (빈 문자열)

```go
var c, python, java bool

func main() {
    var i, j int
    fmt.Println(i, j, c, python, java)
}
```

## 변수 초기화

- 초깃값이 존재한다면 타입을 생략될 수 있다.

```go
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	var golang ="go!"
	fmt.Println(i, j, c, python, java, golang)
}
```

## 짧은 변수 선언

- `:=`를 이용하여 암시적 타입 선언을 할 수 있다.
- 단, 함수 바깥에서는 `func`, `var` 키워드로 시작하기 때문에 `:=`는 사용 불가능하다.

```go
func main() {
	i, j := 3, 4
	fmt.Println(i, j)
}
```

## 기본 자료형

- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // uint8의 별칭
- rune // int32의 별칭
- float32 float64
- complex64 complex128

```go
var (
	BOOL   bool
	INT    int
	INT8   int8
	INT16  int16
	INT32  int32
	INT64  int64
	UINT   uint
	UINT8  uint8
	UINT16 uint16
	UINT32 uint32
	UINT64 uint64
	BYTE   byte
	RUNE rune
	FLOAT32 float32
	FLOAT64 float64
	COMPLEX64 complex64
	COMPLEX128 complex128
)
```

## 타입 변환

- 다른 타입의 요소들 간의 할당에는 명시적인 변환이 필요하다.

```go
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}
```

## 타입 추론

- `=`, `:=` 오른편에 상수일 경우 그 상수의 정확도에 따라 타입이 정해진다.

```go
func main() {
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
```

## 상수

- 상수는 `const` 키워드를 이용해 선언된다.
- 상수는 `:=`를 통해 선언할 수 없다.

```go
const Pi = 3.14

func main() {
	const World = "World!"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

## 숫자형 상수

- 숫자형 상수는 **매우 정확한 값**이다.

```go
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

