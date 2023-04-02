package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	. "obcsapi-go/dao"
	"obcsapi-go/tools"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gin-gonic/gin"
)

// Token1
func ObTodayHandler(c *gin.Context) {
	switch c.Request.Method {
	case "OPTIONS":
		c.Status(200)
	case "GET":
		// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		data, _ := json.Marshal(GetTodayDailyList())
		c.String(200, string(data))
	case "POST":
		decoder := json.NewDecoder(c.Request.Body)
		var memosData MemosData
		err := decoder.Decode(&memosData)
		if err != nil {
			log.Println(err)
		}
		err = DailyTextAppendMemos(memosData.Content)
		if err != nil {
			c.Error(err)
			c.Status(500)
			return
		}
		c.String(200, "Success")
	default:
		c.Status(404)
	}
}

// Token1
func ObPostTodayAllHandler(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var memosData MemosData
	err := decoder.Decode(&memosData)
	if err != nil {
		c.Error(err)
	} else {
		MdTextStore(GetDailyFileKey(), memosData.Content)
	}
	c.String(200, "Success")
}

// Tokne1
func ObGet3DaysHandler(c *gin.Context) {
	three_list := Get3DaysDailyList()
	data, err := json.Marshal(three_list)
	if err != nil {
		c.Status(500)
		c.Error(err)
		return
	}
	c.String(200, string(data))
}

// Token2 静读天下使用的 API
func MoodReaderHandler(c *gin.Context) {
	right_token2, _ := tools.GetToken("token2")
	if c.Request.Header.Get("Authorization") != "Token "+right_token2.TokenString {
		c.Status(401)
		return
	}
	decoder := json.NewDecoder(c.Request.Body)
	var moodReader MoodReader
	err := decoder.Decode(&moodReader)
	if err != nil {
		c.Error(err)
		c.Status(500)
		return
	}
	fmt.Println(moodReader.Highlights[0])
	title := moodReader.Highlights[0].Title
	text := moodReader.Highlights[0].Text
	author := moodReader.Highlights[0].Author
	note := moodReader.Highlights[0].Note
	file_key := fmt.Sprintf("支持类文件/MoonReader/%s.md", tools.ReplaceUnAllowedChars(title))
	append_text := fmt.Sprintf("文: %s\n批: %s\n于: %s\n\n---\n", text, note, tools.TimeFmt("2006-01-02 15:04"))
	exist, _ := CheckObject(file_key)
	if exist {
		err = TextAppend(file_key, append_text)
	} else {
		yaml := fmt.Sprintf("---\ntitle: %s\nauthor: %s\n---\n书名: %s\n作者: %s\n简介: \n评价: \n\n---\n", title, author, title, author)
		err = TextAppend(file_key, yaml+append_text)
	}
	if err != nil {
		c.Error(err)
		c.Status(500)
		return
	}
	c.String(200, "Success")
}

// 安卓软件 fv 悬浮球使用的 API 用于自定义任务的 图片、文字
func fvHandler(c *gin.Context) {
	if c.GetHeader("Content-Type") == "text/plain" {
		content, _ := ioutil.ReadAll(c.Request.Body)
		DailyTextAppendMemos(string(content))
		c.String(200, "Success")
		return
	} else if c.GetHeader("Content-Type") == "application/octet-stream" {
		content, _ := ioutil.ReadAll(c.Request.Body)
		file_key := fmt.Sprintf("%s%s/%s.jpg", tools.ConfigGetString("ob_daily_attachment_dir"), tools.TimeFmt("200601"), tools.TimeFmt("20060102150405"))
		ObjectStore(file_key, content)
		DailyTextAppendMemos(fmt.Sprintf("![](%s)", file_key))
		c.String(200, "Success")
		return
	}
	c.String(404, "Error")
}

// SimpRead WebHook Used
type SimpReadWebHookStruct struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"desc"`
	Tags        string `json:"tags"`
	Content     string `json:"content"`
	Note        string `json:"note"`
}

// SimpRead WebHook Used Token2
func SRWebHook(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var simpReadJson SimpReadWebHookStruct
	err := decoder.Decode(&simpReadJson)
	if err != nil {
		c.Error(err)
		c.Status(500)
		return
	}
	serverTime := tools.TimeFmt("200601021504")
	yaml := fmt.Sprintf("---\ntitle: %s\nsctime: %s\n---\n", simpReadJson.Title, serverTime)
	file_str := fmt.Sprintf("%s[[简悦WebHook生成]]\n生成时间: %s\n原文: %s\n标题: %s\n描述: %s\n标签: %s\n内容: \n%s", yaml, serverTime, simpReadJson.Url, simpReadJson.Title, simpReadJson.Description, simpReadJson.Tags, simpReadJson.Content)
	file_key := fmt.Sprintf("支持类文件/SimpRead/%s %s.md", tools.ReplaceUnAllowedChars(simpReadJson.Title), serverTime)
	MdTextStore(file_key, file_str)
}

// 通用 API 接口 使用 Token2 验证
func GeneralHeader(c *gin.Context) {
	switch c.Request.Method {
	case "OPTIONS":
		c.Status(200)
	case "POST":
		decoder := json.NewDecoder(c.Request.Body)
		var memosData MemosData
		err := decoder.Decode(&memosData)
		if err != nil {
			c.Error(err)
			c.Status(500)
			return
		}
		err = DailyTextAppendMemos(memosData.Content)
		if err != nil {
			c.Error(err)
			c.Status(500)
			return
		}
		c.String(200, "Success")
	default:
		c.Status(404)
	}
}

// Token2
func Url2MdHandler(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var urlStruct UrlStruct
	err := decoder.Decode(&urlStruct)
	if err != nil {
		c.Error(err)
		c.Status(500)
		return
	}
	text, err := tools.Downloader(urlStruct.Url)
	if err != nil {
		c.Error(err)
		c.Status(500)
		return
	}
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(text))
	if err != nil {
		c.Error(err)
		c.Status(500)
		return
	}
	if err != nil {
		c.Error(err)
		c.Status(500)
		return
	}
	serverTime := tools.TimeFmt("200601021504")
	yaml := fmt.Sprintf("---\nurl: %s\nsctime: %s\n---\n", urlStruct.Url, serverTime)
	file_key := fmt.Sprintf("支持类文件/HtmlPages/%s.md", serverTime)
	MdTextStore(file_key, yaml+markdown)
	c.Status(200)
}