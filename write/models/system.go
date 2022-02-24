package models

type Theme struct {
	BackgroundImage string `json:"background_image"`
	GlobalCSS       string `json:"global_css"`
	GlobalJS        string `json:"global_js"`
}
