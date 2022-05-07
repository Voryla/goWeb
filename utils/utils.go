package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func HowDo(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello World!%s", request.URL.Path[1:])
}

func CreatSSL() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	// 生成一个很长的整数作为证书序列号
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}
	// 对证书进行配置
	template := x509.Certificate{
		SerialNumber: serialNumber, // 记录由CA分发的唯一号码
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 有效期为1年
		// KeyUsage 和 ExtKeyUsage 字段值表明这个X.509证书是用于服务器身份验证操作的
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // 常量位图
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // 扩展密钥用法的序列
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},                           // 设置证书仅在ip 127.0.0.1 之上运行
	}
	// 通过调用 crypto/rsa 标准库中的 GenerateKey 函数生成一个RSA私钥，该私钥包含一个能够公开访问的公钥pk.PublicKey，这个公钥用于创建SSL证书
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	// 生成证书  参数一：随机数Reader、参数二：证书模板、参数三：父级证书，若父级证书为模板本身，则该证书是自签名的 参数四：公钥 参数五：签名者的私钥，
	// 返回值：DER编码的证书切片
	// DER:证书的编码格式 X.509证书可以使用多种格式编码，其中一种编码格式是BER（Basic Encoding Rules，基本编码规则）。
	// BER格式指定了一种自解释并且自定义的格式用于对ASN.1数据结构进行编码，而DER格式则是BER的一个子集。DER只提供了一种编码ASN.1值的方法，
	// 这种方法被广泛地应用于密码学当中，尤其是对X.509证书进行加密。
	// 提示：如果证书是由CA签发的，那么证书文件中将同时包含服务器签名以及CA签名，其中服务器签名在前，CA签名在后
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template,
		&template, &pk.PublicKey, pk)
	// 创建 cert.pem 证书文件
	certOut, _ := os.Create("cert.pem")
	// 使用 encoding/pem 标准库将证书编码到 cert.pem 文件中
	// PEM :证书的保存格式 SSL证书可以以多种不同的格式保存，其中一种是PEM（Privacy Enhanced Email，隐私增强邮件）格式，
	// 这种格式会对DER格式的X.509证书实施Base64编码，并且这种格式的文件都以-----BEGIN CERTIFICATE----- 开头，以-----END CERTIFICATE----- 结尾
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	defer certOut.Close()
	// 私钥
	keyOut, _ := os.Create("key.pem")
	//然后继续以PEM编码的方式把之前生成的密钥编码并保存到key.pem文件 里面
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	defer keyOut.Close()
}
