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
- Konsep ini mirip inheritance di OOP
