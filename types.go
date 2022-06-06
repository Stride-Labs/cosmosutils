package cosmosutils

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	// icacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	// icahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	interchainquerytypes "github.com/Stride-Labs/stride/x/interchainquery/types"
	stakeibcmoduletypes "github.com/Stride-Labs/stride/x/stakeibc/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/v3/modules/core/types"
	nfttypes "github.com/crypto-org-chain/chain-main/v3/x/nft/types"
)

func RegisterDefaultInterfaces(interfaceRegistry types.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)

	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	proposal.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	ibctypes.RegisterInterfaces(interfaceRegistry)
	ibctransfertypes.RegisterInterfaces(interfaceRegistry)
	nfttypes.RegisterInterfaces(interfaceRegistry)
	authztypes.RegisterInterfaces(interfaceRegistry)
	feegranttypes.RegisterInterfaces(interfaceRegistry)
	// icacontrollertypes.RegisterInterfaces(interfaceRegistry)
	// icahosttypes.RegisterInterfaces(interfaceRegistry)
	icatypes.RegisterInterfaces(interfaceRegistry)
	stakeibcmoduletypes.RegisterInterfaces(interfaceRegistry)
	interchainquerytypes.RegisterInterfaces(interfaceRegistry)
}
