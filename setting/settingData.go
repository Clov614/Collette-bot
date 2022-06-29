package setting

type SettingData struct {
	Nickname []string `json:"nickname"`
	SelfQQ   string   `json:"selfQQ"`
	Admin    []int    `json:"admin"`
}
