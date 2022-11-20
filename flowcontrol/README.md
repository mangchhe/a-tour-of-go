# 흐름제어

## For

- for 초기화 구문; 조건; 증감식 { ... }
- 초기화 구문과 증감식은 필수가 아니다.
- 세 번째 방식처럼 for문을 while문처럼 활용할 수 있다.
- 네 번째 방식처럼 반복 조건을 생략하면 무한 루프를 만들 수 있다.

```go
func main() {
	var sum int

	// 1.
	for i := 0; i <= 10; i++ {
		sum += i
	}

	// 2.
    for ; sum < 1000; {
        sum += sum
    }

	// 3.
    for sum < 1000 {
        sum += sum
    }
	
	// 4.
	for {
	}

	fmt.Println(sum)
}
```

## if

- `if` 키워드, 조건식 괄호는 생략이 가능하다.

```go
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

## 짧은 구문 if

- 조건문 전에 수행될 짧은 구문을 작성할 수 있다.
- 짧은 구문에서 선언된 변수들은 if scope 내에서만 존재한다.

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

## if와 else

- `else` 키워드

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

## switch

- `switch`, `case`, `default` 키워드
- switch문도 짧은 구문을 작성할 수 있고 자동 break를 지원하기 때문에 작성할 필요가 없다.
- switch case는 상수일 필요가 없다.
- 조건을 적지 않는다면 `switch true {}` 와 동일

```go
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today.")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    }
}
```

## defer

- `defer`은 자신을 둘러싼 함수가 종료될 때까지 함수 실행을 연기한다.
- `defer`로 쌓인 함수들은 후입선출 순서로 실행된다.

```go
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func main() {
    fmt.Println("counting")
    
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("done")
}
```
