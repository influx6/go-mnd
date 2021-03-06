package argument

import (
	"context"
	"math"
	"net/http"
	"strings"
	"time"
)

func example() {
	math.Abs(9.5) // want "Magic number: 9.5"
}

func example2() {
	strings.ToUpper("foo")
}

type s struct {
	count int
}

func (s *s) Add(count int) {
	s.count += count
}

func example3() {
	s := &s{}
	s.Add(10) // want "Magic number: 10"
}

func example4() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second) // want "Magic number: 5"
	defer cancelFunc()

	ctx.Value("foo")
}

func example5() {
	http.StatusText(200) // want "Magic number: 200"
}
