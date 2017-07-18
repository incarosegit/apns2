package apnskey_test

import (
	"testing"
	"fmt"
	"io/ioutil"
//	. "github.com/sideshow/apns2/apnskey"
//	"crypto/x509"
//	"encoding/pem"
	"github.com/sideshow/apns2/apnskey"
)



// push token: <7af020be 581c955d 1ab6cb92 2fc815dd d9a46e16 cf0efe2c 9668bfa2 150d9f29>
func TestAPNSToken_Generation(t *testing.T) {
	bytes, _ := ioutil.ReadFile("AuthKey_L9KK6GBQT9.p8")

	apns, err := apnskey.NewToken(bytes,"1234567", "issuer")
	if err != nil{
		fmt.Println("error: ", err)
	}

	fmt.Println(apns.Generate())

}
