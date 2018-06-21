package main

import (
	"encoding/json"
	"fmt"
)

type IpInfo struct {
	IP   string
	Host string
	Port int
}

type UserObj struct {
	Ret []struct {
		ID               int    `json:"id"`
		ParentID         int    `json:"parent_id"`
		Parent           int    `json:"parent"`
		AllParents       []int  `json:"all_parents"`
		Domain           int    `json:"domain"`
		LastLogin        string `json:"last_login"`
		Currency         string `json:"currency"`
		PasswordExpireAt string `json:"password_expire_at"`
		PasswordReset    bool   `json:"password_reset"`
		LastBank         int    `json:"last_bank"`
		Username         string `json:"username"`
		Block            bool   `json:"block"`
		Bankrupt         bool   `json:"bankrupt"`
		Cash             struct {
			ID          int     `json:"id"`
			UserID      int     `json:"user_id"`
			Balance     float64 `json:"balance"`
			PreSub      int     `json:"pre_sub"`
			PreAdd      int     `json:"pre_add"`
			Currency    string  `json:"currency"`
			LastEntryAt string  `json:"last_entry_at"`
		} `json:"cash"`
		CashFake    interface{} `json:"cash_fake"`
		Card        interface{} `json:"card"`
		EnabledCard interface{} `json:"enabled_card"`
	} `json:"ret"`
	Result  string `json:"result"`
	Profile struct {
		ExecutionTime int    `json:"execution_time"`
		ServerName    string `json:"server_name"`
	} `json:"profile"`
}

func main() {
	Q1()
	Q2()
}

func Q1() {
	apiResult := getApiResult()

	//將回傳值塞入struct
	var result UserObj
	error := json.Unmarshal(apiResult, &result)
	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println("Q1：")
	fmt.Printf("id：%v\n", result.Ret[0].ID)
	fmt.Printf("domain：%v\n", result.Ret[0].Domain)
	fmt.Printf("username：%v\n", result.Ret[0].Username)
	fmt.Printf("cash：%+v\n", result.Ret[0].Cash)
}

func Q2() {
	/* Q2-1
	IP:192.17.55.3	Host:docs.google.com 		Port:80
	IP:192.17.33.17 Host:ts-phpadmin.vir999.com Port:80
	IP:192.17.99.52 Host:jsonviewer.stack.hu 	Port:7711
	*/

	//印出作業2-1
	fmt.Println("\nQ2-1：")
	//將3組資料分別印出
	ipinfo1 := IpInfo{IP: "192.17.55.3", Host: "docs.google.com", Port: 80}
	fmt.Printf("第一組：%+v\n", ipinfo1)
	ipinfo2 := IpInfo{IP: "192.17.33.17", Host: "ts-phpadmin.vir999.com", Port: 80}
	fmt.Printf("第二組：%+v\n", ipinfo2)
	ipinfo3 := IpInfo{IP: "192.17.99.52", Host: "jsonviewer.stack.hu", Port: 7711}
	fmt.Printf("第三組：%+v\n\n", ipinfo3)

	/* Q2-2
	IP:177.18.2.33 Host:github.com Port:80
	*/

	//印出作業2-2
	fmt.Println("\nQ2-2：")
	//宣告slice
	var IpInfoList []IpInfo
	//將資料塞入slice
	IpInfoList = append(IpInfoList, ipinfo1)
	IpInfoList = append(IpInfoList, ipinfo2)
	IpInfoList = append(IpInfoList, ipinfo3)
	//新增一筆資料
	ipinfo4 := IpInfo{IP: "177.18.2.33", Host: "github.com", Port: 80}
	IpInfoList = append(IpInfoList, ipinfo4)
	fmt.Printf("IP群組：%+v\n", IpInfoList)

	// Q2-3
	//updateSetting
	fmt.Println("\nQ2-3：")
	checkString := "jsonviewer.stack.hu" //需要判斷的參數
	//取回要修改的指標位置
	portPointer := updateSetting(IpInfoList, checkString)
	//修改數值
	*portPointer = 80
	fmt.Printf("修改port後：%+v\n", IpInfoList)
}

//取得需要修改的指標位置
func updateSetting(ipArray []IpInfo, checkString string) *int {
	var port *int
	//取得要改變數值的指標位置
	for j := 0; j < len(ipArray); j++ {
		//與傳入的參數做比對,取出需要的指標
		if ipArray[j].Host == checkString {
			port = &ipArray[j].Port
			break
		}
	}

	return port
}

//Api的回傳json
func getApiResult() []byte {
	userData := `{"ret":[{"id":148500286,"parent_id":1414663,"parent":1414663,"all_parents":[1414663,59735,58971,24367,6],"domain":6,"last_login":"2018-06-12T09:17:51+0800","currency":"CNY","password_expire_at":"2020-09-04T18:19:02+0800","password_reset":false,"last_bank":54,"username":"wesley01","block":false,"bankrupt":false,"cash":{"id":71814701,"user_id":148500286,"balance":4923642.0545,"pre_sub":0,"pre_add":0,"currency":"CNY","last_entry_at":"20180612155900"},"cash_fake":null,"card":null,"enabled_card":null}],"result":"ok","profile":{"execution_time":14,"server_name":"ipl-web01.rd5.qa"}}`

	return []byte(userData)
}
