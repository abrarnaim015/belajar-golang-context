<h1>Context</h1>

```golang
background := context.Background()
todo := context.TODO()
```

context bg atau background adalah context kosong, sama dengan context todo bedanya hanya jika masih blm jelas mau menggunakan context apa maka gunakan context.TODO(), tapi sanga jarang penggunaan context todo karna biasanya sudah jelas akan menggunakan context apa pada saat akan menggunakannya

> Memasukan data dan mengambil data di context

```golang
contextA := context.Background()

contextB := context.WithValue(contextA, "b", "B")
contextC := context.WithValue(contextA, "c", "C")

fmt.Println("contextA :", contextA)
fmt.Println("contextB :", contextB)
fmt.Println("contextC :", contextC)

fmt.Println("contextA :", contextD.Value("b"))
// data dari contextA saya ambil di contextD
```

context bawah akan selalu membawa context bapaknya, makanya ketidak saya mengambil value di contextD dengan param "b" yang mana do set di contextB saya dapat ambil karena di contextD juga ada karena bapaknya adalah contextB

> context Cancel

cara membuat context Cancel dengan sebagai berikut

```golang
parent := context.Background()
ctx, cancel := context.WithCancel(parent)
```

untuk mengkatifkan func cancel kita harus memanggil `context.Done()` agar balikan value dari `context.WithCancel(parent)` yaitu `cancel` dapat dipanggil dengan cara `cancel()`.

> context Timeout

```golang
parent := context.Background()
ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
```

logic yang digunakan dengan cancel biasa sama, hanya saja kita dapat mengaktifkan cancel dengan ototmatis setelah waktu yang telah di tentikan maka dari itu kita tinggal memasang `defer cancel()` di code yang kita buat agar jika sudah melebihi waktu yang di tentukan proses yang sedang berjalan dapat di hentikan secara otomatis. sangat berguna pada query ke DB

> context Deadline

```golang
parent := context.Background()
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
```

hampir sama dengan timeout bedanya hanya waktu setnya, jika timeout waktu yang di set tapi jika deadline waktu yang di set, contoh jam 12 siang atau jam 1 siang, jadi patokannya adalah jam bukan waktu yang di set seperti 5 dtk atau sebagainya dan pemanggilan `defer cancel()` sama seperti timeout
