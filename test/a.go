package main

import (
	"fmt"
	"gitee.com/mryy1996/parse-apk/src"
	"io/ioutil"
)

func main() {

	a := src.NewApk("file/jumpqpyx_177271.apk")
	//a := src.NewApk("file/me.zhuque.tdandroid.apk")

	info := a.Parse()

	fmt.Println(info.Label, a.ErrorList)

	fmt.Println(info.TargetSdk, info.TargetSdkName)
	fmt.Println(info.MinSdk, info.MinSdkName)

	ioutil.WriteFile("33.png", info.Icon, 0644)

}
