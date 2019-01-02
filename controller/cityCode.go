package main
import (
    "io/ioutil"
    "net/http"
    "net/url"
    "fmt"
    "encoding/json"
)
 
 
const APPKEY = "" //您申请的APPKEY
 
func main(){
 
    //1.站点简码查询
    CityCode()
 
}
 
//1.站点简码查询
func CityCode(){
    //请求地址
    juheURL :="http://op.juhe.cn/trainTickets/cityCode"
 
    //初始化参数
    param:=url.Values{}
 
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("key",APPKEY) //您申请的appKey
    param.Set("dtype","json") //返回的格式，json或xml，默认json
    param.Set("stationName","") //站点名，如苏州、苏州北，不需要加“站”字
    param.Set("all","1") //如果需要全部站点简码，请将此参数设为1

 
 
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
        }
    }
}
 
 
// get 网络请求
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
 
// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
    resp,err:=http.PostForm(apiURL, params)
    if err!=nil{
        return nil ,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
