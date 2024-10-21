package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten":
		return 10
	case card == "jack":
		return 10
	case card == "queen":
		return 10
	case card == "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerSum := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	if playerSum == 21 {
		if dealerCard == "ace" || dealerCard == "king" || dealerCard == "queen" || dealerCard == "jack" || dealerCard == "ten" {
			return "S"
		}
		return "W"
	}
	if playerSum >= 17 && playerSum <= 20 {
		return "S"
	}
	if playerSum >= 12 && playerSum <= 16 {
		if dealerValue >= 7 {
			return "H"
		}
		return "S"
	}
	if playerSum <= 11 {
		return "H"
	}
	return ""
}
