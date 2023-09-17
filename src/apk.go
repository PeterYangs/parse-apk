package src

import (
	"crypto/md5"
	"fmt"
	"github.com/shogo82148/androidbinary"
	"github.com/shogo82148/androidbinary/apk"
	"io"
	"os"
)

type Apk struct {
	path      string
	ErrorList []error
	Sdk       *Sdk
}

func NewApk(path string) *Apk {

	s := NewSdk()

	s.LoadPermissionList()

	return &Apk{
		path: path,
		Sdk:  s,
	}
}

func (a *Apk) Parse() Info {

	info := Info{}

	md5Str, size, mErr := a.md5FileAndSize(a.path)

	if mErr != nil {

		a.ErrorList = append(a.ErrorList, mErr)

	} else {

		info.Md5 = md5Str
		info.Size = size
	}

	pkg, apkErr := apk.OpenFile(a.path)

	resConfigEN := &androidbinary.ResTableConfig{
		Language: [2]uint8{uint8('z'), uint8('h')},
	}

	if apkErr != nil {

		a.ErrorList = append(a.ErrorList, apkErr)

	} else {

		label, lErr := pkg.Label(resConfigEN)

		if lErr != nil {

			a.ErrorList = append(a.ErrorList, lErr)
		} else {

			info.Label = label

		}

		info.Package = pkg.PackageName()

		code, cErr := pkg.Manifest().VersionCode.Int32()

		if cErr != nil {

			a.ErrorList = append(a.ErrorList, cErr)

		} else {

			info.VersionCode = code
		}

		codeName, nErr := pkg.Manifest().VersionName.String()

		if nErr != nil {

			a.ErrorList = append(a.ErrorList, nErr)

		} else {

			info.VersionName = codeName
		}

		targetSdk, tErr := pkg.Manifest().SDK.Target.Int32()

		if tErr != nil {

			a.ErrorList = append(a.ErrorList, tErr)

		} else {

			info.TargetSdk = targetSdk

		}

		for _, p := range pkg.Manifest().UsesPermissions {

			pe, ee := a.Sdk.GetByKey(p.Name.MustString())

			if ee == nil {

				info.PermissionList = append(info.PermissionList, pe)

			}

		}

	}

	return info

}

func (a *Apk) md5FileAndSize(filename string) (string, int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", 0, err
	}

	info, _ := f.Stat()

	size := info.Size()

	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", size, err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), size, nil
}
