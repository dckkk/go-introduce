package main

import "fmt"

func main()  {
  /* Map adalah tipe data asosiatif yang ada di Golang, berbentuk
  key-value. Untuk setiap data (atau value) yang disimpan, disiapkan juga
  key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier)
  untuk pengaksesan value yang bersangkutan.
  Kalau dilihat, map mirip seperti slice, hanya saja indeks yang digunakan
  untuk pengaksesan bisa ditentukan sendiri tipe-nya
  (indeks tersebut adalah key). */

  // penggunaan map tanpa inisialisasi
  var color map[string]int
  color = map[string]int{}

  color["red"] = 10
  color["yellow"] = 50

  fmt.Println(color)
  /* note : Pengisian data pada map bersifat overwrite,
  ketika variabel sudah memiliki item dengan key yang sama,
  maka value lama akan ditimpa dengan value baru. */

  // pengunaan map dengan inisialisasi
  var fruits = map[string]int{"mango": 10, "apple": 5}
  fmt.Println(fruits)
  // macam macam tipe inisialisasi map yang dapat dilakukan di Golang
  /* var fruits3 = map[string]int{}
  var fruits4 = make(map[string]int)
  var fruits5 = *new(map[string]int) */

  // looping pada map menggunakan for-range
  var months = map[string]int{"januari": 5, "maret": 6, "mei": 10}
  for i, month := range months {
    fmt.Println(i, " \t:", month)
  }

  // penggunaan fungsi delete() pada map
  // delete() digunakan untuk menghapus elemen pada map dengan key tertentu
  var days = map[string]int{"senin": 10, "minggu": 5, "sabtu": 7}
  delete(days, "senin")
  fmt.Println(days)

  // kombinasi slice & map
  // Slice dan map bisa dikombinasikan, dan sering digunakan pada banyak kasus
  var sliceMap = []map[string]string{
    {"name": "Dicky Pratama", "kelas": "X1"},
    {"name": "Dicky P", "kelas": "X2"},
    {"name": "Dicky Prtm", "kelas": "X1"},
  }
  fmt.Println(sliceMap)

}
