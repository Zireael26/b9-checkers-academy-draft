package checkers

import (
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	k.SetSystemInfo(ctx, genState.SystemInfo)
	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, elem)
	}
	// Set all the playerInfo
	for _, elem := range genState.PlayerInfoList {
		k.SetPlayerInfo(ctx, elem)
	}
	// Set if defined
	k.SetLeaderboard(ctx, genState.Leaderboard)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.StoredGameList = k.GetAllStoredGame(ctx)
	genesis.PlayerInfoList = k.GetAllPlayerInfo(ctx)
	// Get all leaderboard
	leaderboard, found := k.GetLeaderboard(ctx)
	if found {
		genesis.Leaderboard = leaderboard
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
