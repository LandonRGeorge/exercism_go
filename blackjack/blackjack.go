package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
    		return 11
        case "two":
    		return 2
        case "three":
    		return 3
		case "four":
    		return 4
        case "five":
    		return 5
		case "six":
    		return 6
		case "seven":
    		return 7
        case "eight":
    		return 8
        case "nine":
    		return 9
        case "ten":
    		return 10
        case "jack":
    		return 10
        case "queen":
    		return 10
        case "king":
    		return 10
        default:
    		return 0   
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cVal1 := ParseCard(card1)
    cVal2 := ParseCard(card2)
    cVal := cVal1 + cVal2
    dVal1 := ParseCard(dealerCard)
    switch {
        case card1 == "ace" && card2 == "ace":
        	return "P"
        case cVal == 21 && !(dVal1 == 10 || dealerCard == "ace"):
        	return "W"
		case cVal >= 17:
        	return "S"
        case cVal >= 12 && !(dVal1 >= 7):
        	return "S"
        default:
    		return "H"
        
    }    
}
