http://themes.semicolonweb.com/html/canvas/intro.html

hava durumu 
https://p.w3layouts.com/demos/june-2016/09-06-2016/new_york_weather_widget/web/

https://p.w3layouts.com/demos/june-2016/01-06-2016/sunny_weather_widget/web/


üst 
https://inspirothemes.com/polo/home-lawyer.html

güzel hava durumu örnekleri 
https://freefrontend.com/css-weather-widgets/


https://codepen.io/teerasak-vichadee/pen/xxKLQKZ


http://css-tricks.github.io/AnythingSlider/video.html#&panel1-4&panel2-2


http://flexslider.woothemes.com/video.html

https://codepen.io/ibrahima92/pen/ExxZVBN


https://codepen.io/BJack/pen/sBefL




https://codepen.io/derekjp/pen/yyVWOq


https://codepen.io/tomlutzenberger/pen/mPNoxj


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.collectapi.com/weather/getWeather?data.lang=tr&data.city=kayseri"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "apikey 47C2njeoSOx0ev4cg9BlbX:29uKQ2dNmReFgjIMXaSQDg")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}