package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"github.com/swaggo/swag/example/celler/model"
)
const APPKEY = "5b433b1f92d41bba340a5bb47464ce32" //您申请的APPKEY
// ShowBottle godoc
// @Summary 查询站点的简码
// @Description 查询站点的简码，一般不会变，请做好缓存
// @Param  stationName path string true "站点名，如苏州、苏州北，不需要加“站”字" 
// @Router /cityCode [post]
func (c *Controller) cityCode(ctx *gin.Context) {
        //请求地址
	juheURL :="http://op.juhe.cn/trainTickets/cityCode"
	stationName :=c.Param("stationName")
	//初始化参数
	param:=url.Values{}
	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("key",APPKEY) //您申请的appKey
	param.Set("dtype","json") //返回的格式，json或xml，默认json
	param.Set("stationName",stationName) //站点名，如苏州、苏州北，不需要加“站”字
	param.Set("all","1") //如果需要全部站点简码，请将此参数设为1
	//发送请求
	data,err:=Post(juheURL,param)
	if err!=nil{
		fmt.Errorf("请求失败,错误信息:\r\n%v",err)
		ctx.JSON(http.StatusBadRequest, err)
	}else{
		var netReturn map[string]interface{}
		json.Unmarshal(data,&netReturn)
		if netReturn["error_code"].(float64)==0{
			fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
			ctx.JSON(http.StatusOK, netReturn["result"])
		}
		}
	
	
}

//发送get请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
	var Url *url.URL
	Url,err=url.Parse(apiURL)
	if err!=nil{
		fmt.Printf("解析url错误:\r\n%v",err)
		return nil,err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery=params.Encode()
	resp,err:=http.Get(Url.String())
	if err!=nil{
		fmt.Println("err:",err)
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
//发送post请求
func Post(apiURL string, params url.Values)(rs[]byte,err error){
	resp,err:=http.PostForm(apiURL, params)
	if err!=nil{
		return nil ,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
