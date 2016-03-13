package main

import (
    "fmt"
    "strconv"
    "github.com/go-martini/martini"
)

func fib(n int) int {
    if n < 2 {
        return n
    }
    return fib(n-2) + fib(n-1)
}

func main() {
    app := martini.Classic()
    app.Get("/(?P<number>[0-9]+)", func(params martini.Params) string {
        number, _ := strconv.Atoi(params["number"])
        return fmt.Sprintf("<html>Go + Martini<hr> fib(%s): %s</html>", params["number"], strconv.Itoa(fib(number)))
    })
    app.Run()
}
