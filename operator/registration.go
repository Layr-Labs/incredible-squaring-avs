package operator

// OUTDATED
// This file contains cli functions for registering an operator with the AVS and printing status
// However, all of this functionality has been moved to the plugin/ package
// we are just waiting for eigenlayer-cli to be open sourced so we can completely get rid of this registration functionality in the operator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
)

func (o *Operator) registerOperatorOnStartup(blsPubkeyCompendiumAddr common.Address) {
	err := o.RegisterOperatorWithEigenlayer()
	if err != nil {
		// This error might only be that the operator was already registered with eigenlayer, so we don't want to fatal
		o.logger.Error("Error registering operator with eigenlayer", "err", err)
	} else {
		o.logger.Infof("Registered operator with eigenlayer")
	}

	err = o.RegisterBLSPublicKey()
	if err != nil {
		// This error might only be that the operator has already registered its bls pubkeys, so we don't want to fatal
		o.logger.Error("Error registering BLS public key with eigenlayer", "err", err)
	} else {
		o.logger.Infof("Registered BLS public key with eigenlayer")
	}

	// TODO(samlaf): these shouldn't be hardcoded
	mockTokenStrategyAddr := common.HexToAddress("0x7a2088a1bFc9d81c55368AE168C2C02570cB814F")
	amount := big.NewInt(1000)
	err = o.DepositIntoStrategy(mockTokenStrategyAddr, amount)
	if err != nil {
		o.logger.Fatal("Error depositing into strategy", "err", err)
	}
	o.logger.Infof("Deposited %s into strategy %s", amount, mockTokenStrategyAddr)

	err = o.RegisterOperatorWithAvs()
	if err != nil {
		o.logger.Fatal("Error registering operator with avs", "err", err)
	}
	o.logger.Infof("Registered operator with avs")
}

func (o *Operator) optOperatorIntoSlashing() error {
	_, err := o.eigenlayerWriter.OptOperatorIntoSlashing(context.Background(), o.credibleSquaringServiceManagerAddr)
	if err != nil {
		o.logger.Errorf("Error opting operator into slashing")
		return err
	}
	return nil
}

func (o *Operator) RegisterOperatorWithEigenlayer() error {
	op := eigenSdkTypes.Operator{
		Address:                 o.operatorAddr.String(),
		EarningsReceiverAddress: o.operatorAddr.String(),
	}
	_, err := o.eigenlayerWriter.RegisterAsOperator(context.Background(), op)
	if err != nil {
		o.logger.Errorf("Error registering operator with eigenlayer")
		return err
	}
	return nil
}

func (o *Operator) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {
	_, tokenAddr, err := o.eigenlayerReader.GetStrategyAndUnderlyingToken(context.Background(), strategyAddr)
	if err != nil {
		o.logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}
	contractErc20Mock, err := o.avsReader.GetErc20Mock(context.Background(), tokenAddr)
	if err != nil {
		o.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}
	txOpts := o.avsWriter.Signer.GetTxOpts()
	tx, err := contractErc20Mock.Mint(txOpts, o.operatorAddr, amount)
	if err != nil {
		o.logger.Errorf("Error assembling Mint tx")
		return err
	}
	o.ethClient.WaitForTransactionReceipt(context.Background(), tx.Hash())

	_, err = o.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
	if err != nil {
		o.logger.Errorf("Error depositing into strategy")
		return err
	}
	return nil
}

// TODO(samlaf): this should eventually be moved to eigensdk, since pubkey compendium should be a shared contract
// between all AVS teams. However it currently is not in that state (not included as part of canonical eigenlayer deployment)
// so for now we just deploy it as part of avs registries and make the call here.
func (o *Operator) RegisterBLSPublicKey() error {
	o.logger.Debugf("registering BLS public keys G1(%s) and G2(%s) with EigenLayer", o.blsKeypair.GetPubKeyG1(), o.blsKeypair.GetPubKeyG2())

	_, err := o.eigenlayerWriter.RegisterBLSPublicKey(context.Background(), o.blsKeypair, eigenSdkTypes.Operator{Address: o.operatorAddr.String()})
	if err != nil {
		return fmt.Errorf("failed to register bls pubkey with compendium: %v", err)
	}

	o.logger.Debugf("registered BLS public keys G1(%s) and G2(%s) with EigenLayer", o.blsKeypair.GetPubKeyG1(), o.blsKeypair.GetPubKeyG2())
	return nil
}

// Registration specific functions
func (o *Operator) RegisterOperatorWithAvs() error {

	// 1. opt operator into getting slashed by credible squaring service manager
	err := o.optOperatorIntoSlashing()
	if err != nil {
		return err
	}

	// 2. register operator with avs coordinator contract

	quorumNumbers := []byte{0}
	socket := "Not Needed"

	pubkey := pubKeyG1ToBN254G1Point(o.blsKeypair.GetPubKeyG1())
	g1Point := regcoord.BN254G1Point{
		X: pubkey.X,
		Y: pubkey.Y,
	}

	_, err = o.avsWriter.RegisterOperatorWithAVSRegistryCoordinator(context.Background(), quorumNumbers, g1Point, socket)
	if err != nil {
		o.logger.Errorf("Unable to register operator with avs registry coordinator")
		return err
	}
	o.logger.Infof("Registered operator with avs registry coordinator.")

	return nil
}

// PRINTING STATUS OF OPERATOR: 1
// operator address: 0xa0ee7a142d267c1f36714e4a8f75612f20a79720
// dummy token balance: 0
// delegated shares in dummyTokenStrat: 200
// operator pubkey hash in AVS pubkey compendium (0 if not registered): 0x4b7b8243d970ff1c90a7c775c008baad825893ec6e806dfa5d3663dc093ed17f
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
	OptedIntoSlashingByAvs bool
	RegisteredWithAvs      bool
	OperatorId             string
	Frozen                 bool
}

func (o *Operator) PrintOperatorStatus() error {
	fmt.Println("Printing operator status")
	pubkeyhash, err := o.eigenlayerReader.GetOperatorPubkeyHash(context.Background(), eigenSdkTypes.Operator{Address: o.operatorAddr.String()})
	if err != nil {
		return err
	}
	pubkeysRegistered := pubkeyhash != [32]byte{}
	serviceManagerCanSlashOperatorUntilBlock, err := o.eigenlayerReader.ServiceManagerCanSlashOperatorUntilBlock(
		context.Background(), o.operatorAddr, o.credibleSquaringServiceManagerAddr,
	)
	if err != nil {
		return err
	}
	curBlockNumber, err := o.ethClient.BlockNumber(context.Background())
	if err != nil {
		return err
	}
	optedIntoSlashingByAvs := curBlockNumber < uint64(serviceManagerCanSlashOperatorUntilBlock)
	registeredWithAvs := o.operatorId != [32]byte{}
	isFrozen, err := o.eigenlayerReader.OperatorIsFrozen(context.Background(), o.operatorAddr)
	if err != nil {
		return err
	}
	operatorStatus := OperatorStatus{
		EcdsaAddress:           o.operatorAddr.String(),
		PubkeysRegistered:      pubkeysRegistered,
		G1Pubkey:               o.blsKeypair.GetPubKeyG1().String(),
		G2Pubkey:               o.blsKeypair.GetPubKeyG2().String(),
		OptedIntoSlashingByAvs: optedIntoSlashingByAvs,
		RegisteredWithAvs:      registeredWithAvs,
		OperatorId:             hex.EncodeToString(o.operatorId[:]),
		Frozen:                 isFrozen,
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
