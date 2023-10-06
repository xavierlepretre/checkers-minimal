package keeper

import (
	"context"

	"github.com/alice/checkers"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *checkers.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, storedGame := range data.StoredGameList {
		if err := k.StoredGameList.Set(ctx, storedGame.Index, storedGame); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*checkers.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var storedGames []checkers.StoredGame
	if err := k.StoredGameList.Walk(ctx, nil, func(index string, storedGame checkers.StoredGame) (bool, error) {
		storedGames = append(storedGames, storedGame)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &checkers.GenesisState{
		Params:         params,
		StoredGameList: storedGames,
	}, nil
}
