// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "math/big"

var (
	// TargetGasLimit is the artificial target
	TargetGasLimit = GenesisGasLimit
)

const (
	//GasLimitBoundDivisor uint64 = 1024 // The bound divisor of the gas limit, used in update calculations.
	GasLimitBoundDivisor uint64 = 100  // The bound divisor of the gas limit, used in update calculations.
	MinGasLimit          uint64 = 5000 // Minimum the gas limit may ever be.
	//GenesisGasLimit      uint64 = 4712388 // Gas limit of the Genesis block.
	GenesisGasLimit uint64 = 10000000 // Gas limit of the Genesis block.

	MaximumExtraDataSize  uint64 = 32    // Maximum size extra data may be after Genesis.
	ExpByteGas            uint64 = 10    // Times ceil(log256(exponent)) for the EXP instruction.
	SloadGas              uint64 = 50    // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	CallValueTransferGas  uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	CallNewAccountGas     uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	TxGas                 uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxGasContractCreation uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroGas         uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv          uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	SstoreSetGas          uint64 = 20000 // Once per SLOAD operation.
	LogDataGas            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend           uint64 = 2300  // Free gas given at beginning of call.

	Sha3Gas     uint64 = 30 // Once per SHA3 operation.
	Sha3WordGas uint64 = 6  // Once per word of the SHA3 operation's data.

	SstoreResetGas  uint64 = 5000  // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClearGas  uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundGas uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.

	NetSstoreNoopGas  uint64 = 200   // Once per SSTORE operation if the value doesn't change.
	NetSstoreInitGas  uint64 = 20000 // Once per SSTORE operation from clean zero.
	NetSstoreCleanGas uint64 = 5000  // Once per SSTORE operation from clean non-zero.
	NetSstoreDirtyGas uint64 = 200   // Once per SSTORE operation from dirty.

	NetSstoreClearRefund      uint64 = 15000 // Once per SSTORE operation for clearing an originally existing storage slot
	NetSstoreResetRefund      uint64 = 4800  // Once per SSTORE operation for resetting to the original non-zero value
	NetSstoreResetClearRefund uint64 = 19800 // Once per SSTORE operation for resetting to the original zero value

	JumpdestGas      uint64 = 1     // Refunded gas, once per SSTORE operation if the zeroness changes to zero.
	EpochDuration    uint64 = 30000 // Duration between proof-of-work epochs.
	CallGas          uint64 = 40    // Once per CALL operation & message call transaction.
	CreateDataGas    uint64 = 200   //
	CallCreateDepth  uint64 = 1024  // Maximum depth of call/create stack.
	ExpGas           uint64 = 10    // Once per EXP instruction
	LogGas           uint64 = 375   // Per LOG* operation.
	CopyGas          uint64 = 3     //
	StackLimit       uint64 = 1024  // Maximum size of VM stack allowed.
	TierStepGas      uint64 = 0     // Once per operation, for a selection of them.
	LogTopicGas      uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	CreateGas        uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	Create2Gas       uint64 = 32000 // Once per CREATE2 operation
	SuicideRefundGas uint64 = 24000 // Refunded following a suicide operation.
	MemoryGas        uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.
	TxDataNonZeroGas uint64 = 68    // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.

	MaxCodeSize = 24576 // Maximum bytecode to permit for a contract

	// Precompiled contract gas prices

	EcrecoverGas                    uint64 = 3000   // Elliptic curve sender recovery gas price
	Sha256BaseGas                   uint64 = 60     // Base price for a SHA256 operation
	Sha256PerWordGas                uint64 = 12     // Per-word price for a SHA256 operation
	Ripemd160BaseGas                uint64 = 600    // Base price for a RIPEMD160 operation
	Ripemd160PerWordGas             uint64 = 120    // Per-word price for a RIPEMD160 operation
	IdentityBaseGas                 uint64 = 15     // Base price for a data copy operation
	IdentityPerWordGas              uint64 = 3      // Per-work price for a data copy operation
	ModExpQuadCoeffDiv              uint64 = 20     // Divisor for the quadratic particle of the big int modular exponentiation
	Bn256AddGas                     uint64 = 500    // Gas needed for an elliptic curve addition
	Bn256AddGasIstanbul             uint64 = 150    // Gas needed for an elliptic curve addition
	Bn256ScalarMulGas               uint64 = 40000  // Gas needed for an elliptic curve scalar multiplication
	Bn256ScalarMulGasIstanbul       uint64 = 6000   // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseGas             uint64 = 100000 // Base price for an elliptic curve pairing check
	Bn256PairingBaseGasIstanbul     uint64 = 45000  // Base price for an elliptic curve pairing check
	Bn256PairingPerPointGas         uint64 = 80000  // Per-point price for an elliptic curve pairing check
	Bn256PairingPerPointGasIstanbul uint64 = 34000  // Per-point price for an elliptic curve pairing check
)

var (
	DifficultyBoundDivisor = big.NewInt(2)       // The bound divisor of the difficulty, used in the update calculations.
	GenesisDifficulty      = big.NewInt(6000000) // Difficulty of the Genesis block.
	MinimumDifficulty      = big.NewInt(2000000) // The minimum that the difficulty may ever be.
	MinimumFruitDifficulty = big.NewInt(2000)
	DurationLimit          = big.NewInt(600) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.

	DifficultyPeriod = big.NewInt(144) //

	FruitBlockRatio = big.NewInt(600) // difficulty ratio between fruit and snail block

	FruitFreshness = big.NewInt(17) // fruit freshness

	MinimumFruits int = 60
	MaximumFruits int = 600

	MinTimeGap = big.NewInt(359)
)

var (
	SnailConfirmInterval = big.NewInt(12)

	SnailRewardInterval = big.NewInt(14)

	SnailMaximumRewardInterval = big.NewInt(20)

	FastToFruitSpace = big.NewInt(1500)

	ElectionPeriodNumber     = big.NewInt(180) // snail block period number
	ElectionSwitchoverNumber = big.NewInt(9600)

	ElectionFruitsThreshold uint64 = 100 // fruit size threshold for committee election

	MaximumCommitteeNumber  = big.NewInt(50)
	ProposalCommitteeNumber = 20
	MinimumCommitteeNumber  = 7
)

var (
	CountInEpoch                      = 20
	MaxRedeemHeight            uint64 = 250000 // about 15 days
	NewEpochLength             uint64 = 25000  // about 1.5 days
	ElectionPoint              uint64 = 200
	FirstNewEpochID            uint64 = 1
	DposForkPoint              uint64 = 0
	EffectOfStakingModify      uint64 = 0      // for mainnet,testnet,devnet
	ElectionMinLimitForStaking        = new(big.Int).Mul(big.NewInt(20000), big.NewInt(1e18))
)
