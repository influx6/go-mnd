package assign

import (
	"net/http"
	"time"
)

func example() interface{} {
	s := struct {
		Amount int
	}{
		Amount: 100, // want "Magic number: 100"
	}

	return s
}

func example2() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second, // want "Magic number: 5"
	}
}

func example3() *http.Client {
	c := &http.Client{
		Timeout: 2 * time.Second, // want "Magic number: 2"
	}

	return c
}
