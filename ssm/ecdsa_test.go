package ssm

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common"
)

func TestCreateECDSAKeyPair(t *testing.T) {
	privKey, pubKey, cpubKey, _ := CreateECDSAKeyPair()
	fmt.Println("privKey=", privKey)
	fmt.Println("pubKey=", pubKey)
	fmt.Println("cpubKey=", cpubKey)
}

// privKey= 79fa51c9db7fbcf50532656d622bc28196c903d77075033207e3379848f64a27
// pubKey= 04b9070ec4f5e9953f5a2c3fa2f5a5f6acbec5c149567f1647fa65db140e5ee8c7dd1fb47ea9118b4615d994add55c8912d5455576284102f09c148387bbd18f76
// cpubKey= 02b9070ec4f5e9953f5a2c3fa2f5a5f6acbec5c149567f1647fa65db140e5ee8c7

func TestSignMessage(t *testing.T) {
	privKey := "79fa51c9db7fbcf50532656d622bc28196c903d77075033207e3379848f64a27"
	message := common.Hash{}.String()
	signature, err := SignMessage(privKey, message)
	if err != nil {
		fmt.Println("sign tx fail")
	}
	fmt.Println("Signature: ", signature)
}

func TestVerifySign(t *testing.T) {
	msgHash := "0000000000000000000000000000000000000000000000000000000000000000"
	publicKey := "04b9070ec4f5e9953f5a2c3fa2f5a5f6acbec5c149567f1647fa65db140e5ee8c7dd1fb47ea9118b4615d994add55c8912d5455576284102f09c148387bbd18f76"
	signature := "1227a6564ed1fc70fcce1822d129f8f0109c54a554ce91854c06451933c44c125a3b46aa4d21d96eb51e83b5f41d73b5f1c1af87b135f7de58a7ff9d77a23f7000"
	ok := VerifySign(publicKey, msgHash, signature)
	fmt.Println("ok==", ok)
}
