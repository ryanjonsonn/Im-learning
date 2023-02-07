package app

import (
	appparams "github.com/Pylons-tech/pylons/app/params"

	"github.com/cosmos/cosmos-sdk/std"
)
func NewSpamMigitationAnteDecorator(pylonsmodulekeeper PylonsKeeper) AnteSpamMigitationDecorator {
	return AnteSpamMigitationDecorator{
		pk: pylonsmodulekeeper,
	}
}
