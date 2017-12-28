package main

import "fmt"

func main()  {
  //string inline
  var firstName string = "Dicky"

  //string
  var lastName string
  lastName = "Pratama"

  //string inference menggunakan
  var agama = "islam"
  //string inference menggunakan := dan hanya bisa didalam func
  gender := "laki-laki"

  //multiple string
  var kota, provinsi, negara string
  kota, provinsi, negara = "Depok", "Jawa Barat", "Indonesia"

  fmt.Printf("halo %s %s! Jenis Kelaminnya %s dan agamanya %s \n", firstName, lastName, gender, agama)

  fmt.Println("Saya tinggal di", kota, provinsi, negara, "!")

  // underscore variable
  /* Golang memiliki aturan unik yaitu tidak diperbolehkan ada variable yang
  menganggur (tidak terpakai), jika ada maka program akan error.
  Underscore (_) adalah predefined variabel yang bisa dimanfaatkan untuk
  menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan
  keranjang sampah. Berikut contohnya : */
  _ = "ini boleh gadipakai"
  name, _ := "dicky", "pratama"
  /* Pada contoh di atas, variabel name akan berisikan text john, sedang nilai
  wick ditampung oleh variabel underscore, menandakan bahwa nilai tersebut
  tidak akan digunakan.+ Variabel underscore adalah predefined, jadi tidak
  perlu menggunakan := untuk pengisian nilai, cukup dengan = saja.
  Namun khusus untuk pengisian nilai multi variabel yang dilakukan dengan
  metode type inference, boleh didalamnya terdapat variabel underscore. Biasanya
   variabel underscore sering dimanfaatkan untuk menampung nilai balik fungsi
   yang tidak digunakan.*/

}
