package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncdatetime"
	"stncCms/app/domain/helpers/stncupload"
	"strings"

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
type KioskSlider struct {
	KioskSliderApp      services.KioskSliderAppInterface
	MediaApp            services.MediaAppInterface
	userApp             services.UserAppInterface
	categoriesKiosk     services.CategoriesKioskAppInterface
	CategoriesKioskJoin services.CategoriesKioskJoinAppInterface
}

const viewPathkioskSlider = "admin/kioskSlider/"

//InitkioskSlider kioskSlider controller constructor
func InitkioskSlider(KiApp services.KioskSliderAppInterface, uApp services.UserAppInterface, mediaApp services.MediaAppInterface,
	catKioskApp services.CategoriesKioskAppInterface, catKioskJoinApp services.CategoriesKioskJoinAppInterface) *KioskSlider {
	return &KioskSlider{
		KioskSliderApp:      KiApp,
		userApp:             uApp,
		MediaApp:            mediaApp,
		categoriesKiosk:     catKioskApp,
		CategoriesKioskJoin: catKioskJoinApp,
	}
}
func (access *KioskSlider) UploadConfig() map[string]string {
	returnData := make(map[string]string)

	maxFiles := 5
	uploadSize := 5
	uploadFile := "kiosk"
	uploadPath := "public/upl/" + uploadFile + "/"
	uploadFsPath := "upload/" + uploadFile + "/"

	returnData["uploadPath"] = uploadPath
	returnData["uploadSymbol"] = uploadFsPath
	returnData["maxFiles"] = stnccollection.IntToString(maxFiles)

	returnData["bigImageWidth"] = stnccollection.IntToString(1815)
	returnData["bigImageHeight"] = stnccollection.IntToString(2510)

	returnData["thumbImageWidth"] = stnccollection.IntToString(150)
	returnData["thumbImageHeight"] = stnccollection.IntToString(150)

	returnData["uploadSize"] = stnccollection.IntToString(uploadSize)
	returnData["fileType"] = "video/mp4,image/jpeg,image/jpg,image/gif,image/png,video/webm" //application/pdf

	return returnData
}

//Index list
func (access *KioskSlider) Index(c *gin.Context) {
	// allkioskSlider, err := access.kioskSliderApp.GetAllkioskSlider()

	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	var tarih stncdatetime.Inow
	var total int64
	access.KioskSliderApp.Count(&total)
	kioskSlidersPerPage := 10
	paginator := pagination.NewPaginator(c.Request, kioskSlidersPerPage, total)
	offset := paginator.Offset()

	kioskSliders, _ := access.KioskSliderApp.GetAllP(kioskSlidersPerPage, offset)

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
		viewPathkioskSlider+"index.html",
		// Pass the data that the page uses
		viewData,
	)
}

//Create all list f
func (access *KioskSlider) Create(c *gin.Context) {
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
		viewPathkioskSlider+"create.html",
		// Pass the data that the page uses
		viewData,
	)
}

//Store save method
func (access *KioskSlider) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var kioskSlider, _, _ = kioskSliderModel(c)
	var savekioskSliderError = make(map[string]string)

	savekioskSliderError = kioskSlider.Validate()

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

		saveData, saveErr := access.KioskSliderApp.Save(&kioskSlider)

		if saveErr != nil {
			savekioskSliderError = saveErr
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
		c.Redirect(http.StatusMovedPermanently, "/admin/kioskSlider/edit/"+lastID)
		return
	}

	viewData := pongo2.Context{
		"title":    "içerik ekleme",
		"catsPost": catsPost,
		"catsData": catsData,
		"csrf":     csrf.GetToken(c),
		"err":      savekioskSliderError,
		"data":     kioskSlider,
	}
	c.HTML(
		http.StatusOK,
		viewPathkioskSlider+"create.html",
		viewData,
	)

}

