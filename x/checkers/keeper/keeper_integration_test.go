package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/xavierlepretre/checkers/app"
	"github.com/xavierlepretre/checkers/x/checkers/keeper"
	"github.com/xavierlepretre/checkers/x/checkers/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	app         *app.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestMsgServerPlayMoveTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := app.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.CheckersKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.msgServer = keeper.NewMsgServerImpl(app.CheckersKeeper)
	suite.ctx = ctx
	suite.queryClient = queryClient
}

func (suite *IntegrationTestSuite) TestGenesisNextGame() {
	nextGame, err := suite.queryClient.NextGame(sdk.WrapSDKContext(suite.ctx), &types.QueryGetNextGameRequest{})
	suite.Require().Nil(err)
	suite.Require().Equal(uint64(1), nextGame.NextGame.IdValue)
}

func (suite *IntegrationTestSuite) TestCreateGameReturnsNextGameId() {
	suite.app.CheckersKeeper.SetNextGame(suite.ctx, types.NextGame{
		Creator: "",
		IdValue: 23,
	})
	black := sdk.AccAddress([]byte("addr1_______________"))
	red := sdk.AccAddress([]byte("addr2_______________"))

	response, err := suite.msgServer.CreateGame(sdk.WrapSDKContext(suite.ctx), &types.MsgCreateGame{
		Creator: black.String(),
		Black:   black.String(),
		Red:     red.String(),
	})
	suite.Require().NoError(err)
	suite.Require().EqualValues(&types.MsgCreateGameResponse{IdValue: "23"}, response)

	nextGame, found := suite.app.CheckersKeeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator: "",
		IdValue: 24,
	}, nextGame)
}
