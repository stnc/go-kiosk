//  golang gin framework mvc and clean code project
//  Licensed under the Apache License 2.0
//  @author Selman TUNÇ <selmantunc@gmail.com>
//  @link: https://github.com/stnc/go-mvc-blog-clean-code
//  @license: Apache License 2.0
package main

import (
	"log"
	"net/http"
	"os"

	"stncCms/app/domain/repository"

	"stncCms/app/web.api/controller/middleware"
	"stncCms/app/web/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/stnc/pongo2gin"
	csrf "github.com/utrack/gin-csrf"
)

//https://github.com/stnc-go/gobyexample/blob/master/pongo2render/render.go
func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	/* //bu sunucuda çalışıyor
		    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	        if err != nil {
	            log.Fatal(err)
	        }
	        environmentPath := filepath.Join(dir, ".env")
	        err = godotenv.Load(environmentPath)
	        fatal(err)
	*/

}

func main() {

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// redisHost := os.Getenv("REDIS_HOST")
	// redisPort := os.Getenv("REDIS_PORT")
	// redisPassword := os.Getenv("REDIS_PASSWORD")
	debugMode := os.Getenv("MODE")
	gormAdvancedLogger := os.Getenv("GORM_ZAP_LOGGER")

	db := repository.DbConnect(dbdriver, user, password, port, host, dbname, debugMode, gormAdvancedLogger)
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()

	// redisService, err := auth.RedisDBInit(redisHost, redisPort, redisPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// upload := stncupload.NewFileUpload()

	kioskSlider := controller.InitKioskSlider(services.KioskSlider, services.User, services.Options)

	switch debugMode {
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)

	case "DEBUG":
		gin.SetMode(gin.DebugMode)

	case "TEST":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(gin.Recovery())

	//TODO: https://github.com/denisbakhtin/ginblog/blob/master/main.go burada memstore kullanımı var ona bakılablir

	store := cookie.NewStore([]byte("SpeedyGonzales"))

	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 7 * 86400}) //Also set Secure: true if using SSL, you should though

	r.Use(sessions.Sessions("myCRM", store))

	r.Use(middleware.CORSMiddleware()) //For CORS

	//TODO: csrf kontrolu nasıl olacak
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "SpeedyGonzales",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.HTMLRender = pongo2gin.TemplatePath("public/view")

	r.MaxMultipartMemory = 1 >> 20 // 8 MiB

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Static("/assets", "./public/static")

	r.StaticFS("/upload", http.Dir("./public/upl"))
	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("optionsDefault", controller.OptionsDefault)

	//kiosk başladı
	kioskGroup := r.Group("/")
	{
		kioskGroup.GET("/", kioskSlider.Index)
		kioskGroup.GET("/ajaxApi", kioskSlider.AjaxApi)
		kioskGroup.GET("bina/:kioskSliderID", kioskSlider.Ekran)
	}

	//Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8888" //localhost
	}
	log.Fatal(r.Run(":" + appPort))

}
