# Go Context
## Apa itu Context?
- Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout, dan sinyal deadline
- Context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request)
- Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses

## Kenapa perlu Context?
- Context di golang biasa digunakan untuk mengirim data request atau sinyal ke proses lain
- Dengan menggunakan context, ketika kita ingin membatalkan semua proses, kita cukup mengirim sinyal ke context, maka secara otomatis semua proses akan dibatalkan
- Hampir semua bagian di Golang memanfaatkan context, seperti database, http server, http client, dan sebagainya
- Bahkan di Google sendiri, ketika menggunakan Golang, context wajib digunakan dan selalu dikirim ke setiap function yang dikirim

## Cara kerja Context
- Dalam sebuah proses, misal ada proses A. proses A mengirim Context ke proses B, dan ke proses C
- jika Context di proses A mengirim data sinyal untuk batalkan proses, maka dia akan membatalkan proses B atau C
- Context adalah sebatas data saja

## Package Context
- Context direpresentasikan di dalam sebuah interface Context
- interface Context terdapat di dalam package context
- Docs : https://golang.org/pkg/context

## Membuat Context
- Karena Context adalah sebuah interface, maka kita butuh struct yang sesuai dengan kontrak interface Context
- Kita tidak perlu secara manual karena di Go sudah ada package context

## Function untuk membuat context
| **Function**             | **Keterangan**                                                                                                                                                                                 |
|--------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **context.Background()** | Membuat context kosong, tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki value apapun. Biasa dipakai di main function atau dalam test, atau dalam awal proses request terjadi |
| **context.TODO()**       | Membuat context kosong, namun biasanya menggunakan ini ketika belum jelas context apa yang ingin digunakan                                                                                             |

## Parent dan Child Context
- Context menganut konsep parent dan child
- Artinya, ketika kita buat context, kita bisa buat juga child-nya
- Parent context bisa memiliki banyak child, namun satu child hanya bisa punya satu parent
- Child context juga bisa memiliki child lagi, disebut-nya sub-child
- Semua fitur/data di parent context akan diwariskan ke semua child context
- Semua fitur/data yang dipunyai child, belum tentu dimiliki oleh parent atau child lain
- Konsep ini mirip inheritance di OOP
- Jika ada suatu parent context di-cancel, maka semua child-nya pun akan ikut ter-cancel
- Jika suatu child context dibatalkan, maka semua sub-child yang menjadi child dari child tersebut juga akan dibatalkan

### Immutable
- Context merupakan object yang immutable, artinya setelah Context dibuat, dia tidak bisa diubah lagi
- Ketika kita menambahkan value ke dalam context, atau menambahkan pengaturan timeout dll, secara otomatis akan membentuk child context baru, bukan mengubah context tersebut

## Context with Value
- Pada saat awal membuat Context, dia tidak memiliki value
- Kita bisa menambahkan value ke dalam context dengan menggunakan function **context.WithValue(parent, key, value)**
- Value yang ditambahkan ke dalam context bersifat pair key-value
- Saat kita menambah value ke context, secara otomatis akan tercipta child context baru, artinya original context tidak akan berubah

## Context with Cancel
- Context bisa dibatalkan dengan menggunakan function **context.WithCancel(parent)**
- Context with Cancel biasa digunakan ketika kita menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke proses tersebut
- Biasanya proses ini berupa Goroutine yang berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi Goroutine, kita tinggal kirim sinyal cancel ke context-nya
- Perlu diingat, kita harus pastikan Goroutine yang menggunakan Context harus melakukan pengecekan terhadap Context-nya, jika tidak, tidak ada gunanya

## Context with Timeout
- Context bisa ditambahkan timeout dengan menggunakan function **context.WithTimeout(parent, duration)**
- Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual untuk membatalkan proses, cancel akan otomatis terjadi jika waktu timeout tercapai
- Penggunaan timeout dengan context biasanya digunakan untuk menghindari infinite loop
- Timeout dengan context juga sangat cocok ketika kita query ke database atau API, karena jika tidak kita hentikan, maka akan terlalu lama menunggu

## Context with Deadline
- Context bisa ditambahkan deadline dengan menggunakan function **context.WithDeadline(parent, time)**
- Pengaturan deadline sedikit berbeda dengan timeout, jika waktu timeout adalah hitungan waktu terhadap waktu kita sekarang, maka deadline adalah waktu kapanpun kita mau, misal jam 12 siang besok
- Deadline dengan context juga cocok untuk menghindari infinite loop
  
