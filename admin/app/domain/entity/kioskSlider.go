package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)

//Kiosk strcut
type KioskSlider struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint64     `gorm:"not null;" json:"user_id"`
	Title     string     `gorm:"size:255 ;not null;" validate:"required" json:"title" `
	Excerpt   string     `gorm:"type:text ;" json:"short_content"`
	MenuOrder int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required,numeric"`
	Timer     int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'10'" validate:"required,numeric"`
	Type      int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"numeric"`
	Status    int        `gorm:"type:smallint ;NOT NULL;DEFAULT:'1'" validate:"required"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//BeforeSave init
func (f *KioskSlider) BeforeSave() {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
	f.Excerpt = html.EscapeString(strings.TrimSpace(f.Excerpt))
}

/*
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
 }
*/

//Prepare init
func (f *KioskSlider) Prepare() {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

//Validate fluent validation
func (f *KioskSlider) Validate() map[string]string {
	var (
		validate *validator.Validate
		uni      *ut.UniversalTranslator
	)
	tr := en.New()
	uni = ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")
	validate = validator.New()
	tr_translations.RegisterDefaultTranslations(validate, trans)

	errorLog := make(map[string]string)

	err := validate.Struct(f)
	fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			// can translate each error one at a time.
			lng := strings.Replace(e.Translate(trans), e.Field(), "Burası", 1)
			errorLog[e.Field()+"_error"] = e.Translate(trans)
			// errorLog[e.Field()] = e.Translate(trans)
			errorLog[e.Field()] = lng
			errorLog[e.Field()+"_valid"] = "is-invalid"
		}
	}
	return errorLog
}
