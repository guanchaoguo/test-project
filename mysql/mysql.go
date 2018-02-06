package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"encoding/json"
)

var engine *xorm.Engine
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(192.168.31.231:3306)/AppDistribution?charset=utf8")
	if err != nil {
		fmt.Println("Connect to mysql error", err)
		return
	}
}

type App_info struct {
	Id          int32  `json:"id" xorm:"autoincr pk id"`
	User_id     int32  `json:"-"`
	Logo        string `json:"logo"`
	Name        string `json:"name"`
	Type        int8   `json:"type"`
	Password    string `json:"password"`
	App_id      string `json:"app_id"`
	Status      int8   `json:"status"`
	Is_password int8   `json:"is_password"`
	Is_merger   int8   `json:"-"`
	Apk_id      int32  `json:"-"`
	Ipa_id      int32  `json:"-"`
	Apk_name    string `json:"-"`
	Ipa_name    string `json:"-"`
	Versions    string `json:"versions"`
	Size        int64  `json:"size"`
	App_url     string `json:"app_url"`
	Shot_url    string `json:"shot_url"`
	Desc        string `json:"desc"`
	Allow_count int32 `json:"allow_count"`
	Updated     string `json:"updated"`
	Created     string `json:"-"`
	Last_ip     string `json:"-"`
}

type StaticAndApp struct {
	App_info `xorm:"extends"`
	Down_count   int32  `json:"down_count" `   //  下载次数
	Upload_count int32  `json:"upload_count" ` //  上传次数
	Scan_count   int32  `json:"scan_count" `   // 浏览次数
}

func (StaticAndApp) TableName() string {
	return "app_info"
}

func (App_info) TableName() string {
	return "app_info"
}

func main(){
	 list := make([]StaticAndApp, 0)
	 sql:= "select a.*, sum(down_count) as down_count,sum(scan_count) as scan_count, sum(upload_count) as upload_count "+
	 	  "from app_info as a  left join  app_statistics as s on  a.id = s.app_id "
	 	  //"where s.app_type = a.type"
	 err:=engine.Sql(sql).Find(&list)


	  data,_:=json.Marshal(list)
	 fmt.Println(err,list,"\n",string(data))


}

