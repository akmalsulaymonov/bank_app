package card

import "github.com/akmalsulaymonov/bank_app/pkg/bank/types"

func Withdraw(card *types.Card, amount types.Money) {
	const withdrawLimit = 20_000_00
	if card.Active == true && card.Balance >= amount && amount > 0 && amount <= withdrawLimit {
		card.Balance -= amount
	}
}
