package std

import (
	"github.com/tendermint/tendermint2/pkgs/amino"
)

var Package = amino.RegisterPackage(amino.NewPackage(
	"github.com/tendermint/tendermint2/pkgs/std",
	"std",
	amino.GetCallersDirname(),
).WithDependencies().WithTypes(

	// Account
	&BaseAccount{}, "BaseAccount",

	// MemFile/MemPackage
	MemFile{}, "MemFile",
	MemPackage{}, "MemPackage",

	// Errors
	InternalError{}, "InternalError",
	TxDecodeError{}, "TxDecodeError",
	InvalidSequenceError{}, "InvalidSequenceError",
	UnauthorizedError{}, "UnauthorizedError",
	InsufficientFundsError{}, "InsufficientFundsError",
	UnknownRequestError{}, "UnknownRequestError",
	InvalidAddressError{}, "InvalidAddressError",
	UnknownAddressError{}, "UnknownAddressError",
	InvalidPubKeyError{}, "InvalidPubKeyError",
	InsufficientCoinsError{}, "InsufficientCoinsError",
	InvalidCoinsError{}, "InvalidCoinsError",
	InvalidGasWantedError{}, "InvalidGasWantedError",
	OutOfGasError{}, "OutOfGasError",
	MemoTooLargeError{}, "MemoTooLargeError",
	InsufficientFeeError{}, "InsufficientFeeError",
	TooManySignaturesError{}, "TooManySignaturesError",
	NoSignaturesError{}, "NoSignaturesError",
	GasOverflowError{}, "GasOverflowError",
))
