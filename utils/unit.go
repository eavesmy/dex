/*
# File Name: ./utils/unit.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: Thu May  5 12:00:02 2022
*/

package utils

import (
	"math"
	"math/big"
)

func UnitToPow(unit string) int {

	str := ""

	switch unit {
	case "wei":
		str = "1"
	case "kwei", "Kwei", "babbage", "femtoether":
		str = "1000"
	case "mwei", "Mwei", "lovelace":
		str = "1000000"
	case "picoether":
		str = "1000000"
	case "gwei":
		str = "1000000000"
	case "Gwei":
		str = "1000000000"
	case "shannon":
		str = "1000000000"
	case "nanoether":
		str = "1000000000"
	case "nano":
		str = "1000000000"
	case "szabo":
		str = "1000000000000"
	case "microether":
		str = "1000000000000"
	case "micro":
		str = "1000000000000"
	case "finney":
		str = "1000000000000000"
	case "milliether":
		str = "1000000000000000"
	case "milli":
		str = "1000000000000000"
	case "ether":
		str = "1000000000000000000"
	case "kether":
		str = "1000000000000000000000"
	case "grand":
		str = "1000000000000000000000"
	case "mether":
		str = "1000000000000000000000000"
	case "gether":
		str = "1000000000000000000000000000"
	case "tether":
		str = "1000000000000000000000000000000"
	}

	return len([]rune(str))
}

func AmountByDecimal(amount *big.Float, decimal int) *big.Float {
	return new(big.Float).Quo(amount, big.NewFloat(math.Pow10(decimal)))
}
