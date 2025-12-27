# TorScraper - Tor Network Automation Tool

Bu proje, **Siber Tehdit İstihbaratı (CTI)** süreçlerinde kullanılmak üzere, Tor ağı üzerindeki .onion uzantılı siteleri otomatize bir şekilde tarayan, durumlarını kontrol eden ve anonim olarak veri toplayan bir Go (Golang) uygulamasıdır.

Bu araç, üniversite siber güvenlik projesi kapsamında **Collection (Toplama)** ve **Automation (Otomasyon)** yetkinliklerini pekiştirmek amacıyla geliştirilmiştir.

* **Anonim Erişim:** Trafiği yerel Tor SOCKS5 proxy (127.0.0.1:9150) üzerinden yönlendirerek gerçek IP adresini gizler.
* **Toplu Tarama:** `targets.yaml` dosyasındaki hedef listesini okur ve sırayla işler (Input Handler).
* **Hata Yönetimi:** Kapanmış veya erişilemeyen (dead) siteleri tespit eder, programı durdurmadan loglar ve bir sonraki hedefe geçer.
* **Durum Raporlama:** Sitelerin HTTP durum kodlarını (Status Code) terminale canlı olarak raporlar.

Projeyi çalıştırmadan önce bilgisayarınızda şunların kurulu olması gerekir:

1.  **Go (Golang):** (https://go.dev/dl/)
2.  **Tor Browser:** Tor ağına bağlanmak için arka planda çalışıyor olmalıdır. (Varsayılan port: 9150)

## Kurulum

1.  Projeyi bilgisayarınıza klonlayın veya indirin.
2.  Proje dizinine gidin:
    ```bash
    cd TorScraper
    ```
3.  Gerekli Go kütüphanesini (`x/net/proxy`) indirin:
    ```bash
    go get golang.org/x/net/proxy
    ```

## Kullanım

1.  **Tor Browser'ı açın** ve bağlanmasını bekleyin (Arka planda açık kalmalı).
2.  `targets.yaml` dosyasını düzenleyin ve taramak istediğiniz .onion adreslerini ekleyin.
3.  Uygulamayı çalıştırın:
    ```bash
    go run main.go
    ```