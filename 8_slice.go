package main

import "fmt"

func main()  {
  /* Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga
  dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
  Karena merupakan data reference, menjadikan perubahan data di tiap elemen
  slice akan berdampak pada slice lain yang memiliki alamat memori yang sama. */

  // contoh iniliasisasi slice (sangat mirip dengan array)
  var slices = []string{"apple", "mango", "banana"}
  fmt.Println(slices)

  var fruitsA = []string{"apple", "grape", "starfruit"} // slice
  var fruitsB = [3]string{"apple", "grape", "starfruit"} // array
  var fruitsC = [...]string{"apple", "grape", "starfruit"} // array
  fmt.Println(fruitsA, fruitsB, fruitsC)

  /* Kalau perbedannya hanya di penentuan alokasi pada saat inisialisasi, kenapa
   tidak menggunakan satu istilah saja? atau adakah perbedaan lainnya?
   Sebenarnya slice dan array tidak bisa dibedakan karena merupakan sebuah
   kesatuan. Array adalah kumpulan nilai atau elemen, sedang slice adalah
   referensi tiap elemen tersebut. Slice bisa dibentuk dari array yang sudah
   didefinisikan, caranya dengan memanfaatkan teknik 2 index atau lebih
   untuk mengambil elemen-nya. Seperti berikut. */
   var friend = []string{"zach", "zidane", "zahra"}
   var newFriend = friend[0:2]
   fmt.Println(newFriend)
   /* Kode fruits[0:2] maksudnya adalah pengaksesan elemen dalam slice fruits
   yang dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2. */

   // fungsi len() dan cap()
   /* Seperti sebelumnya len() adalah untuk menghitung jumlah slot dalam array,
   sedangkan cap() digunakan untuk menghitung lebar atau kapasitas maksimum
   slice. Seperti berikut. */
   var colors = []string{"red", "black", "blue", "orange"}
   fmt.Println("len", len(colors)) //len: 4
   fmt.Println("cap", cap(colors)) //cap: 4

   var color = colors[0:3]
   fmt.Println("len", len(color)) //len: 3
   fmt.Println("cap", cap(color)) //cap: 4

   var col = colors[1:4]
   fmt.Println("len", len(col)) //len: 3
   fmt.Println("cap", cap(col)) //cap: 3

   // fungsi append()
   // append() digunakan untuk menambahkan elemen kedalam slice
   var aFruits = []string{"mango", "pineapple", "melon"}
   var cFruits = append(aFruits, "papaya")
   fmt.Println(cFruits)

   // fungsi copy
   /* copy() digunakan untuk men-copy elemen pada parameter kedua dan
   menggabungkannya ke parameter pertama*/
   var oldFruits = []string{"apple"}
   var newFruits = []string{"melon", "watermelon"}
   var copiedFruits = copy(oldFruits, newFruits)
   fmt.Println(oldFruits, copiedFruits)

}
