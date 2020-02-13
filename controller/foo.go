package foo

import (
    "fmt"
    "net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "foo")
}
