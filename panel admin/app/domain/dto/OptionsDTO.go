package dto

//Kurban dto
type OptionsDto struct {
	PageRenewTime           int    `json:"pageRenewTime"`
	PageRenewStatus         bool   `json:"pageRenewStatus"`
	Dolar                   string `json:"dolar"`
	Euro                    string `json:"euro"`
	Altin                   string `json:"altin"`
	CeyrekAltin             string `json:"ceyrek_altin"`
	WeatherTodayIcon        string `json:"weatherTodayIcon"`
	WeatherTodayDescription string `json:"weatherTodayDescription"`
	WeatherTodayDegree      string `json:"weatherTodayDegree"`
	WeatherTodayNight       string `json:"weatherTodayNight"`
	WeatherTodayHumidity    string `json:"weatherTodayHumidity"`
	Covid19Confirmed        string `json:"covid19Confirmed"`
	Covid19Deaths           string `json:"covid19Deaths"`
	Covid19Recovered        string `json:"covid19Recovered"`
	Covid19Aktive           string `json:"covid19Aktive"`
}
