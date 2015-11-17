package main

import (
	"fmt"
)

// ErrorCode holds the type of error return.
type ErrorCode int

// Here is the list of defined error codes.
const (
	InvalidCharacterSheet ErrorCode = iota

	InvalidHeaderLine
	EmptyHeaderKey
	EmptyHeaderValue
	DuplicateHeaderLine
	InvalidBackgroundOptions
	UndefinedBackground

	UndefinedSessionDate
	InvalidSessionReward
	DuplicateSessionReward
	BadSessionRewardPosition

	InvalidUpgrade
	UndefinedUpgradeName
	UndefinedUpgradeCost
	InvalidUpgradeMark
	InvalidUpgradeCost
	DuplicateUpgradeCost
	BadUpgradeCostPosition
	MismatchMarkCost

	InvalidCharacteristicFormat
	InvalidCharacteristicValue
	UndefinedCharacteristic
	DuplicateCharacteristic
	MissingCharacteristic

	InvalidGaugeValue

	DuplicateTalent

	UndefinedTypeCost
	UndefinedMatchCost
	UndefinedTierCost
)

// errorMsgs contains the messages associated to the error codes.
var errorMsgs = map[ErrorCode]string{

	InvalidCharacterSheet: `the character sheet requires at least a header block and a characteristic block`,

	InvalidHeaderLine:        `line %d: the header line format is invalid`,
	EmptyHeaderKey:           `line %d: the header line key is empty`,
	EmptyHeaderValue:         `line %d: the header line value is empty`,
	DuplicateHeaderLine:      `line %d: the header line is already set`,
	InvalidBackgroundOptions: `line %d: the background options are incorrect`,
	UndefinedBackground:      `line %d: the %s %s is not defined by the universe`,

	UndefinedSessionDate:     `line %d: the session date is not defined`,
	InvalidSessionReward:     `line %d: the session reward is invalid`,
	DuplicateSessionReward:   `line %d: the session reward is already set`,
	BadSessionRewardPosition: `line %d: bad session reward position`,

	InvalidUpgrade:         `line %d: the upgrade format is invalid`,
	UndefinedUpgradeName:   `line %d: the upgrade name is not defined`,
	UndefinedUpgradeCost:   `line %d: the upgrade cost is not defined`,
	InvalidUpgradeMark:     `line %d: the upgrade mark is invalid`,
	InvalidUpgradeCost:     `line %d: the upgrade cost is invalid`,
	DuplicateUpgradeCost:   `line %d: the upgrade cost is already set`,
	BadUpgradeCostPosition: `line %d: bad upgrade cost position`,
	MismatchMarkCost:       `line %d: upgrade with mark "-" expects no cost`,

	InvalidCharacteristicFormat: `line %d: the characteristic format is invalid`,
	InvalidCharacteristicValue:  `line %d: the characteristic value is invalid`,
	UndefinedCharacteristic:     `line %d: the characteristic is not defined`,
	DuplicateCharacteristic:     `line %d: the characteristic is already set`,
	MissingCharacteristic:       `the characteristic %s is not defined for the character`,

	InvalidGaugeValue: `line %d: the gauge value is invalid`,

	DuplicateTalent: `line %d: the talent is already set`,

	UndefinedTypeCost:  `undefined cost for type %s`,
	UndefinedMatchCost: `undefined cost for type %s with %d matching aptitudes`,
	UndefinedTierCost:  `undefined cost for type %s with %d matching aptitudes on tier %d`,
}

// Error is an error holding a code and variadic printable data.
type Error struct {
	Code ErrorCode
	vars []interface{}
}

// NewError build a new error from an error code.
func NewError(code ErrorCode, v ...interface{}) Error {
	return Error{
		Code: code,
		vars: v,
	}
}

// Error implements the Error interface.
func (e Error) Error() string {
	msg, found := errorMsgs[e.Code]
	if !found {
		panic(fmt.Sprintf("undefined error message for code %d", e.Code))
	}
	return fmt.Sprintf(msg, e.vars...)
}
