package src

type Info struct {
	Md5            string
	Size           int64
	Label          string
	Package        string
	VersionCode    int32
	VersionName    string
	TargetSdk      int32
	TargetSdkName  string
	MinSdk         int32
	MinSdkName     string
	PermissionList []Permission
}
