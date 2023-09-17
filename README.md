### 解析安卓apk

**安装**

```shell
go get gitee.com/mryy1996/parse-apk
```

**快速开始**
```go
package main

import (
	"fmt"
	"gitee.com/mryy1996/parse-apk/src"
	"io/ioutil"
)

func main() {

	a := src.NewApk("file/jumpqpyx_177271.apk")

	info := a.Parse()

	//app名称
	fmt.Println(info.Label)

	//目标sdk
	fmt.Println(info.TargetSdk, info.TargetSdkName)

	//最小运行sdk
	fmt.Println(info.MinSdk, info.MinSdkName)

	//权限列表
	fmt.Println(info.PermissionList)

	//版本号
	fmt.Println(info.VersionCode, info.VersionName)

	//文件大小、文件md5
	fmt.Println(info.Size, info.Md5)

	//保存icon
	ioutil.WriteFile("33.png", info.Icon, 0644)

}
```
