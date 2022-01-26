package dto

//GruplarTableName table name
var GruplarTableName string = "gruplar"

//Kurban dto
type Gruplar struct {
	ID                 uint64
	HayvanBilgisiID    uint64
	KesimSiraNo        int
	HissedarAdet       int
	KurbanFiyatiTipi   int
	ToplamKurbanFiyati float64
	Kilo               int
	//Kurban             []Kurban
}

//TableName override
func (gk *Gruplar) TableName() string {
	return GruplarTableName
}

//Kurban dto
type KurbanUpdateRead struct {
	ID           uint64
	KurbanFiyati float64
	KasaBorcu    float64
	KalanUcret   float64
	BorcDurum    int
	Agirlik      int
}
