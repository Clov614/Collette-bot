package bilibiliAnalysis

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"time"
)

var (
	shareDate ShareDate
	videoInfo VideoInfo
)

func BiliAnalysis(receive_msg string) (status bool, message string) {
	re, _ := regexp.Compile("\\[CQ:json,data={\"app\":\"com.tencent")
	if re.MatchString(receive_msg) {
		rawUrl, bvid := handleCQcode(receive_msg)
		getVideoinfo(bvid)
		message = mergeTOcqcode(rawUrl)
		return true, message
	}
	return
}

func BilirawUrlanalysis(receive_msg string) (status bool, message string) {

	re, _ := regexp.Compile("https://www.bilibili.com/video/([\\s\\S]*?)?")
	if re.MatchString(receive_msg) {
		bvids := re.FindStringSubmatch(receive_msg)
		bvid := bvids[1]
		rawUrl := "https://www.bilibili.com/video/" + bvid
		getVideoinfo(bvid)
		message = mergeTOcqcode(rawUrl)
		return true, message
	}
	return false, ""
}

// 处理[CQ: json date=...] 返回视频原链接和bvid
func handleCQcode(receive_msg string) (rawUrl string, bvid string) {
	re, _ := regexp.Compile("\\[CQ:json,data=([\\s\\S]*)]") // 注意: 使用非贪婪匹配防止内容减少
	dates := re.FindStringSubmatch(receive_msg)
	_ = json.Unmarshal([]byte(dates[1]), &shareDate)

	qqdocurl := shareDate.Meta.Detail1.Qqdocurl
	resp, err := grequests.Get(qqdocurl, nil)
	if err != nil {
		log.Error("getqqdocurl error: ", err)
	}
	reRawurl, _ := regexp.Compile("\"url\" content=\"([\\s\\S]*?)/\">")
	reBvid, _ := regexp.Compile("\"url\" content=\"https://www.bilibili.com/video/([\\s\\S]*?)/\">")
	rawUrls := reRawurl.FindStringSubmatch(resp.String())
	bvids := reBvid.FindStringSubmatch(resp.String())
	rawUrl = rawUrls[1]
	bvid = bvids[1]
	return rawUrl, bvid
}

func getVideoinfo(bvid string) {
	rUrl := "https://api.bilibili.com/x/web-interface/view"
	RO := grequests.RequestOptions{
		Params: map[string]string{
			"bvid": bvid,
		},
	}
	resp, err := grequests.Get(rUrl, &RO)
	if err != nil {
		log.Error("getVideoinfo error: ", err)
	}
	_ = json.Unmarshal([]byte(resp.String()), &videoInfo)
}

// 合并为待发送CQ码
func mergeTOcqcode(rawUrl string) (sendMsg string) {
	cqImg := fmt.Sprintf("[CQ:image,file=%s]", videoInfo.Data.Pic)
	tittle := videoInfo.Data.Title
	ctime := videoInfo.Data.Ctime
	tm := time.Unix(int64(ctime), 0)
	timeStr := fmt.Sprintf(tm.Format("2006-01-02 03:04:05 PM"))
	tag := videoInfo.Data.Tname
	view := strconv.Itoa(videoInfo.Data.Stat.View)
	danmaku := strconv.Itoa(videoInfo.Data.Stat.Danmaku)

	reply := strconv.Itoa(videoInfo.Data.Stat.Reply)
	favorite := strconv.Itoa(videoInfo.Data.Stat.Favorite)
	like := strconv.Itoa(videoInfo.Data.Stat.Like)
	coin := strconv.Itoa(videoInfo.Data.Stat.Coin)
	share := strconv.Itoa(videoInfo.Data.Stat.Share)
	bvid := videoInfo.Data.Bvid
	sendMsg = "标题: " + tittle + "\n" + cqImg + "\n" + "分区: " + tag + " 投稿时间: " + timeStr + "\n"
	sendMsg = sendMsg + "播放量: " + view + " 弹幕数: " + danmaku + " 评论数: " + reply + "\n"
	sendMsg = sendMsg + "点赞: " + like + " 投币: " + coin + " 收藏: " + favorite + " 分享: " + share + "\n"
	sendMsg = sendMsg + rawUrl + "\n" + "\t\t" + "\n" + "\t"
	sendMsg = sendMsg + bvid
	return sendMsg
}
