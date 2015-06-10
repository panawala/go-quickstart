package helper

import (
	"fmt"
	. "github.com/qiniu/api/conf"
	qiuniu_io "github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"github.com/satori/go.uuid"
	"io"
	"strings"
)

func uptoken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
		//CallbackUrl: callbackUrl,
		//CallbackBody:callbackBody,
		//ReturnUrl:   returnUrl,
		//ReturnBody:  returnBody,
		//AsyncOps:    asyncOps,
		//EndUser:     endUser,
		//Expires:     expires,
	}
	return putPolicy.Token(nil)
}

func init() {
	ACCESS_KEY = ""
	SECRET_KEY = ""
}

func UploadFile(data io.Reader, size int64, mime_type string) string {
	var err error
	var ret qiuniu_io.PutRet
	uptoken := uptoken("pan2")
	fmt.Println(uptoken)
	extra := &qiuniu_io.PutExtra{}
	extra.MimeType = mime_type
	fmt.Println(size)
	var key string = uuid.NewV4().String() + "." + strings.Split(mime_type, "/")[1]
	err = qiuniu_io.Put2(nil, &ret, uptoken, key, data, size, extra)
	if err != nil {
		fmt.Println("io.Put failed", err)
		fmt.Println(ret)
		return ""
	}
	cdn_url := "http://7xjfrz.com1.z0.glb.clouddn.com/"
	fmt.Println(ret)
	fmt.Println(cdn_url + ret.Key)
	fmt.Println("upload success")
	return cdn_url + ret.Key
}
