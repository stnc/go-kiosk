package controller

import (
	"encoding/json"
	"net/http"
	"stncCms/app/domain/entity"
	stnccollection "stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncdatetime"

	"stncCms/app/services"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

//KioskSlider constructor
type KioskSlider struct {
	KioskSliderApp services.KioskSliderAppInterface
	userApp        services.UserAppInterface
	optionsApp     services.OptionsAppInterface
}

const viewPathKioskSlider = "kioskSlider/"

//InitKioskSlider KioskSlider controller constructor
func InitKioskSlider(pApp services.KioskSliderAppInterface, uApp services.UserAppInterface, oApp services.OptionsAppInterface) *KioskSlider {
	return &KioskSlider{
		KioskSliderApp: pApp,
		userApp:        uApp,
		optionsApp:     oApp,
	}
}

//Index list
func (access *KioskSlider) Index(c *gin.Context) {
	// allKioskSlider, err := access.KioskSliderApp.GetAllKioskSlider()

	KioskSliders, _ := access.KioskSliderApp.GetAll()

	viewData := pongo2.Context{
		"title":                   "Slider",
		"allData":                 KioskSliders,
		"pageRenewTime":           access.optionsApp.GetOption("pageRenewTime"),
		"pageRenewStatus":         access.optionsApp.GetOption("pageRenewStatus"),
		"dolar":                   access.optionsApp.GetOption("dolar"),
		"euro":                    access.optionsApp.GetOption("euro"),
		"altin":                   access.optionsApp.GetOption("altin"),
		"ceyrek_altin":            access.optionsApp.GetOption("ceyrek_altin"),
		"weatherTodayIcon":        access.optionsApp.GetOption("weatherTodayIcon"),
		"weatherTodayDescription": access.optionsApp.GetOption("weatherTodayDescription"),
		"weatherTodayDegree":      access.optionsApp.GetOption("weatherTodayDegree"),
		"weatherTodayNight":       access.optionsApp.GetOption("weatherTodayNight"),
		"weatherTodayHumidity":    access.optionsApp.GetOption("weatherTodayHumidity"),
		"Covid19Confirmed":        access.optionsApp.GetOption("Covid19Confirmed"),
		"Covid19Deaths":           access.optionsApp.GetOption("Covid19Deaths"),
		"Covid19Recovered":        access.optionsApp.GetOption("Covid19Recovered"),
		"Covid19Aktive":           access.optionsApp.GetOption("Covid19Aktive"),
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		viewPathKioskSlider+"index.html",
		// Pass the data that the page uses
		viewData,
	)
}

//Edit edit data
func (access *KioskSlider) Ekran(c *gin.Context) {
	postID := stnccollection.StringToint(c.Param("kioskSliderID"))
	if posts, err := access.KioskSliderApp.GetAllCatID(postID); err == nil {
		viewData := pongo2.Context{
			"allData":                 posts,
			"EkranId":                 postID,
			"pageRenewTime":           access.optionsApp.GetOption("pageRenewTime"),
			"pageRenewStatus":         access.optionsApp.GetOption("pageRenewStatus"),
			"dolar":                   access.optionsApp.GetOption("dolar"),
			"euro":                    access.optionsApp.GetOption("euro"),
			"altin":                   access.optionsApp.GetOption("altin"),
			"ceyrek_altin":            access.optionsApp.GetOption("ceyrek_altin"),
			"weatherTodayIcon":        access.optionsApp.GetOption("weatherTodayIcon"),
			"weatherTodayDescription": access.optionsApp.GetOption("weatherTodayDescription"),
			"weatherTodayDegree":      access.optionsApp.GetOption("weatherTodayDegree"),
			"weatherTodayNight":       access.optionsApp.GetOption("weatherTodayNight"),
			"weatherTodayHumidity":    access.optionsApp.GetOption("weatherTodayHumidity"),
			"Covid19Confirmed":        access.optionsApp.GetOption("Covid19Confirmed"),
			"Covid19Deaths":           access.optionsApp.GetOption("Covid19Deaths"),
			"Covid19Recovered":        access.optionsApp.GetOption("Covid19Recovered"),
			"Covid19Aktive":           access.optionsApp.GetOption("Covid19Aktive"),
		}
		c.HTML(
			http.StatusOK,
			viewPathKioskSlider+"index.html",
			viewData,
		)

	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}

}

//Index list
func (access *KioskSlider) AjaxApi(c *gin.Context) {
	// allKioskSlider, err := access.KioskSliderApp.GetAllKioskSlider()

	var tarih stncdatetime.Inow

	KioskSliders, _ := access.KioskSliderApp.GetAjaxData()
	var connectiondata = entity.KioskConnection{}
	connectiondata.Title = "Bağlandı sağlandı"
	data, _ := json.Marshal(KioskSliders)
	connectiondata.Content = string(data)
	structureNo := stnccollection.StringToint(c.DefaultQuery("bina", "0"))
	connectiondata.StructureNo = structureNo
	access.KioskSliderApp.Save(&connectiondata)

	viewData := pongo2.Context{
		"jsonData": KioskSliders,
		"tarih":    tarih,
	}
	c.JSON(http.StatusOK, viewData)
	return

}
