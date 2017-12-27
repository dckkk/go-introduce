package main

import "fmt"

func main()  {
  // dalam golang looping hanya dapat menggunakan satu keyword saja yaitu for

  // looping menggunakan for tipe 1
  for i := 0; i < 3; i++ {
    fmt.Println("Angka", i)
  }

  // looping menggunakan for tipe 2
  var a = 0
  for a < 5 {
    fmt.Println("Nomor", a)
    a++
  }

  // looping menggunakan for tipe 3
  var b = 0
  for {
    fmt.Println("Angka", b)
    b++
    
    if b == 5 {
      break
    }
  }

  // looping  menggunakan for-range dalam array
  var example = [3]string{"apple", "orange", "banana"}

  for i, ex := range example {
    fmt.Printf("Matriks %d : %s\n", i, ex)
  }

}
