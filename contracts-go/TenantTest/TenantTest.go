package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	//tenant "Tenant"
	tenant "github.com/eluv-io/contracts-evm-builds/contracts-go/Tenant"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	qrypto "github.com/qluvio/elv-master/crypto"
	"github.com/qluvio/elv-master/format"
)

var PRIVATE_KEY = ""
var ADDRESS_SMART_CONTRACT = "0x6958a6Ff096CEB64CA51124377F3Fd77fc3200E8"

func main() {
	contentId := byte32PutString("content1")

	testTenantCommitEIP191(contentId, true, "commit1")

}

func byte32PutString(s string) [32]byte {
	var a [32]byte
	if len(s) > 32 {
		copy(a[:], s)
	} else {
		copy(a[32-len(s):], s)
	}
	return a
}

func deploy(tenantId [32]byte) {
	client, err := ethclient.Dial("https://host-766.contentfabric.io/eth")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := tenant.DeployTenant(auth, client, tenantId)
	if err != nil {
		log.Fatal(err)
	}

	ADDRESS_SMART_CONTRACT = address.Hex()
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}

func signTenantCommitEIP191(contentId [32]byte, force bool, newHash string) (sigV uint8, sigR, sigS [32]byte) {
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}
	data, err := format.AbiEncodePacked([]interface{}{contentId, force, newHash})
	data = crypto.Keccak256(data)
	data = []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data))
	data = crypto.Keccak256(data)

	signature, err := crypto.Sign(data, privateKey)

	sigV, sigR, sigS = qrypto.EthSplitSignature(signature)

	return
}

func testTenantCommitEIP191(contentId [32]byte, force bool, newHash string) {
	sigV, sigR, sigS := signTenantCommitEIP191(contentId, force, newHash)

	client, err := ethclient.Dial("https://host-766.contentfabric.io/eth")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(ADDRESS_SMART_CONTRACT)

	instance, err := tenant.NewTenant(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.commitSigned(contentId, force, newHash, sigV, sigR, sigS)
	if err != nil {
		log.Fatal(err)
	}
}

func signTenantCommit(contentId [32]byte, force bool, newHash string) (sigV uint8, sigR, sigS [32]byte) {
	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}
	data, err := format.AbiEncodePacked([]interface{}{contentId, force, newHash})
	data = crypto.Keccak256(data)

	signature, err := crypto.Sign(data, privateKey)

	sigV, sigR, sigS = qrypto.EthSplitSignature(signature)

	return
}

func testTenantCommit(contentId [32]byte, force bool, newHash string) {
	sigV, sigR, sigS := signTenantCommit(contentId, force, newHash)

	client, err := ethclient.Dial("https://host-766.contentfabric.io/eth")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(ADDRESS_SMART_CONTRACT)

	instance, err := tenant.NewTenant(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.commitSigned(contentId, force, newHash, sigV, sigR, sigS)
	if err != nil {
		log.Fatal(err)
	}
}
