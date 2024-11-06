package rpc

import (
	"context"
	"errors"
	"strconv"

	"github.com/CavnHan/wallet-chain-sign/leveldb"
	"github.com/CavnHan/wallet-chain-sign/protobuf/wallet"
	"github.com/CavnHan/wallet-chain-sign/ssm"
	"github.com/ethereum/go-ethereum/log"
)

func (s *RpcServer) GetSupportSignWay(ctx context.Context, in *wallet.SupportSignWayRequest) (*wallet.SupportSignWayResponse, error) {
	if in.Type == "ecdsa" || in.Type == "eddsa" {
		return &wallet.SupportSignWayResponse{
			Code:    strconv.Itoa(1),
			Msg:     "Support this sign way",
			Support: true,
		}, nil
	} else {
		return &wallet.SupportSignWayResponse{
			Code:    strconv.Itoa(0),
			Msg:     "Do not support this sign way",
			Support: false,
		}, nil
	}

}

func (s *RpcServer) ExportPublicKeyList(ctx context.Context, in *wallet.ExportPublicKeyRequest) (*wallet.ExportPublicKeyResponse, error) {
	if in.Number > 10000 {
		return &wallet.ExportPublicKeyResponse{
			Code: strconv.Itoa(1),
			Msg:  "Number must be less than 100000",
		}, nil
	}
	var keyList []leveldb.Key
	var retKeyList []*wallet.PublicKey

	for counter := 0; counter <= int(in.Number); counter++ {
		priKeyStr, pubKeyStr, decPubkeyStr, err := ssm.CreateECDSAKeyPair()
		if err != nil {
			log.Error("create key pair fail", "err", err)
			return nil, err
		}
		keyItem := leveldb.Key{
			PrivateKey:     priKeyStr,
			CompressPubkey: pubKeyStr,
		}
		pukItem := &wallet.PublicKey{
			CompressPubkey:   pubKeyStr,
			DecompressPubkey: decPubkeyStr,
		}
		retKeyList = append(retKeyList, pukItem)
		keyList = append(keyList, keyItem)
	}
	isOk := s.db.StoreKeys(keyList)
	if !isOk {
		log.Error("store keys fail", "isOk", isOk)
		return nil, errors.New("store keys fail")
	}
	return &wallet.ExportPublicKeyResponse{
		Code:      strconv.Itoa(1),
		Msg:       "create keys success",
		PublicKey: retKeyList,
	}, nil
}

func (s *RpcServer) SignMessage(ctx context.Context, in *wallet.SignTxMessageRequest) (*wallet.SignTxMessageResponse, error) {
	priKey, isOk := s.db.GetPrivKey(in.PublicKey)
	if !isOk {
		return nil, errors.New("get private key by public key fail")
	}
	signature, err := ssm.SignMessage(priKey, in.MessageHash)
	if err != nil {
		log.Error("sign message fail", "err", err)
		return nil, err
	}
	return &wallet.SignTxMessageResponse{
		Code:      strconv.Itoa(1),
		Msg:       "sign tx message success",
		Signature: signature,
	}, nil
}
