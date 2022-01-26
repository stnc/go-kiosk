package controller

import (
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/domain/repository"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const viewPathIndex = "admin/index/"

//Index all list f
func Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	viewData := pongo2.Context{
		"title": "Posts",
		"csrf":  csrf.GetToken(c),
	}
	c.HTML(
		http.StatusOK,
		viewPathIndex+"index.html",
		viewData,
	)
}

//OptionsDefault all list f
func OptionsDefault(c *gin.Context) {
	// stncsession.IsLoggedInRedirect(c)

	//buraya bir oprion otılacak bunlar giriş yaptıktan sonra veri varmı yok mu bakacak

	db := repository.DB

	option1 := entity.Options{OptionName: "pageRenewTime", OptionValue: "50"}
	db.Debug().Create(&option1)

	pageRenewStatus := entity.Options{OptionName: "pageRenewStatus", OptionValue: "1"}
	db.Debug().Create(&pageRenewStatus)

	dolar := entity.Options{OptionName: "dolar", OptionValue: ""}
	db.Debug().Create(&dolar)

	euro := entity.Options{OptionName: "euro", OptionValue: ""}
	db.Debug().Create(&euro)

	altin := entity.Options{OptionName: "altin", OptionValue: ""}
	db.Debug().Create(&altin)

	ceyrek_altin := entity.Options{OptionName: "ceyrek_altin", OptionValue: ""}
	db.Debug().Create(&ceyrek_altin)

	weatherTodayIcon := entity.Options{OptionName: "weatherTodayIcon", OptionValue: ""}
	db.Debug().Create(&weatherTodayIcon)

	weatherTodayDescription := entity.Options{OptionName: "weatherTodayDescription", OptionValue: ""}
	db.Debug().Create(&weatherTodayDescription)

	weatherTodayDegree := entity.Options{OptionName: "weatherTodayDegree", OptionValue: ""}
	db.Debug().Create(&weatherTodayDegree)

	weatherTodayNight := entity.Options{OptionName: "weatherTodayNight", OptionValue: ""}
	db.Debug().Create(&weatherTodayNight)

	weatherTodayHumidity := entity.Options{OptionName: "weatherTodayHumidity", OptionValue: ""}
	db.Debug().Create(&weatherTodayHumidity)

	covid19Confirmed := entity.Options{OptionName: "covid19Confirmed", OptionValue: ""}
	db.Debug().Create(&covid19Confirmed)

	covid19Deaths := entity.Options{OptionName: "covid19Deaths", OptionValue: ""}
	db.Debug().Create(&covid19Deaths)

	covid19Recovered := entity.Options{OptionName: "covid19Recovered", OptionValue: ""}
	db.Debug().Create(&covid19Recovered)

	covid19Aktive := entity.Options{OptionName: "covid19Aktive", OptionValue: ""}
	db.Debug().Create(&covid19Aktive)

	//Teknopark2018
	user := entity.User{FirstName: "İnfo  ", LastName: "erciyes", Email: "info@erciyesteknopark.com", Password: "cb5e6834e30cf762b38387db44c936daac667559"}
	db.Debug().Create(&user)

	c.JSON(http.StatusOK, "yapıldı")
}
