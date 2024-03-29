package api

import (
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//	接受用户名和密码 返回对应的json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		log.Println(err)
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
