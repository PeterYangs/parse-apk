package main

import (
	"fmt"
	"gitee.com/mryy1996/parse-apk/src"
)

func main() {

	a := src.NewApk("file/jumpqpyx_177271.apk")

	info := a.Parse()

	fmt.Println(info.Label, a.ErrorList, info.PermissionList)

}
