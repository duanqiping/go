package main

import (
	"familyaccount/utils"
)

func main()  {

	account := utils.NewAccount()
	account.Verify()

	familyAccount := utils.NewFamilyAccount()
	familyAccount.MainMenu()
}
