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

}
