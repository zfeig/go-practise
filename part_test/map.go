package main 
import "fmt"

var (
   cMap = map[string]int{"MonDay":1,"ThuesDay":2,"ThirsDay":3,"FriDay":4,"WendsDay":5,"SatuDay":6,"SunDay":7}
   aMap = map[string]func() int{"aa":func() int {return 100},"bb":func() int {return 200},"cc":func() int {return 300}}
)

func checkDay(day int) (dayName string) {
	for name,value := range cMap {
		if value == day {
           dayName = name
		}
	}
	return dayName
}


func getDayName(day int) string {
	dayName := ""
	for k,v := range cMap {
		if v == day {
           dayName = k
		}
	}
	return dayName
}


func isDayExist(name string) bool {
	res := false
	if _,ok := cMap[name];ok {
      res = true
	}
	return res
}

func main() {
	// aMap := map[string]func() int{
 //        "aa":func() int {return 100},
 //        "bb":func() int {return 200},
 //        "cc":func() int {return 300},
	// }
	fmt.Println(aMap)
	num := aMap["bb"]()
	fmt.Printf("Map bb is: %d\n", num)
	fmt.Printf("ready to range map:\n")
    for k,v := range aMap {
    	fmt.Printf("func name:%s  return value is :%d\n",k, v())
    }

    //make map and int
    bMap := make(map[string]int)
    bMap["aa"] = 1
    bMap["bb"] = 2
    bMap["cc"] = 3
    for k,v := range bMap {
    	fmt.Printf("map name:%s  res is :%d\n",k, v)
    }
   fmt.Printf("dayName is :%s\n",checkDay(5))
   fmt.Printf("dayName is :%s\n",getDayName(6))
   fmt.Printf("Monday is exist: %t\n",isDayExist("MonDay"))
   fmt.Printf("Monday is exist: %t\n",isDayExist("whatDay"))
   
   //reverse map
   rMap := make(map[int]string,len(cMap))
   for k,v := range cMap{
   	  rMap[v] = k 
   }

   for k,v := range rMap{
   	 fmt.Printf("rMap k:%d=>v:%s\n",k,v)
   }

}