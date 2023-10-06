package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"github.com/alice/checkers"
	"github.com/alice/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	k Keeper
}

var _ checkers.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) checkers.MsgServer {
	return &msgServer{k: keeper}
}

// CreateGame defines the handler for the MsgCreateGame message.
func (ms msgServer) CreateGame(ctx context.Context, msg *checkers.MsgCreateGame) (*checkers.MsgCreateGameResponse, error) {
	if _, err := ms.k.StoredGameList.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
	}

	newBoard := rules.New()
	storedGame := checkers.StoredGame{
		Index: msg.Index,
		Board: newBoard.String(),
		Turn:  rules.PieceStrings[newBoard.Turn],
		Black: msg.Black,
		Red:   msg.Red,
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.StoredGameList.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	ctx1 := sdk.UnwrapSDKContext(ctx)
	ctx1.EventManager().EmitEvent(
		sdk.NewEvent("game-created", sdk.NewAttribute("black", msg.Black)),
	)

	return &checkers.MsgCreateGameResponse{}, nil
}
