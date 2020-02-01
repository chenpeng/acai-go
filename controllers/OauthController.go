package controllers

import (
	"acai-go/models"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"strings"
)

type OauthController struct {
	beego.Controller
}

var publicKeyStr = ""

// @Title 下发公钥
// @Description get publicKey
// @Success 200
// @router /publicKey [get]
func (mrc *OauthController) PublicKey() {
	if publicKeyStr != "" {
		result := models.Result{Code: 0, Data: publicKeyStr, Message: "获取成功"}
		mrc.Data["json"] = result
	} else {
		publicKeyS, err := GenRsaKey(1024)
		if err != nil {
			result := models.Result{Code: 1, Data: nil, Message: "获取失败"}
			mrc.Data["json"] = result
		} else {
			publicKeyStr = publicKeyS
			result := models.Result{Code: 0, Data: publicKeyStr, Message: "获取成功"}
			mrc.Data["json"] = result
		}
	}
	mrc.ServeJSON()
}

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var privateKeyPem string
var publicKeyPem string

func GenRsaKey(size int) (pkp string, err error) {
	// 生成私钥文件
	privateKey, err = rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		errors.New("出问题啦")
	}
	if bits := privateKey.N.BitLen(); bits != size {
		errors.New("key too short")
	}
	privateKeyB := x509.MarshalPKCS1PrivateKey(privateKey)
	//privateKeyStr = base64.StdEncoding.EncodeToString(privateKeyB)
	priBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyB,
	}
	privateKeyPem = string(pem.EncodeToMemory(priBlock))
	fmt.Printf("=======私钥文件内容=========%v", privateKeyPem)
	// 生成公钥文件
	publicKey = &privateKey.PublicKey
	publicKeyB, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		errors.New("出问题啦")
	}
	//publicKeyStr = base64.StdEncoding.EncodeToString(publicKeyB)
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyB,
	}
	publicKeyPem = string(pem.EncodeToMemory(publicBlock))
	fmt.Printf("=======公钥文件内容=========%v", publicKeyPem)
	if err != nil {
		errors.New("出问题啦")
	}
	return publicKeyPem, err
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//block, _ := pem.Decode(privateKeyB)
	//if block == nil {
	//	return nil, errors.New("private key error!")
	//}
	//priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	//if err != nil {
	//	return nil, err
	//}
	//return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	return rsa.DecryptPKCS1v15(nil, privateKey, ciphertext)
}

// @Title 登录
// @Description post login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200
// @router /login [post]
func (mrc *OauthController) Login() {
	var inputDto models.InputDto
	json.Unmarshal(mrc.Ctx.Input.RequestBody, &inputDto)
	textB, err := RsaDecrypt([]byte(inputDto.Text))
	fmt.Printf(string(textB))
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "获取失败"}
		mrc.Data["json"] = result
	} else {
		result := models.Result{Code: 0, Data: string(textB), Message: "获取成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

var random io.Reader

// @Title 加密
// @Description post login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200
// @router /encrypt [post]
func (mrc *OauthController) Encrypt() {
	var inputDto models.InputDto
	json.Unmarshal(mrc.Ctx.Input.RequestBody, &inputDto)
	text := inputDto.Text
	random = rand.Reader
	//bb,err := rsa.EncryptPKCS1v15(random, publicKey, []byte(text))
	//aa,err := rsa.DecryptPKCS1v15(random, privateKey, bb)
	aa, err := rsa.EncryptPKCS1v15(random, publicKey, []byte(text))
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "加密失败"}
		mrc.Data["json"] = result
	} else {
		str := base64.StdEncoding.EncodeToString(aa)
		result := models.Result{Code: 0, Data: str, Message: "加密成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title 解密
// @Description post login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200
// @router /decrypt [post]
func (mrc *OauthController) Decrypt() {
	var inputDto models.InputDto
	json.Unmarshal(mrc.Ctx.Input.RequestBody, &inputDto)
	text := inputDto.Text
	random := inputDto.Random
	qq, err := base64.StdEncoding.DecodeString(text)
	ab := strings.NewReader(random)
	aa, err := rsa.DecryptPKCS1v15(ab, privateKey, qq)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "获取失败"}
		mrc.Data["json"] = result
	} else {
		str := string(aa)
		result := models.Result{Code: 0, Data: str, Message: "获取成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}

// @Title 测试
// @Description post login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200
// @router /ceshi [post]
func (mrc *OauthController) Ceshi() {
	var inputDto models.InputDto
	json.Unmarshal(mrc.Ctx.Input.RequestBody, &inputDto)
	text := inputDto.Text
	random := inputDto.Random
	ab := strings.NewReader(random)
	bb, err := rsa.EncryptPKCS1v15(ab, publicKey, []byte(text))
	cc := base64.StdEncoding.EncodeToString(bb)
	dd, err := base64.StdEncoding.DecodeString(cc)
	aa, err := rsa.DecryptPKCS1v15(ab, privateKey, dd)
	if err != nil {
		result := models.Result{Code: 1, Data: nil, Message: "获取失败"}
		mrc.Data["json"] = result
	} else {
		str := string(aa)
		//str := base64.StdEncoding.EncodeToString(aa)
		result := models.Result{Code: 0, Data: str, Message: "获取成功"}
		mrc.Data["json"] = result
	}
	mrc.ServeJSON()
}
