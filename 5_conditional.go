package main

import "fmt"

func main()  {
  point := 5

  // penggunaan if,else if, else
  if point > 10 {
    fmt.Println("sangat sempurna")
    }  else if point == 10 {
      fmt.Println("sempurna")
    } else if point >= 7 {
      fmt.Println("lulus")
    } else {
      fmt.Println("gagal")
    }

    // penggunaan switch - case - default
    switch point {
    case 10:
      fmt.Println("sangat sempurna")
    case 7,8,9:
      fmt.Println("lulus")
    case 6,5,3,2,1:
      fmt.Println("gagal")
    default:
      {
        fmt.Println("you need to learn more !")
      }
    }

    // penggunaan switch - case - default dengan clause seperti if else
    switch {
    case point == 10:
      fmt.Println("sangat sempurna")
    case point >= 7 :
      fmt.Println("lulus")
    case point <= 6:
      fmt.Println("gagal")
    default:
      {
        fmt.Println("you need to learn more !")
      }
    }

    // penggunaan switch - case - default dengan fallthrough
    /* note : dalam bahasa lain ketika sebuah case terpenuhi maka
    pengecekkan kondisi tidak akan diteruskan ke case-case berikutnya,
    dalam golang memiliki keyword fallthrough untuk meneruskan ke case-case
    berikutnya meskipun case sebelumnya sudah terpenuhi */
    switch {
    case point == 10:
      fmt.Println("sangat sempurna")
    case point >= 7  :
      fmt.Println("lulus")
    case (point <= 6) && (point > 3):
      fmt.Println("gagal")
      fallthrough
    case point == 5:
      fmt.Println("selamat anda mendapat jackpot !")
    default:
      {
        fmt.Println("you need to learn more !")
      }
    }

    // ternary operator atau if pendek
    /* var results = (point == 5) ? "tebakan anda benar" : "tebakan anda salah"
    fmt.Print(results) */
     // Oops, maaf sayangnya golang tidak mendukung untuk pengecekkan kondisi dengan ternary


}
