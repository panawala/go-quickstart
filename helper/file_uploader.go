package helper

import (
	"github.com/astaxie/beego"
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
	}
	return putPolicy.Token(nil)
}

func init() {
	ACCESS_KEY = "fuI-VbB3VrpleFvmJYVwTaan60h9Yu_hWgpaJRgd"
	SECRET_KEY = "XdsSJybSynYqQlsuoCoI2sOF5_br-smlB27hfmGH"
}

func UploadFile(data io.Reader, size int64, mime_type string) string {
	var err error
	var ret qiuniu_io.PutRet
	uptoken := uptoken(beego.AppConfig.String("qiniu_bucket"))
	extra := &qiuniu_io.PutExtra{}
	extra.MimeType = mime_type
	var key string = uuid.NewV4().String() + "." + strings.Split(mime_type, "/")[1]
	err = qiuniu_io.Put2(nil, &ret, uptoken, key, data, size, extra)
	if err != nil {
		beego.Error("io.Put2 failed", err)
		return ""
	}
	beego.Info("upload success")
	return beego.AppConfig.String("qiniu_cdn_url") + ret.Key
}
