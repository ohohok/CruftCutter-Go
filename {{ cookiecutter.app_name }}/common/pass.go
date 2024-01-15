package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

const PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAsJbTZtSs41F5Ioq/qcS0Cw23P69SpMSZ2nEYolNhTY3WCf3u\newnlm+mHDZnGt2lP5jQp6c4y0QO+OQtcdSKfpO3ouX56Iu0VPI9U4zDYctqZBmQ5\n7j8LhN9bFnmEMw8YwmUNLWxOgO8REltjXj5wbHCxQFHBz2MFA2uulDWh5LaEn6cQ\nG0qrwm1UkoFeoIsMEpIyucIIFua2E6ZFpCdDe7sStikgkUvlFboJ4xH7RWjZiH8w\ngrOnA9syY49IsQJVijrc/MfFA4noFdPiUuD4J3lbZFDZnV34uvzh3+HRowEoJ2Ei\neyq/mk1/npqTnDLV98B1jo0Lo6IFAUI5gFBdaQIDAQABAoIBAHYUtNHLHrx2a2jk\nnJr59GZ7ynBkXP/ekv6Vp6JL3QHN+TT/Puu5R3cFJhC7JjzYx9uoP+qevi4zsYxV\ng4K0H0pa58a5wxP9srinI10z+Vh7wd2bQX2FXL+B11fk87hsvOKoTbJ0/N2Mfr7m\nq0CGwghRJsVaph12GhEafUePwWy8W/HmZ4n8B+NGMIoXmoLT2NNrqR0D952mn/iS\nHkM1hp9ht8s7BgHodXqFrR9OqpmVNBBSYA62vNuHi2rfo7L5S+mJnFFHAGBvkOUT\nV/6Kpb0essJkyQexeIl3Wjr+KXeeoF+LiLkezEnLJRPPf2dsLvj+cQIQEck5ZaOR\nDbbmbG0CgYEAt5pgT9NMZaXgpOAlZhPUzWiZDNuz8T7SxmJpm75lxfSUR7wCFGvf\nQiTzaTiaaoqMqseuSWj5igVY3nXlu0hqh2FhZcqf0BhoZnHfrblll6s8ZrVkcjF1\nqST+fXLXv1Y2Xi0VYbZ2k1ecFNGDLM2kClRDqDrthBpOV37uFgFcXi8CgYEA9jhx\nDdJJMDcSAZntsyXg8ZNYdqjlMByO+dsQGTP1q1JIRsBBtpzZVhNni2z2EsamVgkR\nuEXnKrP6g9DnRDJwzRLwq2vUQlXyRMWwMLum601zFGWx3iycI82JfYaPMeMglHeH\nyVEYCLddzNSlrinZIOialroz37z+1v82w9mvb+cCgYEArZIf4tcari22xaexM2Mk\nFlv/9Ivu/O4rTKjUtgu1IwMo+vfd73hbQ6izBJIiSP5aJUlIltXKspSDgytp1zeK\nnhmNfjGkC5JHgDG/B/jw9gVHwMFTCDGBzjnO7MGY/KWAGF/+irV4O6rjPzsiC0UN\nA0bN/0hWKkDENysj4WG/9LkCgYBHao2YJbNm0cJeRoiEmusJBuT03f5EGzR5Ukn0\nFXcfbylZpDArSIldhxlUfpFNVuMuN0k1eskXQMbb7v71b+/5+5FlF0ykxQsVWYXG\nTNeYjolflDOoLqZHWDmV+C6WmLt9dMIk6WmgNZd2bwNzZd39xpn4OCANuGLud24q\nFnzlZwKBgAoEWphjO3EzAt1w7dG22NTK8Y9FQDjerKQg3RkxNiWoFcb7Qlvh3m0K\n9OWPxEi6WHOCBeLYOQhoQ99GsRx6LOIfKSxeqLFneNFAaPcHWBDV781nUeGi5WPB\nA5JxJLGGPIE38DIWT8/IrZ5k90ZaHUDDwMGKFhXhTgnz9QDm06O7\n-----END RSA PRIVATE KEY-----"

func RSADecrypt(cipherText string, privateKey []byte) ([]byte, error) {
	// 先base64解码
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	//pem解码
	block, _ := pem.Decode(privateKey)
	//X509解码
	privateKeyItem, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyItem, data)
	if err != nil {
		return nil, err
	}
	//返回明文
	return plainText, nil
}
