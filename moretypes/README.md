# More Type

## Pointers

- 포인터는 값의 메모리 주소를 가진다.
- `*T` 타입은 T 값을 가리키는 포인터이다. zero value는 `nil`이다.
- `&`는 피연산자에 대한 포인터를 생성한다.
- `*`는 포인터가 가리키는 주소의 값을 나타낸다.

```go
func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	fmt.Println(*p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
```

## Structs (구조체)

- 구조체는 필드의 집합체이다.
- `type 구조체명 struct {}`
- 필드에는 .(dot)으로 접근할 수 있다.

```go
type Vertex struct {
	X int
	Y int
}

func main() {
    vertex := Vertex{1, 2}
    fmt.Println(vertex.X, vertex.Y)
    
	// 포인터로도 접근 가능
    vertex2 := &vertex
    vertex2.X = 111
    fmt.Println(vertex.X, vertex.Y)
}
```

## Struct Literals

- 구조체에 값을 할당하는 방법에는 여러 가지 방법이 있다.

```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

## Arrays

- `[n]T` 타입은 타입이 T인 n개의 값의 배열이다.
- 파이썬과 같이 배열 슬라이스를 지원한다. `a[low, high]`

```go
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

## Slices are like references to arrays

- 슬라이싱하여 만들어진 배열은 새로운 주소를 가지는 배열이 아니다.
- 배열을 수정하게 되면 기존 배열도 같이 수정되게 된다.

```go
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

## Slice defaults

- 슬라이스 표현법

```go
func main() {
    var a [10]int
	a[0:10]
	a[:10]
	a[0:]
	a[:]
}
```

## Slice length and capacity

- slice는 length, capacity를 가진다.
- length : 슬라이드가 포함하는 요소의 개수
- capacity : 슬라이스의 첫 번째 요소부터 기본 배열의 요소의 개수

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

## Nil slices

- 슬라이스 zero value는 nil이다.
- nil 슬라이스의 길이와 용량은 0이며, 기본 배열을 가지고 있지 않다.

```go
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

## Creating a slice with make

- 슬라이스는 내장된 `make` 함수로 생성할 수 있다. 이건 동적 크기의 배열을 생성하는 방법이다.

```go
func main() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5) // type, length, capacity
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
```

## Slices of Slices

- 슬라이스는 다른 슬라이스를 포함하여 모든 타입을 담을 수 있다.

```go
func main() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

## Appeding to a slice

- 슬라이스에 새로운 요소를 추가할 때 일반적으로 사용하는 메서드
- `func append(s []T, vs ...T) []T`
  - 첫 번째 : 추가될 배열
  - 두 번째 : 추가해야 할 요소

```go
func main() {
	var s []int
	printSlice3(s)

	s = append(s, 0)
	printSlice3(s)

	s = append(s, 1)
	printSlice3(s)

	s = append(s, 2, 3, 4)
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

## Range

- for문에서 `range`는 맵 또는 슬라이스 요소들을 순회할 때 사용한다.
- `i, v := range <slice>`
  - 첫 번째 : 인덱스
  - 두 번째 : 해당 인덱스의 값의 복사본

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

## Range continued

- _를 이용해 인덱스 또는 값을 생략할 수 있다.

```go
for i, _ := range pow
for _, value := range pow
for i := range pow
```

## Maps

- zero value는 nil이고 nil은 키도 없고 키를 추가할 수도 없다.
- `make` 함수를 이용해 주어진 타입으로 초기화되고 사용 준비가 된 맵을 반환한다.

```go
type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func main() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

## Map literal

- 최상위 타입이 타입 이름일 경우, 리터럴 요소에 생략할 수 있다.

```go
var m = map[string]Vertex{
    "Bell Labs": Vertex{
    40.68433, -74.39967,
},
"Google": Vertex{
    37.42202, -122.08408,
    },
}

var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
```

## Mutating Maps

- m[key] = elem : 추가하거나 업데이트
- elem = m[key] : 요소 검색
- delete(m, key) : 요소 제거
- elem, ok = m[key] : 키가 존재하는지 검색
  - ok : 데이터가 있으면 true 없으면 false
  - elem: 데이터가 없으면 해당 타입 zero value

```go
func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```

## Function values

- 함수들도 값으로 취급하기 때문에 인자나 리턴 값으로 넘겨줄 수 있다.

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(5, 12)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

## Function closures

- 클로저는 함수의 외부로부터 오는 변수를 참조하는 함수 값이다.
- 함수는 참조된 변수에 접근하여 할당될 수 있다.

```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

## Fibonacci closure (예제)

```go
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := second
		first, second = second, first + second
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```