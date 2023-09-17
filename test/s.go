package main

import (
	"fmt"
	"gitee.com/mryy1996/parse-apk/src"
)

func main() {

	s := src.NewSdk()

	s.LoadPermissionList()

	//fmt.Println(s.GetByKey("android.permission.ACCESS_WIFI_STATE"))
	fmt.Println(s.GetPermissionByKey("android.permission.ACCESS_COARSE_LOCATION"))

}
