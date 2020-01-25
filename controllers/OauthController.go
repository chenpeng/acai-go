package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/astaxie/beego"
)

type OauthController struct {
	beego.Controller
}

// @Title 下发公钥
// @Description get publicKey
// @Success 200
// @router /publicKey [get]
func (mrc *OauthController) PublicKey() {

}

func GenRsaKey(bits int) error {

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	fmt.Printf("=======私钥文件内容=========%v", string(pem.EncodeToMemory(priBlock)))
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	fmt.Printf("=======公钥文件内容=========%v", string(pem.EncodeToMemory(publicBlock)))

	if err != nil {
		return err
	}
	return nil
}
