package v5

import (
	"github.com/WaveHashProtocol/wavehash/v6/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/gaia/v9/x/globalfee"
	ibchookstypes "github.com/osmosis-labs/osmosis/x/ibc-hooks/types"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v5"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV5UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			globalfee.ModuleName,
			ibchookstypes.StoreKey,
		},
	},
}
