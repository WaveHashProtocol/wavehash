package interchaintest

import (
	"context"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestBasicWavehashStart is a basic test to assert that spinning up a Wavehash network with one validator works properly.
func TestBasicWavehashStart(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	// Create chain factory with Wavehash
	numVals := 1
	numFullNodes := 1

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "wavehash",
			ChainConfig:   wavehashConfig,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	wavehash := chains[0].(*cosmos.CosmosChain)

	// Create a new Interchain object which describes the chains, relayers, and IBC connections we want to use
	ic := interchaintest.NewInterchain().AddChain(wavehash)

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	ctx := context.Background()
	client, network := interchaintest.DockerSetup(t)

	err = ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: true,

		// This can be used to write to the block database which will index all block data e.g. txs, msgs, events, etc.
		// BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})
}
