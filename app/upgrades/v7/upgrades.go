package v7

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

func CreateUpgradeHandler(mm *module.Manager, configurator module.Configurator,
	wasmKeeper *wasm.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// Set wasm old version to 1 if we want to call wasm's InitGenesis ourselves
		// in this upgrade logic ourselves
		// vm[wasm.ModuleName] = wasm.ConsensusVersion

		// otherwise we run this, which will run wasm.InitGenesis(wasm.DefaultGenesis())
		// and then override it after
		newVM, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return newVM, err
		}

		// Since we provide custom DefaultGenesis (privileges StoreCode) in app/genesis.go rather than
		// the wasm module, we need to set the params here when migrating (is it is not customized).

		params := wasmKeeper.GetParams(ctx)
		params.CodeUploadAccess = wasmtypes.AllowNobody
		wasmKeeper.SetParams(ctx, params)

		// override here
		return newVM, err
	}
}
