package operator

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"

	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/metrics"
	"github.com/Layr-Labs/incredible-squaring-avs/tests"
)

// Name starts with Integration test because we don't want it to run with go test ./...
// since this starts a chain and takes longer to run
// TODO(samlaf): buggy test, fix it
func IntegrationTestOperatorRegistration(t *testing.T) {
	anvilCmd := tests.StartAnvilChainAndDeployContracts()
	defer anvilCmd.Process.Kill()
	operator, err := createMockOperator()
	assert.Nil(t, err)
	err = operator.RegisterOperatorWithEigenlayer()
	assert.Nil(t, err)
}

func createMockOperator() (*Operator, error) {
	logger := sdklogging.NewNoopLogger()
	reg := prometheus.NewRegistry()
	noopMetrics := metrics.NewNoopMetrics()

	avsEcdsaPrivateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ecdsaPrivateKey, err := ecdsa.NewPrivateKeyFromHex(avsEcdsaPrivateKeyHex)
	if err != nil {
		return nil, err
	}

	operator := &Operator{
		logger:             logger,
		avsEcdsaPrivateKey: ecdsaPrivateKey,
		operatorId:         ecdsa.PrivateKeyToOperatorId(ecdsaPrivateKey),
		operatorAddr:       ecdsa.PrivateKeyToOperatorId(ecdsaPrivateKey), // we'll use the same.. think this is fine..?
		metricsReg:         reg,
		metrics:            noopMetrics,
		newTaskCreatedChan: make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated),
	}
	return operator, nil
}
