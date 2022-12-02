# 메서드

## Value Receiver

- `func (v Vertex) 함수명()` : 함수명 앞에 붙는 걸 Receiver 라고 부른다.
- 리시버가 붙은 함수는 리시버 타입의 메서드가 된다.

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

- 구조체가 아닌 형식에 대해서도 메서드를 선언할 수 있다.

```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
		return float64(-f)
    }
    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
```

## Pointer Receiver

- 리시버 유형으로 `*T`를 가질 수 있다.
- 포인터 리시버는 리시버가 가리키는 값을 메서드 내에서 수정할 수 있다.

```go
type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{3, 4}
	v.Scale(10)
	fmt.Println(v.X, v.Y)
}
```

## 포인터와 함수

- 함수 파라미터도 주소로 전달하게 되면 메서드 내부에서 값을 수정할 수 있다.

```go
type Vertex3 struct {
	X, Y float64
}

func Scale(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	Scale(&v, 10)
	fmt.Println(v)
}
```

## 메서드와 포인터 indireciton

- 함수에서 포인터 인자를 받기 위해서는 `&`를 사용해야 한다.
- 하지만 포인터 리시버는 포인터가 아니라 값이라도 `(&T)` 포인터 리시버가 있는 메서드를 자동으로 호출한다.

```go
type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex4{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex4{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

- 위와 다르게 포인터를 값 인자로 사용할 경우 역참조가 일어나야 한다.
- 인자에 `*`를 붙여야 하고 리시버 같은 경우에는 `(*T)` 처럼 자동으로 역참조가 일어난다.

```go
func (v Vertex4) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    fmt.Println(p.Abs())
    fmt.Println(AbsFunc(*p))
}
```

## interfaces

- `interface type`은 메서드의 시그니처 집합으로 정의된다.
- `interface` 유형의 값은 해당 메서드를 구현하는 모든 값을 보유할 수 있다.
- `a = v` Vertex 구조체는 포인터 유형에서만 정의되기 때문에 오류를 발생한다.

```go
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex5{3, 4}

	a = f
	a = &v
	a = v // Error

	fmt.Println(a.Abs())
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex5 struct {
	X, Y float64
}

func (v *Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

## 인터페이스의 암시적 구현

```go
type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

## Nil 인터페이스 값

- 인터페이스 자체 내부의 콘트리트 값이 0일 경우, 그 메서드는 nil 리시버로 호출됩니다.

```go
type I2 interface {
	M()
}

type T2 struct {
	S string
}

func (t *T2) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I2

	var t *T2
	i = t
	describe(i)
	i.M()

	i = &T2{"hello"}
	describe(i)
	i.M()
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## Nil 인터페이스 값

- Nil 인터페이스 값은 값 또는 콘크리트 유형을 모두 가지지 않는다.
- Nil 인터페이스에서 메서드를 호출하는 것은 런타임 에러이다. 이유는 어떠한 구체적인 메서드를 호출할지 나타내는 인터페이스 튜플 내부의 타입이 없기 때문이다.

```go
type I3 interface {
	M()
}

func main() {
	var i I
	describe2(i)
	i.M()
}

func describe2(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```