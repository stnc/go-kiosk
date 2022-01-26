package dto

//hayvan kartı olacak - ahırı - hayvan kilosu falan olacak hayvanın alış fiyatı satış fiyatı et kilos gibi hesaplamalar olacak

//HayvanBilgisi yapı
type HayvanBilgisi struct {
	ID          uint64 `json:"id"`
	HayvanCinsi string `json:"hayvanCinsi"`
	Agirlik     int    `json:"agirlik"`
	KupeNo      uint64 `json:"kupeNo"`
	Resim       string `json:"resim" `
	Durum       int    `json:"durum"`
	// Fiyat11              float64            `gorm:"-" json:"fiyat1"` //db den okumaz
	Fiyat1               float64            `json:"fiyat1"`
	Fiyat2               float64            `json:"fiyat2"`
	Fiyat3               float64            `json:"fiyat3"`
	SatisFiyati1         float64            `json:"satisFiyati1"`
	SatisFiyati2         float64            `json:"satisFiyati2"`
	SatisFiyati3         float64            `json:"satisFiyati3"`
	KisiBasiDusen1       float64            `json:"kisiBasiDusen1"`
	KisiBasiDusen2       float64            `json:"kisiBasiDusen2"`
	KisiBasiDusen3       float64            `json:"kisiBasiDusen3"`
	KisiBasiKilo         float64            `json:"kisiBasiKilo"`
	HayvanSatisYerleri   HayvanSatisYerleri `json:"hayvanSatisYerleri"`
	HayvanSatisYerleriID uint64             `json:"hayvanSatisYerleriID"`
}

//HayvanSatisYerleri strcut
type HayvanSatisYerleri struct {
	ID         uint64 `json:"id"`
	YerAdi     string `json:"yer_adi"`
	Adresi     string `json:"adresi"`
	IlgiliKisi string `json:"ilgiliKisi"`
	Telefon    string `json:"telefon"`
	Durum      int    `json:"durum"`
}
