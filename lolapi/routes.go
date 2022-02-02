package lolapi

type Route string

const (
	RouteAmericas Route = "https://americas.api.riotgames.com/riot"
	RouteAsia     Route = "https://asia.api.riotgames.com"
	RouteEsports  Route = "https://esports.api.riotgames.com"
	RouteEurope   Route = "https://europe.api.riotgames.com"
)
