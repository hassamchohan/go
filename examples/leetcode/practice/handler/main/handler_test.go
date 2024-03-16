package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {

	tests := []struct {
		name  string
		value string
		want  int
		err   string
	}{
		{name: "double of two", value: "2", want: 4},
		{name: "double of three", value: "3", want: 6},
		{name: "missing value", value: "", err: "missing parameter"},
		{name: "not a number", value: "a", err: "not a number: a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/double?param="+tt.value, nil)
			if err != nil {
				t.Fatalf("could not create request : %v", err)
			}

			rec := httptest.NewRecorder()
			doubleHandler(rec, req)
			res := rec.Result()
			defer res.Body.Close()

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response : %v", err)
			}

			if tt.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request, got : %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tt.err {
					t.Errorf("expected message %q, got  %q", tt.err, msg)
				}

				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK, got  %v", res.Status)
			}
			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("expected an integer, got : %s", b)
			}
			if d != tt.want {
				t.Errorf("expected double to be %v, got %v", tt.want, d)
			}

		})
	}

}
