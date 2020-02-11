package main

import (
	"fmt"
	"github.com/truechain/truechain-engineering-code/cmd/utils"
	"github.com/truechain/truechain-engineering-code/core/types"
	"gopkg.in/urfave/cli.v1"
	"math/big"
)

var AppendCommand = cli.Command{
	Name:   "append",
	Usage:  "Append validator deposit staking count",
	Action: utils.MigrateFlags(AppendImpawn),
	Flags:  ImpawnFlags,
}

func AppendImpawn(ctx *cli.Context) error {
	loadPrivate(ctx)

	conn, url := dialConn(ctx)
	printBaseInfo(conn, url)

	value := trueToWei(ctx, false)

	input := packInput("append")
	txHash := sendContractTransaction(conn, from, types.StakingAddress, value, priKey, input)

	getResult(conn, txHash)

	return nil
}

var UpdateFeeCommand = cli.Command{
	Name:   "updatefee",
	Usage:  "Update delegate fee will take effect in next epoch",
	Action: utils.MigrateFlags(UpdateFeeImpawn),
	Flags:  ImpawnFlags,
}

func UpdateFeeImpawn(ctx *cli.Context) error {
	loadPrivate(ctx)

	conn, url := dialConn(ctx)
	printBaseInfo(conn, url)

	fee = ctx.GlobalUint64(FeeFlag.Name)
	checkFee(new(big.Int).SetUint64(fee))
	fmt.Println("Fee", fee)

	input := packInput("setFee", new(big.Int).SetUint64(fee))

	txHash := sendContractTransaction(conn, from, types.StakingAddress, new(big.Int).SetInt64(0), priKey, input)

	getResult(conn, txHash)
	return nil
}

var cancelCommand = cli.Command{
	Name:   "cancel",
	Usage:  "Call this staking will cancelled at the next epoch",
	Action: utils.MigrateFlags(cancelImpawn),
	Flags:  ImpawnFlags,
}

func cancelImpawn(ctx *cli.Context) error {
	loadPrivate(ctx)
	conn, url := dialConn(ctx)
	printBaseInfo(conn, url)

	value := trueToWei(ctx, false)

	input := packInput("cancel", value)
	txHash := sendContractTransaction(conn, from, types.StakingAddress, new(big.Int).SetInt64(0), priKey, input)

	getResult(conn, txHash)
	return nil
}

var withdrawCommand = cli.Command{
	Name:   "withdraw",
	Usage:  "Call this will instant receive your deposit money",
	Action: utils.MigrateFlags(withdrawImpawn),
	Flags:  ImpawnFlags,
}

func withdrawImpawn(ctx *cli.Context) error {
	loadPrivate(ctx)
	conn, url := dialConn(ctx)
	printBaseInfo(conn, url)
	PrintBalance(conn, from)

	value := trueToWei(ctx, false)

	input := packInput("withdraw", value)

	txHash := sendContractTransaction(conn, from, types.StakingAddress, new(big.Int).SetInt64(0), priKey, input)

	getResult(conn, txHash)
	PrintBalance(conn, from)
	return nil
}

var queryStakingCommand = cli.Command{
	Name:   "querystaking",
	Usage:  "Query staking info, can cancel info and can withdraw info",
	Action: utils.MigrateFlags(queryStakingImpawn),
	Flags:  ImpawnFlags,
}

func queryStakingImpawn(ctx *cli.Context) error {
	loadPrivate(ctx)
	conn, url := dialConn(ctx)
	printBaseInfo(conn, url)

	queryStakingInfo(conn)
	return nil
}
