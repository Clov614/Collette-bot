package bilibiliAnalysis

type VideoInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizId   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State     int `json:"state"`
		Duration  int `json:"duration"`
		MissionId int `json:"mission_id"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			ArgueMsg   string `json:"argue_msg"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		NoCache bool `json:"no_cache"`
		Pages   []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
		Staff []struct {
			Mid   int    `json:"mid"`
			Title string `json:"title"`
			Name  string `json:"name"`
			Face  string `json:"face"`
			Vip   struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path        string `json:"path"`
					Text        string `json:"text"`
					LabelTheme  string `json:"label_theme"`
					TextColor   string `json:"text_color"`
					BgStyle     int    `json:"bg_style"`
					BgColor     string `json:"bg_color"`
					BorderColor string `json:"border_color"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			} `json:"vip"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"official"`
			Follower   int `json:"follower"`
			LabelStyle int `json:"label_style"`
		} `json:"staff"`
		UserGarb struct {
			UrlImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
	} `json:"data"`
}

//type ShareData struct {
//	App   string `json:"app"`
//	Desc  string `json:"desc"`
//	Extra struct {
//		AppType int `json:"app_type"`
//		Appid   int `json:"appid"`
//		Uin     int `json:"uin"`
//	} `json:"extra"`
//}

type ShareDateAndroid struct {
	App    string `json:"app"`
	Config struct {
		AutoSize int    `json:"autoSize"`
		Ctime    int    `json:"ctime"`
		Forward  int    `json:"forward"`
		Height   int    `json:"height"`
		Token    string `json:"token"`
		Type     string `json:"type"`
		Width    int    `json:"width"`
	} `json:"config"`
	Desc  string `json:"desc"`
	Extra struct {
		AppType int `json:"app_type"`
		Appid   int `json:"appid"`
		Uin     int `json:"uin"`
	} `json:"extra"`
	Meta struct {
		Detail1 struct {
			AppType       int    `json:"appType"`
			Appid         string `json:"appid"`
			Desc          string `json:"desc"`
			GamePoints    string `json:"gamePoints"`
			GamePointsURL string `json:"gamePointsUrl"`
			Host          struct {
				Nick string `json:"nick"`
				Uin  int    `json:"uin"`
			} `json:"host"`
			Icon              string `json:"icon"`
			Preview           string `json:"preview"`
			Qqdocurl          string `json:"qqdocurl"`
			Scene             int    `json:"scene"`
			ShareTemplateData struct {
			} `json:"shareTemplateData"`
			ShareTemplateID string `json:"shareTemplateId"`
			ShowLittleTail  string `json:"showLittleTail"`
			Title           string `json:"title"`
			URL             string `json:"url"`
		} `json:"detail_1"`
	} `json:"meta"`
	NeedShareCallBack bool   `json:"needShareCallBack"`
	Prompt            string `json:"prompt"`
	Ver               string `json:"ver"`
	View              string `json:"view"`
}

type ShareDataIos struct {
	App    string `json:"app"`
	Config struct {
		Ctime   int    `json:"ctime"`
		Forward bool   `json:"forward"`
		Token   string `json:"token"`
		Type    string `json:"type"`
	} `json:"config"`
	Desc  string `json:"desc"`
	Extra struct {
		AppType int `json:"app_type"`
		Appid   int `json:"appid"`
		Uin     int `json:"uin"`
	} `json:"extra"`
	Meta struct {
		News struct {
			Action         string `json:"action"`
			AndroidPkgName string `json:"android_pkg_name"`
			AppType        int    `json:"app_type"`
			Appid          int    `json:"appid"`
			Ctime          int    `json:"ctime"`
			Desc           string `json:"desc"`
			JumpURL        string `json:"jumpUrl"`
			Preview        string `json:"preview"`
			SourceIcon     string `json:"source_icon"`
			SourceURL      string `json:"source_url"`
			Tag            string `json:"tag"`
			Title          string `json:"title"`
			Uin            int    `json:"uin"`
		} `json:"news"`
	} `json:"meta"`
	Prompt string `json:"prompt"`
	Ver    string `json:"ver"`
	View   string `json:"view"`
}
