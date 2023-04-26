package domain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type KeyPair struct {
	PublicKey  string `bson:"publicKey"`
	PrivateKey string `bson:"privateKey"`
}

func NewKeyPair() (*KeyPair, error) {

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	privateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	publicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
	})

	return &KeyPair{
		PublicKey:  string(publicKey),
		PrivateKey: string(privateKey),
	}, nil
}

func ConvertKeyStringToPemBlock(key string) *pem.Block {
	block, _ := pem.Decode([]byte(key))
	return block
}
