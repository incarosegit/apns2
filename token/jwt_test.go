package token_test

import (
	"testing"
	"fmt"
	"github.com/sideshow/apns2/token"
)



// push token: <7af020be 581c955d 1ab6cb92 2fc815dd d9a46e16 cf0efe2c 9668bfa2 150d9f29>
func TestAPNSToken_Generation(t *testing.T) {

	key, err := token.ECKeyFromFile("AuthKey_L9KK6GBQT9.p8")

	apns := token.NewToken(key,"1234567", "issuer")

	if t,_:=apns.Generate();t {
		fmt.Println(apns.Raw())
	}

}
