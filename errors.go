package checkers

import "cosmossdk.io/errors"

var (
	ErrInvalidBlack     = errors.Register(ModuleName, 2, "black address is invalid: %s")
	ErrInvalidRed       = errors.Register(ModuleName, 3, "red address is invalid: %s")
	ErrGameNotParseable = errors.Register(ModuleName, 4, "game cannot be parsed")
)
