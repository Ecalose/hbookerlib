package hbookermodel

type Login struct {
	Tip
	Data LoginData `json:"data"`
}

type LoginData struct {
	LoginToken string     `json:"login_token"`
	UserCode   string     `json:"user_code"`
	ReaderInfo ReaderInfo `json:"reader_info"`
	PropInfo   PropInfo   `json:"prop_info"`
	IsSetYoung string     `json:"is_set_young"`
}

type Register struct {
	Code string      `json:"code"`
	Tip  interface{} `json:"tip"`
	Data LoginData   `json:"data"`
}
