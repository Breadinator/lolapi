package lolapi

import "strings"

type Region string

const (
	RegionInvalid           Region = "invalid"
	RegionBrasil            Region = "br1"
	RegionEuropeNorthEast   Region = "eun1"
	RegionEuropeWest        Region = "euw1"
	RegionJapan             Region = "jp1"
	RegionKorea             Region = "kr"
	RegionLatinAmericaNorth Region = "la1"
	RegionLatinAmericaSouth Region = "la2"
	RegionNorthAmerica      Region = "na1"
	RegionOceania           Region = "oc1"
	RegionTurkey            Region = "tr1"
	RegionRussia            Region = "ru"
	RegionPBE               Region = "pbe1"
)

func RouteFromRegion(region Region) Route {
	switch region {
	case RegionBrasil, RegionLatinAmericaNorth, RegionLatinAmericaSouth, RegionNorthAmerica, RegionOceania:
		return RouteAmericas
	case RegionEuropeNorthEast, RegionEuropeWest, RegionTurkey, RegionRussia:
		return RouteEurope
	case RegionJapan, RegionKorea:
		return RouteAsia
	default:
		return RouteAmericas
	}
}

func RegionFromString(region string) Region {
	switch strings.ToLower(region) {
	case "br1", "br", "brazil", "brasil":
		return RegionBrasil
	case "eune", "eun", "eun1", "europe north east", "europe nordic and east", "europe nordic & east":
		return RegionEuropeNorthEast
	case "euw", "euw1", "europe west":
		return RegionEuropeWest
	case "jp", "jp1", "japan":
		return RegionJapan
	case "kr", "korean":
		return RegionKorea
	case "la1", "latin america north", "latam north":
		return RegionLatinAmericaNorth
	case "oc1", "oceania":
		return RegionOceania
	case "tr", "tr1", "turkey":
		return RegionTurkey
	case "ru", "russia":
		return RegionRussia
	case "pbe", "pbe1":
		return RegionPBE
	default:
		return RegionInvalid
	}
}
