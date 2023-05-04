package main
import ("fmt"
         "math")

type Person struct {
  first_Name string
  last_Nmae string
  favorite_Country []string
  
}
  
  type vehicle struct {
  doors int
  color string
  
  }
  
  
  type truck struct {
      vehicle
      fourWheel bool
  }
  
  
  type sedan struct {
      vehicle
      luxury bool
  }
  
  
  type square struct {
      length float64
      width float64
  }
  
  
  type circle struct {
      radius float64
      
  }
func (s square) area() float64 {
      return s.length * s.width
  }
  
  
  
  func (c circle) area() float64 {
      return math.Pi * c.radius * c.radius
  }
  
  type SHAPE interface {
    area() float64
  }
  
  
  func INFO(s SHAPE) {
    fmt.Printf("Area: %f\n", s.area())
}
    
func main() {
  var p1 Person
  var p2 Person
  
  p1.first_Name = "John"
  p1.last_Nmae = "Snow"
  p1.favorite_Country = []string{ "USA", "Austrilia", "Italy"}
  
  fmt.Println("First person",  p1)
  fmt.Println("first_Name", p1)
  for idx, val := range p1.favorite_Country {
     fmt.Printf("%v\t%v\n", idx, val)
  
  }
  
  p2.first_Name = "Tyrion"
  p2.last_Nmae = "Lannister"
  p2.favorite_Country = []string{ "Austria", "Turkey", "Dubai"}
  
  fmt.Println("Second person", p2)
  fmt.Println("first_Name", p2)
  for idx, val := range p2.favorite_Country {
     fmt.Printf("%v\t%v\n", idx, val)
  
  
  }
  
  var a = make(map[string]Person)
  var b = make(map[string]Person)
  a[p1.last_Nmae] = p1
  b[p2.last_Nmae] = p2
  
  for idx, val := range a {
     fmt.Printf("%v\t%v\n", idx, val)
     for idx, val := range p1.favorite_Country {
     fmt.Printf("%v\t%v\n", idx, val)
  
  }
  fmt.Println()
  }
      for idx, val := range b {
         fmt.Printf("%v\t%v\n", idx, val)   
  
  for idx, val := range p2.favorite_Country {
     fmt.Printf("%v\t%v\n", idx, val)
  
  
  }
  fmt.Println()
  
      }
      
  
  var t = truck {
     vehicle : vehicle {
         doors: 3,
         color: "Black" ,
      },
      fourWheel: true,
  }
  
 fmt.Println(t, t.doors)
  
  var s = sedan {
     vehicle : vehicle {
         doors: 5,
         color: "Green" ,
      },
      luxury: true,
  }
  
 fmt.Println(s, s.luxury)
  
  
  square := square {
        length: 2,
        width:  3,
  
  }
  
  circle := circle{
        radius: 4,
    }
  
  INFO(square)

    
    INFO(circle)
  
  }