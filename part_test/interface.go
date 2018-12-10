package main 
import "fmt"

type Shape interface{
	Area() float32
}


type Square struct {
	side float32
}

func (sq Square) Area() float32 {
	return sq.side*sq.side
}

type Rectange struct{
	length,width float32
}

func (re Rectange) Area() float32 {
	return re.length*re.width
}
var aShape Shape

func main() {
	s := Square{5}
	r := Rectange{6,3}

    if _,ok := aShape.(*Square); ok {
    	fmt.Printf("aShape has Square:%b\n",ok)
    } else {
    	fmt.Println("aShape does not has Square")
    }
   
	shapes := []Shape{s,r}
	for k,_:= range shapes {
		fmt.Printf("Shape details: %T =>%v\n", shapes[k],shapes[k])
        fmt.Println("Area of this shape is: ", shapes[k].Area())
	}

	switch t := aShape.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
    case *Rectange:
    	fmt.Printf("Type Rectange %T with value %v\n", t, t)
    case nil:
    	fmt.Printf("nil value: nothing to check?\n")
    default:
    	fmt.Printf("Unexpected type %T\n", t)		
	}
}

