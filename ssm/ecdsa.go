package ssm

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

/**
 * @brief: CreateECDSAKeyPair
 * @param:
 * @return: string, string, string, error
 */
func CreateECDSAKeyPair() (string, string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Error("generate key fail", "err==", err)
		return "0x00", "0x00", "0x00", err
	}
	priKeyStr := hex.EncodeToString(crypto.FromECDSA(privateKey))
	pubKeyStr := hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey))
	decPubKeyStr := hex.EncodeToString(crypto.CompressPubkey(&privateKey.PublicKey))
	return priKeyStr, pubKeyStr, decPubKeyStr, nil
}

/**
 * @brief: SignMessage
 * @param: privKey string, txMsg string
 * @return: string, error
 */
func SignMessage(privKey, txMsg string) (string, error) {
	hash := common.HexToHash(txMsg)
	log.Info("txMsg hash===", hash)
	privByte, err := hex.DecodeString(privKey)
	if err != nil {
		log.Error("decode private key fail", "err===", err)
		return "0x000", err
	}
	privkeyECDSA, err := crypto.ToECDSA(privByte)
	if err != nil {
		log.Error("Byte private key to ecdsa key fail", "err===", err)
		return "0x000", err
	}
	signatureByte, err := crypto.Sign(hash[:], privkeyECDSA)
	if err != nil {
		log.Error("sign transaction fail", "err===", err)
		return "0x000", err
	}
	return hex.EncodeToString(signatureByte), nil
}

/**
 * @brief: VerifySign
 * @param: pubKey, msgHash, sign string
 * @return: bool
 */
func VerifySign(pubKey, msgHash, sign string) bool {
	publicKey, err := hex.DecodeString(pubKey)
	if err != nil {
		log.Error("decode public key fail", "err===", err)
		return false
	}
	messageHash, err := hex.DecodeString(msgHash)
	if err != nil {
		log.Error("decode msghash key fail", "err===", err)
		return false
	}
	signature, err := hex.DecodeString(sign)
	if err != nil {
		log.Error("decode sign fail", "err===", err)
		return false
	}
	log.Info("publickey length===", len(publicKey))
	log.Info("messageHash length===", len(messageHash))
	log.Info("signature length===", len(signature))
	// 在使用 crypto.VerifySignature 函数时，签名数据应该是 r || s 的组合，不包括恢复字节 v。恢复字节 v 通常是 1 个字节，位于签名的末尾。
	//签名去掉v值,不包括恢复字节
	return crypto.VerifySignature(publicKey, messageHash, signature[:64])
}
