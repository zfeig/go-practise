package main
import "fmt"
type Book struct {
 name int
 page int
}

func main() {
 book := Book{26,357}
 fmt.Println("BOOk:",book)
 p := &book
 p.name = 100
 fmt.Println("After Book:",book)
}
