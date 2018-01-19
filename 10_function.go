package main

import "fmt"
import "strings"

/* Fungsi main merupakan fungsi yang paling utama pada program Golang.
Cara membuat fungsi cukup mudah, yaitu dengan menuliskan keyword func, diikuti
setelahnya nama fungsi, kurung yang berisikan parameter, dan kurung kurawal
untuk membungkus blok kode. */

// berikut beberapa contoh penggunaan fungsi
func main()  {
  // memanggil fungsi tanpat return
  var names = []string{"Dicky", "Pratama"}
  printName("halo", names)

  // memanggil fungsi dengan return
  var segitiga = luasSegitiga(2, 3)
  fmt.Println("luas segitiga adalah", segitiga)

  // memanggil fungsi dengan multiple return
  var luasPersegi, kelilingPersegi = hitungPersegi(6)
  fmt.Printf("luas persegi adalah %d, dan kelilingnya %d\n", luasPersegi, kelilingPersegi)

  // memanggil fungsi dengan return di diawal
  var simpleLuas, simpleKeliling = simplePersegi(3)
  fmt.Printf("simple luas persegi adalah %d, dan simple kelilingnya %d\n", simpleLuas, simpleKeliling)

  // memanggil fungsi variadic
  var average = calculate(1,2,3,4,5,1,2,3,3,4,5)
  fmt.Println("rata rata nya adalah", average)

  // fungsi closure, maksudnya adalah fungsi dalam variable
  var allName = func(names []string) string {
    var allNames = strings.Join(names,",")

    return allNames
  }
  var named = allName([]string{"Dicky", "Pratama", "Stevie", "McQueen"})
  fmt.Println(named)

}

// fungsi tanpa return
func printName(msg string, arr []string) {
  var nameString = strings.Join(arr, " ")
  fmt.Println(msg, nameString)
}

// fungsi dengan return
func luasSegitiga(a, t int) int{
  var luas = int (a * t) / 2

  return luas
}

// fungsi dengan multiple return
func hitungPersegi(s int) (int, int){
  luas := s * s
  keliling := 4 * s

  return luas, keliling
}

// fungsi dengan return diawal (ini adalah salah satu keunikan pada golang)
func simplePersegi (s int) (luas int, keliling int) {
  luas = s * s
  keliling = 4 * s

  return
}

// fungsi variadic, maksudnya adalah infinite parameter
func calculate(numbers ...int) float64 {
  total := 0
  for _, number := range numbers {
    total += number
  }
   avg := float64(total) / float64(len(numbers))

   return avg

}
