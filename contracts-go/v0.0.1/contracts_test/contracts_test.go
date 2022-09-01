package contracts_test

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/eluv-io/contracts-evm-builds/contracts-go/v0.0.1/contracts"
)

var testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

// Test_TenantOwner is a simple test showing how to set up a 'simulated' backed
// and exercise contracts functions.
func Test_TenantOwner(t *testing.T) {
	backend := backends.NewSimulatedBackend(
		core.GenesisAlloc{
			crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(10000000000000000)},
		},
		10000000,
	)
	defer func() { _ = backend.Close() }()

	// just because it's easier, the caller uses the same private key as the backend
	// 1337 is the conventional chain_id for tests
	auth, _ := bind.NewKeyedTransactorWithChainID(testKey, big.NewInt(1337))

	// generate tenant id
	var tenantId [32]byte
	_, _ = rand.Read(tenantId[:])

	// deploy contract
	addr, tx, tenant, err := contracts_v0_0_1.DeployTenant(auth, backend, tenantId)
	require.NoError(t, err)
	// simulated backend requires committing after one or more transactions.
	backend.Commit()

	// and wait it's mined
	minedAddr, err := bind.WaitDeployed(context.Background(), backend, tx)
	require.NoError(t, err)
	require.Equal(t, addr, minedAddr)

	// call some function and verify it works
	o, err := tenant.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, crypto.PubkeyToAddress(testKey.PublicKey), o)
}
