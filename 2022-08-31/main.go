package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

//var token = "c87d80b350c25c0ac277d7b98703be26"
var token = "a118416cfcd54a5ac8b59cd87bbdcccd"

var baseUrl = "https://apis.jobui.com/cooperation/yizhanchi/"
//var baseUrl = "http://xcxapi.jobui.com/cooperation/yizhanchi/"

type constantData struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data struct {
		JobList []struct {
			Category string     `json:"category"`
			List     []listItem `json:"list"`
		} `json:"jobList"`
		DistrictList []struct {
			Province string     `json:"province"`
			CityList []cityList `json:"cityList"`
		} `json:"districtList"`
		EduList []string `json:"eduList"`
		ExpList []string `json:"expList"`
	} `json:"data"`
	Error interface{} `json:"error"`
}

//职位列表 结构体
type listItem struct {
	Job   string `json:"job"`
	JobID string `json:"jobID"`
}

//城市列表结构体
type cityList struct {
	City     string   `json:"city"`
	AreaList []string `json:"areaList"`
}

//请求响应数据 结构体
type responseData struct {
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

//常量数据变量
var constants constantData

var wait sync.WaitGroup

//原子计数器
var errorNum uint64 = 0

//主函数
func main() {
	//获取常量数据
	getConstant()

	//变量常量数组
	if constants.Data.JobList != nil {
		for _, list := range constants.Data.JobList {
			wait.Add(1)
			//go operationApi(list.List)
			operationApi(list.List)
		}
	}

	wait.Wait()

	errorFinal := atomic.LoadUint64(&errorNum)
	fmt.Println("errorNum:", errorFinal)

	var writeString = "errorNum:" + strconv.FormatUint(errorFinal, 10) + "\n"
	var filename = "./online_test_bug.txt"
	var f *os.File

	defer f.Close()
	//使用 io.WriteString
	if checkFileIsExist(filename) { //如果文件存在
		f, _ = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, _ = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}

	n, _ := io.WriteString(f, writeString) //写入文件(字符串)

	fmt.Printf("写入 %d 个字节n", n)
}

//
//  checkFileIsExist
//  @Description: 判断文件是否存在  存在返回 true 不存在返回false
//  @param filename
//  @return bool
//
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//
//  getConstant
//  @Description: 
//
func getConstant() {
	urls := baseUrl + "constant/"
	urlValues := url.Values{}
	urlValues.Add("token", token)
	resp, _ := http.PostForm(urls, urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &constants)
}

//
//  requestApi
//  @Description: 请求内容
//  @param requestUrl
//  @param jobName
//  @param city
//  @return responseData
//
func requestApi(requestUrl string, jobName string, city string) responseData {
	//wait.Add(1)
	urlValues := url.Values{}
	urlValues.Add("job", jobName)
	urlValues.Add("token", token)
	if city != "" {
		urlValues.Add("city", city)
	}
	resp, _ := http.PostForm(requestUrl, urlValues)
	body, _ := ioutil.ReadAll(resp.Body)

	var data responseData

	json.Unmarshal(body, &data)

	fmt.Println(data)
	if data.Meta.Code != 200 {
		// 使用 `AddUint64` 来让计数器自动增加，使用
		// `&` 语法来给出 `errorNum` 的内存地址。
		atomic.AddUint64(&errorNum, 1)

		// 允许其它 Go 协程的执行
		runtime.Gosched()
	}

	//wait.Done()
	return data
}

//
//  operationApi
//  @Description: 操作api请求
//  @param list
//
func operationApi(list []listItem) {
	for _, item := range list {
		fmt.Println(item.Job)
		//for _, infoList := range constants.Data.DistrictList {
		//	for _, cityInfoList := range infoList.CityList {
		//		go requestApi(baseUrl+"trends/",item.Job,cityInfoList.City)
		//		go requestApi(baseUrl+"salary/",item.Job,cityInfoList.City)
		//		//for _,area := range cityInfoList.AreaList {
		//		//	
		//		//}
		//		
		//	}
		//}
		requestApi(baseUrl+"skill/",item.Job,"")
		requestApi(baseUrl+"major/",item.Job,"")
		requestApi(baseUrl+"experience/",item.Job,"")
		requestApi(baseUrl+"education/",item.Job,"")
		requestApi(baseUrl+"trends/",item.Job,"")
		requestApi(baseUrl+"salary/", item.Job, "")
		time.Sleep(time.Second)
	}

	wait.Done()
}



