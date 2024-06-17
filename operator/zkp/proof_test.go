package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	threshold := 500
	// Create a valid proof
	proofValid, vk, witness, _ := createProof(35, threshold)
	assert.True(t, verify(proofValid, vk, witness), "Expected proofValid to verify successfully")

	// Create an invalid proof
	proofInvalid, vkInvalid, witnessInvalid, _ := createProof(1000, threshold)
	assert.False(t, verify(proofInvalid, vkInvalid, witnessInvalid), "Expected proofInvalid to fail verification")
}
