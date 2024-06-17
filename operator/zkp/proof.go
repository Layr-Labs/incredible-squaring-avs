package main

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// AssertLessThan defines a simple circuit
type AssertLessThanCircuit struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints
// bound >= x
func (circuit *AssertLessThanCircuit) Define(api frontend.API) error {
	api.AssertIsLessOrEqual(circuit.X, circuit.Y)
	return nil
}

func createProof(value int, bound int) (groth16.Proof, groth16.VerifyingKey, witness.Witness, error) {
	// compiles our circuit into a R1CS
	var circuit AssertLessThanCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	pk, vk, _ := groth16.Setup(ccs)
	assignment := AssertLessThanCircuit{X: value, Y: bound}
	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	proof, _ := groth16.Prove(ccs, pk, witness)
	return proof, vk, witness, nil
}

func verify(proof groth16.Proof, vk groth16.VerifyingKey, witness witness.Witness) bool {
	result := groth16.Verify(proof, vk, witness)
	return result != nil
}
