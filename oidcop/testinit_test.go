package oidcop

import (
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
)

func setupTest() {
	utils.SetIsTest()

	db.SetIsTest()

	types.SetIsTest("oidcop")

	schema.SetIsTest()
}

func teardownTest() {
	schema.UnsetIsTest()

	types.UnsetIsTest("oidcop")

	db.UnsetIsTest()

	utils.UnsetIsTest()
}