//Edit edit data
func (access *KioskSlider) Edit(c *gin.Context) {
	//strconv.Atoi(c.Param("id"))
	//kioskSliderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	if kioskSliderID, err := strconv.ParseUint(c.Param("KioskSliderID"), 10, 64); err == nil {

		if kioskSliders, err := access.KioskSliderApp.GetByID(kioskSliderID); err == nil {
			var catsPost []string
			kioskSliderIDint := stnccollection.StringToint(c.Param("KioskSliderID"))
			catsPostData, _ := access.CategoriesKioskJoin.GetAllforKioskID(kioskSliderID)
			mediaData, _ := access.MediaApp.GetAll(1, kioskSliderIDint)

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

			accData := access.UploadConfig()

			viewData := pongo2.Context{
				"title":      "içerik düzenleme",
				"catsPost":   catsPost,
				"catsData":   catsData,
				"data":       kioskSliders,
				"medias":     mediaData,
				"csrf":       csrf.GetToken(c),
				"flashMsg":   flashMsg,
				"fileConfig": accData,
			}
			c.HTML(
				http.StatusOK,
				viewPathkioskSlider+"edit.html",
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
func (access *KioskSlider) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	var kioskSlider, idN, id = kioskSliderModel(c)
	var savekioskSliderError = make(map[string]string)

	savekioskSliderError = kioskSlider.Validate()

	catsPost := c.PostFormArray("cats")
	mediaData, _ := access.MediaApp.GetAll(1, int(idN))

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

		saveData, saveErr := access.KioskSliderApp.Update(&kioskSlider)
		if saveErr != nil {
			savekioskSliderError = saveErr
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
		c.Redirect(http.StatusMovedPermanently, "/admin/kioskSlider/edit/"+id)
		return
	}

	viewData := pongo2.Context{
		"title":  "içerik düzenleme",
		"err":    savekioskSliderError,
		"csrf":   csrf.GetToken(c),
		"medias": mediaData,
		"data":   kioskSlider,
	}
	c.HTML(
		http.StatusOK,
		viewPathkioskSlider+"edit.html",
		viewData,
	)
}

//TODO: resim silmeyi unutma
//Delete data
func (access *KioskSlider) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if kioskSliderID, err := strconv.ParseUint(c.Param("KioskSliderID"), 10, 64); err == nil {

		access.KioskSliderApp.Delete(kioskSliderID)
		stncsession.SetFlashMessage("Başarı ile silindi", c)

		c.Redirect(http.StatusMovedPermanently, "/admin/kioskSlider")
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//Delete data
func (access *KioskSlider) MediaDelete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	accData := access.UploadConfig()

	uploadPath := accData["uploadPath"]

	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {

		if mediaData, err := access.MediaApp.GetByID(ID); err == nil {
			mediaName := mediaData.MediaName
			stncupload.FileDelete(uploadPath, mediaName)
			stncupload.FileDelete(uploadPath+"big/", mediaName)
			stncupload.FileDelete(uploadPath+"thumb/", mediaName)
			fmt.Println(mediaData)
		}

		access.MediaApp.Delete(ID)

		viewData := pongo2.Context{
			"status": "ok",
			"msg":    "Kayıt Başarı ile Silindi",
		}
		fmt.Println("girer")
		c.JSON(http.StatusOK, viewData)

		// c.JSON(http.StatusBadGateway, uploadError)

		return

	}
}

func (access *KioskSlider) Status(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if kioskSliderID, err := strconv.ParseUint(c.Param("KioskSliderID"), 10, 64); err == nil {
		status := stnccollection.StringToint(c.DefaultQuery("status", "false"))
		access.KioskSliderApp.SetKioskSliderUpdate(kioskSliderID, status)
		stncsession.SetFlashMessage("Durum Değiştirildi", c)

		c.Redirect(http.StatusMovedPermanently, "/admin/kioskSlider")
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//Delete data
func (access *KioskSlider) Upload(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	kioskSliderID, _ := strconv.ParseUint(c.Param("KioskSliderID"), 10, 64)
	kioskSliderIDint := stnccollection.StringToint(c.Param("KioskSliderID"))
	filenameForm, _ := c.FormFile("file")

	var uploadError string
	var filename string

	accData := access.UploadConfig()

	uploadPath := accData["uploadPath"]

	maxFiles := stnccollection.StringToint(accData["maxFiles"])
	uploadSize := stnccollection.StringToint(accData["uploadSize"])
	fileType := accData["fileType"]
	fileTypes := strings.Split(fileType, ",")

	_, err := os.Stat(uploadPath)
	if os.IsNotExist(err) {
		uploadError = uploadPath + " klasörü yok "
		//Create a folder/directory at a full qualified path
		err = os.Mkdir(uploadPath, 0755)
		if err != nil {
			// log.Fatal(err)
			log.Default()
		}

		err = os.Mkdir(uploadPath+"big/", 0755)
		if err != nil {
			// log.Fatal(err)
			log.Default()
		}

		err = os.Mkdir(uploadPath+"thumb/", 0755)
		if err != nil {
			// log.Fatal(err)
			log.Default()
		}
	}

	var total int

	access.MediaApp.Count(1, kioskSliderIDint, &total)

	if total >= maxFiles {
		uploadError = "Yükleme sayısını aştınız maksimum " + stnccollection.IntToString(maxFiles) + " dosya yükleyebilirsiniz"
		c.JSON(http.StatusBadGateway, uploadError)
		return
	}

	upl := stncupload.FileUpload{
		UploadPath: uploadPath,
		UploadSize: uploadSize,
		MaxFiles:   maxFiles,
		// Types:      []string{"video/mp4", "image/jpeg", "image/jpg", "image/gif", "image/png", "video/webm", "application/pdf"},
		// Types: []string{},
		Types: fileTypes,
	}

	var up stncupload.UploadFileInterface = upl

	filename, uploadError = up.UploadFile(filenameForm)

	// countTypes := len(upl.Types)

	var typeControlError bool = true

	var filetype string

	typeControlError, filetype = up.RealFileType(filename)

	switch filetype {
	case "image/jpeg", "image/jpg":
		access.imageResize(filename, uploadPath)
		stncupload.FileDelete(uploadPath, filename)
	case "image/gif":
		access.imageResize(filename, uploadPath)
		stncupload.FileDelete(uploadPath, filename)
	case "image/png":
		access.imageResize(filename, uploadPath)
		stncupload.FileDelete(uploadPath, filename)
		// default:
		// 	returnData = false
	}

	var mediaData = entity.Media{}
	mediaData.MediaName = filename
	mediaData.ModulID = 1
	mediaData.ContentID = kioskSliderID
	mediaData.UserID = 1 //TODO: user id git
	mediaData.MimeType = filetype

	saveData, saveErr := access.MediaApp.Save(&mediaData)

	if saveErr != nil {
		uploadError = "veritabanı hatası"
	}

	var returnID int
	returnID = saveData.ID
	if typeControlError == false {
		// stringSlice := strings.Split(filetype, "/") stringSlice[1]
		access.MediaApp.Delete(uint64(returnID))
		stncupload.FileDelete(uploadPath, filename)
		uploadError = "HATA: Yükleyeceğiniz dosya tipi uygun değildir, bunlar biri olmalıdır mp4, webm, jpeg, jpg, gif, png "
		c.JSON(http.StatusBadGateway, uploadError)
		return
	}

	// https://play.golang.org/p/UKZbcuJUPP

	if uploadError == "" {
		c.JSON(http.StatusOK, "Başarı ile yuklendi")
	} else {
		c.JSON(http.StatusBadGateway, uploadError)
	}
	fmt.Println(uploadError)
	return
}

//form kioskSlider model
func kioskSliderModel(c *gin.Context) (kioskSlider entity.KioskSlider, idD uint64, idStr string) {
	id := c.PostForm("ID")
	title := c.PostForm("Title")
	excerpt := c.PostForm("Excerpt")

	menuOrder := stnccollection.StringToint(c.PostForm("MenuOrder"))
	idInt, _ := strconv.Atoi(id)
	var idN uint64
	idN = uint64(idInt)
	//	var kioskSlider = entity.kioskSlider{}
	kioskSlider.ID = idN
	kioskSlider.UserID = 1
	kioskSlider.Title = title
	kioskSlider.Timer = stnccollection.StringToint(c.PostForm("Timer"))

	kioskSlider.Status = stnccollection.StringToint(c.PostForm("Status"))
	kioskSlider.MenuOrder = menuOrder

	kioskSlider.Excerpt = excerpt
	return kioskSlider, idN, id
}

func (access *KioskSlider) imageResize(filename string, path string) {

	fmt.Println("girer")
	src, err := imaging.Open(path + filename)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// src = imaging.Resize(src, 1815, 2510, imaging.Center)
	// var srcBig image.Image

	accData := access.UploadConfig()

	widthBig := stnccollection.StringToint(accData["bigImageWidth"])
	heightHeight := stnccollection.StringToint(accData["bigImageHeight"])

	weightThumb := stnccollection.StringToint(accData["thumbImageWidth"])
	heightThumb := stnccollection.StringToint(accData["thumbImageHeight"])

	var srcBig = imaging.Resize(src, widthBig, heightHeight, imaging.Lanczos)

	errBig := imaging.Save(srcBig, path+"big/"+filename)
	if errBig != nil {
		log.Fatalf("failed to save image: %v", errBig)
	}

	var srcThumb = imaging.Resize(src, weightThumb, heightThumb, imaging.Lanczos)

	errth := imaging.Save(srcThumb, path+"thumb/"+filename)
	if errth != nil {
		log.Fatalf("failed to save image: %v", errth)
	}

}
