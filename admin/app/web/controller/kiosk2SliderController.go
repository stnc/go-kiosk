package controller

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnc2upload"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncdatetime"

	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"
	"strconv"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/disintegration/imaging"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//kioskSlider constructor
type Kiosk2Slider struct {
	Kiosk2SliderApp     services.Kiosk2SliderAppInterface
	userApp             services.UserAppInterface
	categoriesKiosk     services.CategoriesKioskAppInterface
	CategoriesKioskJoin services.CategoriesKioskJoinAppInterface
}

const viewPathkioskSlider2 = "admin/kiosk2Slider/"

//InitkioskSlider kioskSlider controller constructor
func Initkiosk2Slider(KiApp services.Kiosk2SliderAppInterface, uApp services.UserAppInterface,
	catKioskApp services.CategoriesKioskAppInterface, catKioskJoinApp services.CategoriesKioskJoinAppInterface) *Kiosk2Slider {
	return &Kiosk2Slider{
		Kiosk2SliderApp:     KiApp,
		userApp:             uApp,
		categoriesKiosk:     catKioskApp,
		CategoriesKioskJoin: catKioskJoinApp,
	}
}

//Index list
func (access *Kiosk2Slider) Index(c *gin.Context) {
	// allkioskSlider, err := access.Kiosk2SliderApp.GetAllkioskSlider()

	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	var tarih stncdatetime.Inow
	var total int64
	access.Kiosk2SliderApp.Count(&total)
	kioskSlidersPerPage := 10
	paginator := pagination.NewPaginator(c.Request, kioskSlidersPerPage, total)
	offset := paginator.Offset()

	kioskSliders, _ := access.Kiosk2SliderApp.GetAllP(kioskSlidersPerPage, offset)

	// var tarih stncdatetime.Inow
	// fmt.Println(tarih.TarihFullSQL("2020-05-21 05:08"))
	// fmt.Println(tarih.AylarListe("May"))
	// fmt.Println(tarih.Tarih())
	// //	tarih.FormatTarihForMysql("2020-05-17 05:08:40")
	//	tarih.FormatTarihForMysql("2020-05-17 05:08:40")

	viewData := pongo2.Context{
		"paginator": paginator,
		"title":     "İçerik Ekleme",
		"allData":   kioskSliders,
		"tarih":     tarih,
		"flashMsg":  flashMsg,
		"csrf":      csrf.GetToken(c),
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathkioskSlider2+"index.html",
		// Pass the data that the page uses
		viewData,
	)
}

//Create all list f
func (access *Kiosk2Slider) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	cats, _ := access.categoriesKiosk.GetAll()
	viewData := pongo2.Context{
		"title":    "İçerik Ekleme",
		"catsData": cats,
		"csrf":     csrf.GetToken(c),
	}
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathkioskSlider2+"create.html",
		// Pass the data that the page uses
		viewData,
	)
}

