package operator

// OUTDATED
// This file contains cli functions for registering an operator with the AVS and printing status
// However, all of this functionality has been moved to the plugin/ package
// we are just waiting for eigenlayer-cli to be open sourced so we can completely get rid of this registration
// functionality in the operator

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
)

func (o *Operator) registerOperatorOnStartup(
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	mockTokenStrategyAddr common.Address,
	registryAddr common.Address,
	avsAddress common.Address,
	operatorSetsIds []uint32,
	waitForReceipt bool,
	blsKeyPair bls.KeyPair,
	socket string,
) {
	err := o.RegisterOperatorWithEigenlayer()
	if err != nil {
		// This error might only be that the operator was already registered with eigenlayer, so we don't want to fatal
		o.logger.Error("Error registering operator with eigenlayer", "err", err)
	} else {
		o.logger.Info("Registered operator with eigenlayer")
	}

	// TODO(samlaf): shouldn't hardcode number here
	// Use SetString for large numbers
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // Base 10

	err = o.DepositIntoStrategy(mockTokenStrategyAddr, amount)
	if err != nil {
		o.logger.Fatal("Error depositing into strategy", "err", err)
	}
	o.logger.Infof("Deposited %s into strategy %s", amount, mockTokenStrategyAddr)

	err = o.RegisterForOperatorSets(
		registryAddr,
		avsAddress,
		operatorSetsIds,
		waitForReceipt,
		blsKeyPair,
		socket,
		operatorEcdsaPrivateKey,
	)
	if err != nil {
		o.logger.Fatal("Error registering operator with avs", "err", err)
	}
	o.logger.Info("Registered operator with avs")
}

func (o *Operator) RegisterOperatorWithEigenlayer() error {
	op := eigenSdkTypes.Operator{
		Address:                   o.operatorAddr.String(),
		DelegationApproverAddress: o.operatorAddr.String(),
	}
	_, err := o.eigenlayerWriter.RegisterAsOperator(context.Background(), op, true)
	if err != nil {
		o.logger.Error("Error registering operator with eigenlayer", "err", err)
		return err
	}
	return nil
}

func (o *Operator) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {

	_, tokenAddr, err := o.eigenlayerReader.GetStrategyAndUnderlyingToken(context.TODO(), strategyAddr)
	if err != nil {
		o.logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}
	o.logger.Info(tokenAddr.String())
	contractErc20Mock, err := o.avsReader.GetErc20Mock(context.Background(), tokenAddr)
	if err != nil {
		o.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}
	txOpts, err := o.avsWriter.TxMgr.GetNoSendTxOpts()
	if err != nil {
		o.logger.Errorf("Error in GetNoSendTxOpts")
		return err
	}
	tx, err := contractErc20Mock.Mint(txOpts, o.operatorAddr, amount)
	if err != nil {
		o.logger.Errorf("Error assembling Mint tx")
		return err
	}
	_, err = o.avsWriter.TxMgr.Send(context.Background(), tx, true)
	if err != nil {
		o.logger.Errorf("Error submitting Mint tx")
		return err
	}

	_, err = o.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount, true)
	if err != nil {
		o.logger.Errorf("Error depositing into strategy", "err", err)
		return err
	}
	return nil
}

// Registration specific functions
func (o *Operator) RegisterForOperatorSets(
	registryAddr common.Address,
	avsAddress common.Address,
	operatorSetIds []uint32,
	waitForReceipt bool,
	blsKeyPair bls.KeyPair,
	socket string,
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
) error {
	operatorAddress := crypto.PubkeyToAddress(operatorEcdsaKeyPair.PublicKey)

	registrationRequest := elcontracts.RegistrationRequest{
		OperatorAddress: operatorAddress,
		AVSAddress:      avsAddress,
		OperatorSetIds:  operatorSetIds,
		WaitForReceipt:  waitForReceipt,
		BlsKeyPair:      &blsKeyPair,
		Socket:          socket,
	}

	_, err := o.eigenlayerWriter.RegisterForOperatorSets(
		context.Background(),
		registryAddr,
		registrationRequest,
	)

	if err != nil {
		o.logger.Errorf("Unable to register operator with the operator set")
		return err
	}
	o.logger.Info("Registered operator with operator set")

	return nil
}

// PRINTING STATUS OF OPERATOR: 1
// operator address: 0xa0ee7a142d267c1f36714e4a8f75612f20a79720
// dummy token balance: 0
// delegated shares in dummyTokenStrat: 200
// operator pubkey hash in AVS pubkey compendium (0 if not registered):
// 0x4b7b8243d970ff1c90a7c775c008baad825893ec6e806dfa5d3663dc093ed17f
// operator is opted in to eigenlayer: true
// operator is opted in to playgroundAVS (aka can be slashed): true
// operator status in AVS registry: REGISTERED
//
//	operatorId: 0x4b7b8243d970ff1c90a7c775c008baad825893ec6e806dfa5d3663dc093ed17f
//	middlewareTimesLen (# of stake updates): 0
//
// operator is frozen: false
type OperatorStatus struct {
	EcdsaAddress string
	// pubkey compendium related
	PubkeysRegistered bool
	G1Pubkey          string
	G2Pubkey          string
	// avs related
	RegisteredWithAvs bool
	OperatorId        string
}

func (o *Operator) PrintOperatorStatus() error {
	fmt.Println("Printing operator status")
	operatorId, err := o.avsReader.GetOperatorId(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		return err
	}
	pubkeysRegistered := operatorId != [32]byte{}
	registeredWithAvs := o.operatorId != [32]byte{}
	operatorStatus := OperatorStatus{
		EcdsaAddress:      o.operatorAddr.String(),
		PubkeysRegistered: pubkeysRegistered,
		G1Pubkey:          o.blsKeypair.GetPubKeyG1().String(),
		G2Pubkey:          o.blsKeypair.GetPubKeyG2().String(),
		RegisteredWithAvs: registeredWithAvs,
		OperatorId:        hex.EncodeToString(o.operatorId[:]),
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(operatorStatusJson))
	return nil
}

func pubKeyG1ToBN254G1Point(p *bls.G1Point) regcoord.BN254G1Point {
	return regcoord.BN254G1Point{
		X: p.X.BigInt(new(big.Int)),
		Y: p.Y.BigInt(new(big.Int)),
	}
}
