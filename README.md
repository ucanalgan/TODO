# TODO

Terminal üzerinden çalışan, Go ile yazılmış basit bir görev yöneticisi. Görevler JSON dosyasına kaydedilir, program kapanıp açıldığında kaldığı yerden devam eder.

## Özellikler

- Görev ekleme, listeleme, düzenleme ve silme
- Görevleri tamamlandı olarak işaretleme
- Silme işleminde onay sorusu
- JSON dosyasına otomatik kayıt her değişiklik anında yazılır
- Renkli durum göstergesi (tamamlanan yeşil, bekleyen sarı)
- Listede toplam / tamamlanan / bekleyen sayısı
- Geçersiz girdilere karşı doğrulama (boş başlık, hatalı ID, sayı olmayan menü seçimi)

## Kurulum

Go 1.26 veya üzeri gerekir.

```bash
git clone https://github.com/ucanalgan/TODO.git
cd TODO
go run .
```

Derlenmiş bir çalıştırılabilir dosya oluşturmak için:

```bash
go build .
```

> `go run main.go` çalışmaz proje birden fazla dosyaya bölündüğü için tüm paketi derleyen `go run .` kullanılmalıdır.

## Kullanım

Program açıldığında menü görünür:

```
1. Add Mission
2. List Missions
3. Complete Mission
4. Delete Mission
5. Edit Mission
6. Exit
Enter your choice:
```

Örnek bir listeleme çıktısı:

```
Missions:
Mission ID: 1, Title: Go dokümantasyonunu oku, Status: Completed
Mission ID: 2, Title: README yaz, Status: Incomplete
Total Missions: 2, Completed: 1, Incomplete: 1
```

Görevler çalışma dizinindeki `missions.json` dosyasında tutulur. Dosya yoksa program boş bir listeyle başlar ve ilk görev eklendiğinde dosyayı oluşturur.

## Proje Yapısı

```
.
├── main.go       # Menü döngüsü ve komut yönlendirmesi
├── mission.go    # Mission veri modeli
├── input.go      # Terminal girdisi okuma ve doğrulama
├── storage.go    # JSON dosyasına kaydetme ve okuma
└── missions.json # Görev verisi (otomatik oluşur, sürüm kontrolüne dahil değil)
```

Dört dosya da `package main` altındadır; ayrı paket bölünmesi bu boyuttaki bir proje için gereksiz görülmüştür.

### Veri modeli

```go
type Mission struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```

ID'ler program başlarken mevcut en yüksek ID'nin bir fazlasından devam eder; bu sayede silme işleminden sonra ID çakışması oluşmaz.

## Bilinen sınırlamalar

- Görev başlıkları tek satırlıdır
- Silinen bir görev geri alınamaz (silmeden önce onay sorulur)
- Aynı anda birden fazla program örneği çalıştırılırsa dosya yazımları çakışabilir

## Yol haritası

- [ ] Öncelik alanı ve önceliğe göre sıralama
- [ ] Tamamlananları gizleme / filtreleme
- [ ] Başlığa göre arama
- [ ] Oluşturulma tarihi
- [ ] Menü yerine CLI komutları (`todo add`, `todo list`)
- [ ] Birim testleri
