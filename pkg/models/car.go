package models

type CarModel struct {
	Hsn          string `json:"hsn"`
	Tsn          string `json:"tsn"`
	Name         string `json:"name"`
	HaendlerName string `json:"haendler_hsn"`
}

type CarResponse struct {
	Name string `json:"name"`
	Tsn  string `json:"tsn"`
}

type HaendlerResponse struct {
	Name string `json:"name"`
	Hsn  string `json:"hsn"`
}
