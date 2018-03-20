package main

import (
	"github.com/valyala/fasthttp"
	"github.com/tidwall/gjson"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xormplus/xorm"
	"fmt"
)

type Video struct {
	Title     string
	Comment string
	Play    int64
	CreatedAt      int64
	Length   string
	Video_Review   int64
	Favorites int64
}

func main() {
	var users []Video
	engine, _ := xorm.NewEngine("sqlite3", "./test.db")
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI("https://space.bilibili.com/ajax/member/getSubmitVideos?mid=2301165&pagesize=30&tid=0&page=1&keyword=%E6%AC%85%E4%BC%9A%E4%B8%8D%E4%BC%9A%E5%86%99&order=pubdate")
	fasthttp.Do(req, resp)
	for _, x := range gjson.GetBytes(resp.Body(), "data.vlist").Array() {
		user := Video{Title: x.Get("title").String(), Comment: x.Get("comment").String(), Play: x.Get("play").Int(), CreatedAt: x.Get("created").Int(), Length: x.Get("length").String(), Video_Review: x.Get("video_review").Int(), Favorites: x.Get("favorites").Int()}
		users = append(users, user)
	}
	engine.Sync2(new(Video))
	affected, err := engine.Insert(&users)
	fmt.Println(affected,err)
}
