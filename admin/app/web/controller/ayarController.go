package controller

import (
	"net/http"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//Options constructor
type Options struct {
	OptionsApp services.OptionsAppInterface
}

const viewPathOptions = "admin/options/"

//InitOptions post controller constructor
func InitOptions(OptionsApp services.OptionsAppInterface) *Options {
	return &Options{
		OptionsApp: OptionsApp,
	}
}

//Index list
func (access *Options) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	pageRenewTime := access.OptionsApp.GetOption("pageRenewTime")
	pageRenewStatus := access.OptionsApp.GetOption("pageRenewStatus")
	dolar := access.OptionsApp.GetOption("dolar")
	euro := access.OptionsApp.GetOption("euro")
	altin := access.OptionsApp.GetOption("altin")
	ceyrek_altin := access.OptionsApp.GetOption("ceyrek_altin")
	weatherTodayIcon := access.OptionsApp.GetOption("weatherTodayIcon")
	weatherTodayDescription := access.OptionsApp.GetOption("weatherTodayDescription")
	weatherTodayDegree := access.OptionsApp.GetOption("weatherTodayDegree")
	weatherTodayNight := access.OptionsApp.GetOption("weatherTodayNight")
	weatherTodayHumidity := access.OptionsApp.GetOption("weatherTodayHumidity")
	covid19Confirmed := access.OptionsApp.GetOption("covid19Confirmed")
	covid19Deaths := access.OptionsApp.GetOption("covid19Deaths")
	covid19Recovered := access.OptionsApp.GetOption("covid19Recovered")
	covid19Aktive := access.OptionsApp.GetOption("covid19Aktive")

	viewData := pongo2.Context{
		"title":                   "Ayarlar",
		"csrf":                    csrf.GetToken(c),
		"pageRenewTime":           pageRenewTime,
		"pageRenewStatus":         pageRenewStatus,
		"dolar":                   dolar,
		"euro":                    euro,
		"altin":                   altin,
		"ceyrek_altin":            ceyrek_altin,
		"weatherTodayIcon":        weatherTodayIcon,
		"weatherTodayDescription": weatherTodayDescription,
		"weatherTodayDegree":      weatherTodayDegree,
		"weatherTodayNight":       weatherTodayNight,
		"weatherTodayHumidity":    weatherTodayHumidity,
		"covid19Confirmed":        covid19Confirmed,
		"covid19Deaths":           covid19Deaths,
		"covid19Recovered":        covid19Recovered,
		"covid19Aktive":           covid19Aktive,
	}

	c.HTML(
		http.StatusOK,
		viewPathOptions+"index.html",
		viewData,
	)
}

//Update list
func (access *Options) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	access.OptionsApp.SetOption("pageRenewTime", c.PostForm("pageRenewTime"))
	access.OptionsApp.SetOption("pageRenewStatus", c.PostForm("pageRenewStatus"))
	access.OptionsApp.SetOption("dolar", c.PostForm("dolar"))
	access.OptionsApp.SetOption("euro", c.PostForm("euro"))
	access.OptionsApp.SetOption("altin", c.PostForm("altin"))
	access.OptionsApp.SetOption("ceyrek_altin", c.PostForm("ceyrek_altin"))
	access.OptionsApp.SetOption("weatherTodayIcon", c.PostForm("weatherTodayIcon"))
	access.OptionsApp.SetOption("weatherTodayDescription", c.PostForm("weatherTodayDescription"))
	access.OptionsApp.SetOption("weatherTodayDegree", c.PostForm("weatherTodayDegree"))
	access.OptionsApp.SetOption("weatherTodayNight", c.PostForm("weatherTodayNight"))
	access.OptionsApp.SetOption("weatherTodayHumidity", c.PostForm("weatherTodayHumidity"))
	access.OptionsApp.SetOption("covid19Confirmed", c.PostForm("covid19Confirmed"))
	access.OptionsApp.SetOption("covid19Deaths", c.PostForm("covid19Deaths"))
	access.OptionsApp.SetOption("covid19Recovered", c.PostForm("covid19Recovered"))
	access.OptionsApp.SetOption("covid19Aktive", c.PostForm("covid19Aktive"))

	c.Redirect(http.StatusMovedPermanently, "/admin/options")
	return

}
