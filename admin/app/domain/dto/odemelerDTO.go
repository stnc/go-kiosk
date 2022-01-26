package dto

import "time"

//OdemelerTableName table name
var OdemelerTableName string = "odemeler"

//Odemeler dto layer
type Odemeler struct {
	ID       uint64 `json:"id"`
	KurbanID uint64 `json:"kurbanId"`
	UserID   uint64 `json:"userId"`
	Aciklama string `json:"aciklama"`
	Makbuz   string `json:"makbuz" `
	//durum=  taksit eklemişse durum 1 / İlk Eklenen Fiyat değeri 2 / kasa borçlu kalmışsa değer 3 olur
	BorcDurum      int       `validate:"required"`
	VerilenUcret   float64   `validate:"required,numeric" `
	KalanUcret     float64   `validate:"numeric" `
	KasaBorcu      float64   `validate:"numeric" `
	VerildigiTarih time.Time `json:"verildigiTarih"`
}

//OdemelerSonFiyat dto layer
type OdemelerSonFiyat struct {
	ID           uint64
	VerilenUcret float64 `validate:"required,numeric"`
	KalanUcret   float64 `validate:"required,numeric"`
	KasaBorcu    float64 `validate:"required,numeric"`
}

//TableName override
func (gk *Odemeler) TableName() string {
	return OdemelerTableName
}
