package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/minio/minio-go/v7"
	"github.com/pquerna/otp/totp"
	"log"
	"strconv"
	"videoweb/biz/model/hertz/user"
	"videoweb/biz/utils"
	"videoweb/config"
	"videoweb/database/DB/dao"
	"videoweb/response"
)

func GetMfa(c context.Context, ctx *app.RequestContext) (interface{}, error) {
	claim, err := utils.ParseToken(string(ctx.GetHeader("access_token")))
	if err != nil {
		return response.BadResponse(), err
	}

	usrDao := dao.NewUserDao(c)
	usr, err := usrDao.FindUserByUid(strconv.FormatInt(claim.Uid, 10))
	if err != nil {
		return response.BadResponse(), err
	}

	img, key, err := dao.GenerateQRCode(usr.Username, claim.Uid)
	if err != nil {
		return response.BadResponse(), err
	}

	imgByBase64, err := dao.ImageToBase64(key, img)
	if err != nil {
		return response.BadResponse(), err
	}

	decodedImg, err := base64.StdEncoding.DecodeString(imgByBase64)
	if err != nil {
		log.Println(err)
		return response.BadResponse(), err
	}

	buf := bytes.NewBuffer(decodedImg)
	storeQRCodePath := "qrcode/" + strconv.FormatInt(claim.Uid, 10) + usr.Username + ".png"
	config.MinioClient.PutObject(context.Background(), config.BucketName, storeQRCodePath, buf, -1, minio.PutObjectOptions{
		ContentType: "image/png", // 设置内容类型为图片
	})

	return response.MFAResponse(key.Secret(), imgByBase64), nil
}

func BindMFA(c context.Context, ctx *app.RequestContext, MFAReq *user.UserBindMFAReq) (interface{}, error) {
	valid := totp.Validate(MFAReq.GetSecret(), MFAReq.GetCode()) //key设置为全局变量

	if !valid {
		log.Println("Invalid passcode!")
		return response.BadResponse(), errors.New("Invalid MFA secret!")
	}
	//codeBytes, err := base64.StdEncoding.DecodeString(MFAReq.GetCode())
	//if err != nil {
	//	return response.BadResponse(), err
	//}

	return response.GoodResponse(), nil
}