//Store save method
func (access *Kiosk2Slider) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var kiosk2Slider, _, _ = kiosk2SliderModel(c)
	var savekioskSliderError = make(map[string]string)

	savekioskSliderError = kiosk2Slider.Validate()

	sendFileName := "Picture"
	filenameForm, _ := c.FormFile(sendFileName)

	//multiple
	// form, _ := c.MultipartForm()
	// files := form.File[sendFileName]
	// stnc2upload.NewFileUpload("str").MultipleUploadFile(files, c.PostForm("Resim2"))

	filename, uploadError := stnc2upload.NewFileUpload().UploadFile(filenameForm, c.PostForm("Resim2"))
	if filename == "false" {
		savekioskSliderError[sendFileName+"_error"] = uploadError
		savekioskSliderError[sendFileName+"_valid"] = "is-invalid"
	}

	catsPost := c.PostFormArray("cats")
	//fmt.Println(catsPost)
	catsData, _ := access.categoriesKiosk.GetAll()
	fmt.Println(reflect.ValueOf(catsData).Kind())
	for key, row := range catsData {
		catsData[key].ID = row.ID
		catsData[key].Name = row.Name
		//a, _ := strconv.Atoi(catsPost[key])
		finding := strconv.FormatInt(int64(row.ID), 10)
		_, found := stnccollection.FindSlice(catsPost, finding)
		if found {
			catsData[key].SelectedID = row.ID
		}
	}

	if len(savekioskSliderError) == 0 {

		_, filetype := stnc2upload.NewFileUpload().RealImage("public/upl/kiosk/" + filename)
		if filetype != "video/mp4" {
			kiosk2Slider.Type = 1
		} else {
			kiosk2Slider.Type = 2
		}
		kiosk2Slider.Picture = filename
		saveData, saveErr := access.Kiosk2SliderApp.Save(&kiosk2Slider)

		if saveErr != nil {
			savekioskSliderError = saveErr
		} else {
			if filetype != "video/mp4" {
				if filenameForm != nil {
					src, err := imaging.Open("public/upl/kiosk/" + filename)
					if err != nil {
						log.Fatalf("failed to open image: %v", err)
					}

					// src = imaging.Resize(src, 1815, 2510, imaging.Center)
					// var srcBig image.Image
					var srcBig = imaging.Resize(src, 1815, 2510, imaging.Lanczos)

					errBig := imaging.Save(srcBig, "public/upl/kiosk/big/"+filename)
					if errBig != nil {
						log.Fatalf("failed to save image: %v", errBig)
					}

					var srcThumb = imaging.Resize(src, 150, 150, imaging.Lanczos)

					errth := imaging.Save(srcThumb, "public/upl/kiosk/thumb/"+filename)
					if errth != nil {
						log.Fatalf("failed to save image: %v", errth)
					}
				}
			}
		}

		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		var catPost = entity.CategoriesKioskJoin{}
		for _, row := range catsPost {
			catID, _ := strconv.ParseUint(row, 10, 64)
			catPost.CategoryID = catID
			catPost.KioskID = saveData.ID
			saveCat, _ := access.CategoriesKioskJoin.Save(&catPost)

			catPost.ID = saveCat.ID + 1
		}
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi, düzenleme ekranından düzenlemeye başlayabilirsiniz", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/kiosk2Slider/edit/"+lastID)
		return
	}

	viewData := pongo2.Context{
		"title":    "içerik ekleme",
		"catsPost": catsPost,
		"catsData": catsData,
		"csrf":     csrf.GetToken(c),
		"err":      savekioskSliderError,
		"data":     kiosk2Slider,
	}
	c.HTML(
		http.StatusOK,
		viewPathkioskSlider2+"create.html",
		viewData,
	)

}

