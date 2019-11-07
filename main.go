package main

import (
	"fmt"
	"seq"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(seq.Fibo(6))

	s1 := []string{"hello", "world!"}
	fmt.Println(strings.Join(s1, " "))

	s2 := strings.Split("Hello, world", " ")
	fmt.Println(s2[1])
	s3 := strings.Fields("Hello, world!")
	fmt.Println(s3[1])

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	s4 := strings.FieldsFunc("Hello안녕Hello", f)
	fmt.Println(s4)

	fmt.Println(strings.Repeat("hello", 3))

	fmt.Println(strings.Replace("Hello, world!", "world", "Sex", 1))
	fmt.Println(strings.Replace("Hello Hello", "llo", "sex", 2))
}
