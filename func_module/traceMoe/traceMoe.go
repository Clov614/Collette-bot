package traceMoe

import (
	"Collette_bot/BaseEvent"
	"Collette_bot/setting"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

type EchoCache struct {
	UserId    int
	TimeStamp int64
}

// 将二次判定缓存于内存
var (
	Echos       map[int]EchoCache = make(map[int]EchoCache, 10)
	resTracemoe ResTracemoe
)

// 消息二次判断机制
func echoDo(msgEvent BaseEvent.PluginsMsg) {
	Echos[msgEvent.UserId] = EchoCache{
		UserId:    msgEvent.UserId,
		TimeStamp: time.Now().Unix(),
	}
}

func getImgurl(msg string) (imgUrl string) {
	reTypeimg, _ := regexp.Compile("CQ:image")
	reImgurl, _ := regexp.Compile("url=([\\s\\S]*?)]")
	if reTypeimg.MatchString(msg) {
		imgUrlslice := reImgurl.FindStringSubmatch(msg)
		imgUrl = imgUrlslice[1]
	} else {
		imgUrl = "NULL"
	}
	return imgUrl
}

func sendtoTracemoe(imgUrl string) (sendMsg string) {
	RO := grequests.RequestOptions{Headers: map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36",
		"Content-Type": "application/json",
	}}

	sendUrl := fmt.Sprintf("https://api.trace.moe/search?anilistInfo&url=%s", imgUrl)
	resp, err := grequests.Get(sendUrl, &RO)
	if err != nil {
		log.Error(err)
	}

	_ = json.Unmarshal(resp.Bytes(), &resTracemoe)
	result := resTracemoe.Result[0]
	anilist := result.Anilist
	title := anilist.Title.Native    // 标题
	romaji := anilist.Title.Romaji   // 罗马音
	english := anilist.Title.English // 英文
	episode := result.Episode        // 集数
	start := (result.From) / 60      // 开始时刻
	end := result.To / 60            // 结束时刻
	similarity := result.Similarity  // 相似度
	//video := result.Video            // 视频
	image := result.Image // 图片
	//video_CQ := fmt.Sprintf("[CQ:video,file=%s]", video)
	sendMsg = fmt.Sprintf("搜索结果:\n番名:%s\n罗马音:%s\nenglish:%s\n[CQ:image,file=%s]\n位置:第%d集\nStart:/%.2f/\nEnd:/%.2f/\n图片时刻:/%.2f/\n相似度:%.2f\npower by trace_moe！！", title, romaji, english, image, episode, start, end, start, similarity)
	return sendMsg
}

func TraceMoe(msgEvent BaseEvent.PluginsMsg, dataCheck *BaseEvent.PluginsData) {
	receiveMsg := msgEvent.Message

	reImgurl, _ := regexp.Compile("url=([\\s\\S]*?)]")
	echo, ok := Echos[msgEvent.UserId]
	if ok {
		timeNow := time.Now().Unix()
		if (timeNow-echo.TimeStamp) <= 60 && reImgurl.MatchString(receiveMsg) {
			delete(Echos, msgEvent.UserId)
			imgUrl := getImgurl(receiveMsg)
			sendMsg := sendtoTracemoe(imgUrl)
			dataCheck.Status = true
			dataCheck.SendMsg = sendMsg
		} else {
			delete(Echos, msgEvent.UserId)
		}
		delete(Echos, msgEvent.UserId)
	}

	// 判断是否有at机器人的变量
	var ATstatus bool
	re, _ := regexp.Compile("\\[CQ:at,qq=" + setting.Data.SelfQQ + "]")
	// 判断是否有at机器人
	for _, V := range setting.Data.Nickname {
		reNick, _ := regexp.Compile(V)
		if reNick.MatchString(receiveMsg) {
			ATstatus = true
		}
	}
	if re.MatchString(receiveMsg) {
		ATstatus = true
	}
	if ATstatus {
		targets := []string{"以图搜番", "搜番", "trace_moe", "traceMoe"}
		for _, v := range targets {
			reServer, _ := regexp.Compile(v)
			if reServer.MatchString(receiveMsg) {
				echoDo(msgEvent)
				dataCheck.Status = true
				dataCheck.SendMsg = "请发出需要识别的图片(1分钟内)"
			}
		}
	}
}