//Edit edit data
func (access *Kiosk2Slider) Edit(c *gin.Context) {
	//strconv.Atoi(c.Param("id"))
	//kioskSliderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	if kioskSliderID, err := strconv.ParseUint(c.Param("KioskSliderID"), 10, 64); err == nil {

		if kioskSliders, err := access.Kiosk2SliderApp.GetByID(kioskSliderID); err == nil {
			var catsPost []string

			catsPostData, _ := access.CategoriesKioskJoin.GetAllforKioskID(kioskSliderID)

			for _, row := range catsPostData {
				str := strconv.FormatUint(row.CategoryID, 10) //uint64 to stringS
				catsPost = append(catsPost, str)
			}
			catsData, _ := access.categoriesKiosk.GetAll()
			for key, row := range catsData {
				catsData[key].ID = row.ID
				catsData[key].Name = row.Name
				//a, _ := strconv.Atoi(catsPost[key])
				finding := strconv.FormatInt(int64(row.ID), 10)
				_, found := stnccollection.FindSlice(catsPost, finding)
				if found {
					catsData[key].SelectedID = row.ID
				}
			}
			viewData := pongo2.Context{
				"title":    "içerik düzenleme",
				"catsPost": catsPost,
				"catsData": catsData,
				"data":     kioskSliders,
				"csrf":     csrf.GetToken(c),
				"flashMsg": flashMsg,
			}
			c.HTML(
				http.StatusOK,
				viewPathkioskSlider2+"edit.html",
				viewData,
			)

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//Update data
func (access *Kiosk2Slider) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	var kiosk2Slider, idN, id = kiosk2SliderModel(c)
	var savekioskSliderError = make(map[string]string)

	savekioskSliderError = kiosk2Slider.Validate()

	sendFileName := "Picture"

	filenameForm, _ := c.FormFile(sendFileName)
	filename, uploadError := stnc2upload.NewFileUpload().UploadFile(filenameForm, c.PostForm("Resim2"))

	if filename == "false" {
		savekioskSliderError[sendFileName+"_error"] = uploadError
		savekioskSliderError[sendFileName+"_valid"] = "is-invalid"
	}

	catsPost := c.PostFormArray("cats")

	catsData, _ := access.categoriesKiosk.GetAll()
	for key, row := range catsData {
		catsData[key].ID = row.ID
		catsData[key].Name = row.Name
		finding := strconv.FormatInt(int64(row.ID), 10)
		_, found := stnccollection.FindSlice(catsPost, finding)
		if found {
			catsData[key].SelectedID = row.ID
		}
	}
	if len(savekioskSliderError) == 0 {
		kiosk2Slider.Picture = filename
		_, filetype := stnc2upload.NewFileUpload().RealImage("public/upl/kiosk/" + filename)

		if filetype != "video/mp4" {
			kiosk2Slider.Type = 1
		} else {
			kiosk2Slider.Type = 2
		}
		saveData, saveErr := access.Kiosk2SliderApp.Update(&kiosk2Slider)
		if saveErr != nil {
			savekioskSliderError = saveErr
		} else {
			if filetype != "video/mp4" {
				if filenameForm != nil {
					src, err := imaging.Open("public/upl/kiosk/" + filename)
					if err != nil {
						log.Fatalf("failed to open image: %v", err)
					}

					// src = imaging.Resize(src, 1815, 2510, imaging.Center)
					// var srcBig image.Image
					var srcBig = imaging.Resize(src, 1815, 2510, imaging.Lanczos)

					errBig := imaging.Save(srcBig, "public/upl/kiosk/big/"+filename)
					if errBig != nil {
						log.Fatalf("failed to save image: %v", errBig)
					}

					var srcThumb = imaging.Resize(src, 150, 150, imaging.Lanczos)

					errth := imaging.Save(srcThumb, "public/upl/kiosk/thumb/"+filename)
					if errth != nil {
						log.Fatalf("failed to save image: %v", errth)
					}
				}
			}
		}

		var catPost = entity.CategoriesKioskJoin{}
		access.CategoriesKioskJoin.DeleteForKioskID(idN)
		for _, row := range catsPost {
			catID, _ := strconv.ParseUint(row, 10, 64)
			catPost.CategoryID = catID
			catPost.KioskID = saveData.ID
			saveCat, _ := access.CategoriesKioskJoin.Save(&catPost)
			catPost.ID = saveCat.ID + 1
		}
		stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/kiosk2Slider/edit/"+id)
		return
	}

	viewData := pongo2.Context{
		"title": "içerik düzenleme",
		"err":   savekioskSliderError,
		"csrf":  csrf.GetToken(c),
		"data":  kiosk2Slider,
	}
	c.HTML(
		http.StatusOK,
		viewPathkioskSlider2+"edit.html",
		viewData,
	)
}

//TODO: resim silmeyi unutma
//Delete data
func (access *Kiosk2Slider) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if kioskSliderID, err := strconv.ParseUint(c.Param("KioskSliderID"), 10, 64); err == nil {

		access.Kiosk2SliderApp.Delete(kioskSliderID)
		stncsession.SetFlashMessage("Başarı ile silindi", c)

		c.Redirect(http.StatusMovedPermanently, "/admin/kioskSlider")
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (access *Kiosk2Slider) Status(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if kioskSliderID, err := strconv.ParseUint(c.Param("KioskSliderID"), 10, 64); err == nil {
		status := stnccollection.StringToint(c.DefaultQuery("status", "false"))
		access.Kiosk2SliderApp.SetKioskSliderUpdate(kioskSliderID, status)
		stncsession.SetFlashMessage("Durum Değiştirildi", c)

		c.Redirect(http.StatusMovedPermanently, "/admin/kiosk2Slider")
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//Delete data
func (access *Kiosk2Slider) Upload(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	filenameForm, _ := c.FormFile("file")

	var uploadError string
	var filename string

	filename, uploadError = stnc2upload.NewFileUpload().UploadFile(filenameForm, c.PostForm("Resim2"))

	_, filetype := stnc2upload.NewFileUpload().RealImage("public/upl/kiosk/" + filename)

	if filetype == "video/mp4" {
		uploadError = "hata mp4 olamaz"
	}

	if uploadError == "" {
		uploadError = "Başarlı ile yuklendi"
	}
	c.JSON(http.StatusBadGateway, uploadError)
	// c.JSON(http.StatusOK, "ok")
}

//form kioskSlider model
func kiosk2SliderModel(c *gin.Context) (kioskSlider entity.Kiosk2Slider, idD uint64, idStr string) {
	id := c.PostForm("ID")
	title := c.PostForm("Title")

	idInt, _ := strconv.Atoi(id)
	var idN uint64
	idN = uint64(idInt)
	//	var kioskSlider = entity.kioskSlider{}
	kioskSlider.ID = idN
	kioskSlider.UserID = 1
	kioskSlider.Title = title

	kioskSlider.Status = stnccollection.StringToint(c.PostForm("Status"))

	return kioskSlider, idN, id
}
