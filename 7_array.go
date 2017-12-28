package main

import "fmt"

func main()  {
  // contoh array sederhana
  var fruits = [4]string{"apple", "banana", "orange", "tomato"}
  fmt.Println("Jumlah array \t\t", len(fruits))
  fmt.Println("isinya \t", fruits)

  // contoh array kosong berisi 3 slot
  var names [3]string
  // cara mengisi array
  names[0] = "uzumaki"
  names[1] = "naruto"
  names[2] = "kun"
  // names[3] = "error" <-- jika dilakukan ini akan error karena jumlah slot array hanya dibatasi 3 diawal
  fmt.Println(names[0], names[1], names[2])

  // contoh array tanpa batasan slot
  var number = [...]int{1,2,3,4,5,6,7}
  fmt.Println(number)

  // contoh array multi-dimensi
  var multi = [2][3]int{{3, 2, 3}, {3, 4, 5}}
  fmt.Println(multi)

  // contoh array menggunakan keyword "make"
  var makes = make([]string, 2)
  makes[0] = "apple"
  makes[1] = "banana"
  fmt.Println(makes) 

  // looping array menggunakan for biasa
  var loops = [5]int{1,2,3,4,5}
  for i:=0; i<len(loops); i++ {
    fmt.Println("elemen", i, loops[i])
  }
  // i = key, loops[i] = value

  // looping array menggunakan for-range
  var samples = [4]int{4,5,6,7}
  for i, sample := range samples {
    fmt.Println("elemen", i, sample)
  }
  // i = key, sample = value

  // looping array dengan underscore(_) variable
  /* hasil looping pada Golang selalu menghasilkan key, dan valuesnya.
  Namun kadang mungkin saja kita hanya membutuhkan salah satunya, hal itu
  tentunya akan menghasilkan error mengingat Golang tidak memperbolehkan adanya
  variable tidak terpakai, solusinya adalah menggunakan underscore(_) variable*/
  var unders = [4]int{2,3,4,5}
  for _, under := range unders {
    fmt.Println(under)
  }


}
