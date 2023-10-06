package checkers

import "cosmossdk.io/collections"

const ModuleName = "checkers"

var (
	ParamsKey         = collections.NewPrefix("Params")
	StoredGameListKey = collections.NewPrefix("StoredGame/value/")
)
