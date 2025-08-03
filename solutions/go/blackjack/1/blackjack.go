package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	case card == "ace":
		return 11
	default:
		return 0

	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {

	case card1 == "ace" && card2 == "ace":
		return "P"
	case (handValue) == 21 && !((dealerCard == "ace") || (dealerCard == "jack") || (dealerCard == "queen") || (dealerCard == "king")):
		return "W"
	case (handValue) == 21 && ((dealerCard == "ace") || (dealerCard == "jack") || (dealerCard == "queen") || (dealerCard == "king")):
		return "S"
	case (handValue) > 16 && (handValue) < 21:
		return "S"
	//case (ParseCard(card1)+ParseCard(card2)) > 11 && (ParseCard(card1)+ParseCard(card2)) < 17:
	//return "S"
	case (handValue) > 11 && (handValue) < 17 && !(dealerValue >= 7):
		return "S"
	default:
		return "H"
	}
}
