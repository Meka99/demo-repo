package main 

import "fmt"

func main () {


   for i:= 1; i < 101; i++ {
     fmt.Println(i)
   }  


   for j:= 1; j < 51; j++ {
     if j % 2 != 1 {
         continue
     }
   fmt.Println(j)
   }

   for k:= 1; k < 51; k++ {
     if k % 2 == 1 {
         continue
     }
   fmt.Println(k)
   }

   for l:= 50; l < 106; l++ {
       m := l/6
     fmt.Println(m)

   }
    var text string
    fmt.Print ("Enter The text : ")
    fmt.Scan(&text) 
    
    
    if (text == "golangtutorial") {
        fmt.Println("welcome")
    } else {
    fmt.Println("end")
    }
    
//     for n:= 1; n<81; n++ {
//     if n % 2 == 0 {
//         fmt.Println("Golang")
//   } else {
//         fmt.Println(n)
        
//     }  
    
//   }
    for o:= 1; o<81; o++ {
        if o % 4 == 0 {
            if  o % 2 == 0 {
            fmt.Println("Golang tutorial")
            }else {
                fmt.Println("tutorial")
            }
            
       } else if o % 2 == 0 {
            if  o % 4 == 0 {
            fmt.Println("Golang tutorial")
            }else {
                fmt.Println("Golang")
            }
            
      }else {
            fmt.Println(o)
        }  
    
  }
  
  
  
  
  
  
}