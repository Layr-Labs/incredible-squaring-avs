package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// JSON processing: reading and extracting values
	jsonFile, err := os.Open("contracts/script/output/31337/credible_squaring_avs_deployment_output.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	var jsonData map[string]map[string]string
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	if err := json.Unmarshal(byteValue, &jsonData); err != nil {
		fmt.Println("Error unmarshalling JSON data:", err)
		return
	}

	// YAML processing: loading, updating, and writing back
	yamlFile, err := os.Open("config-files/operator-docker-compose.anvil.yaml")
	yamlFile2, err2 := os.Open("config-files/operator.anvil.yaml")
	if err != nil {
		fmt.Println("Error opening YAML file:", err)
		return
	}
	if err2 != nil {
		fmt.Println("Error opening YAML file2:", err)
		return
	}
	defer yamlFile.Close()
	defer yamlFile2.Close()

	// For the Docker
	var yamlData yaml.Node
	decoder := yaml.NewDecoder(yamlFile)
	if err := decoder.Decode(&yamlData); err != nil {
		fmt.Println("Error decoding YAML file:", err)
		return
	}

	// Operator Yaml
	var yamlData2 yaml.Node
	decoder2 := yaml.NewDecoder(yamlFile2)
	if err2 := decoder2.Decode(&yamlData2); err2 != nil {
		fmt.Println("Error decoding YAML file:", err2)
		return
	}

	// Update Docker data based on JSON values
	updateYAMLScalar(&yamlData, "avs_registry_coordinator_address", jsonData["addresses"]["registryCoordinator"])
	updateYAMLScalar(&yamlData, "operator_state_retriever_address", jsonData["addresses"]["operatorStateRetriever"])
	updateYAMLScalar(&yamlData, "token_strategy_addr", jsonData["addresses"]["erc20MockStrategy"])

	// Update Operator YAML
	updateYAMLScalar(&yamlData2, "avs_registry_coordinator_address", jsonData["addresses"]["registryCoordinator"])
	updateYAMLScalar(&yamlData2, "operator_state_retriever_address", jsonData["addresses"]["operatorStateRetriever"])
	updateYAMLScalar(&yamlData2, "token_strategy_addr", jsonData["addresses"]["erc20MockStrategy"])

	// Writing back to docker file, preserving comments
	outputFile, err := os.Create("config-files/operator-docker-compose.anvil.yaml")
	if err != nil {
		fmt.Println("Error reopening YAML file for writing:", err)
		return
	}
	defer outputFile.Close()

	// Writing back to the operator file
	outputFile2, err2 := os.Create("config-files/operator.anvil.yaml")
	if err2 != nil {
		fmt.Println("Error reopening YAML file for writing:", err)
		return
	}
	defer outputFile2.Close()

	// Docker
	encoder := yaml.NewEncoder(outputFile)
	encoder.SetIndent(2)
	if err := encoder.Encode(&yamlData); err != nil {
		fmt.Println("Error encoding YAML data:", err)
		return
	}
	encoder.Close()

	// Operator
	encoder2 := yaml.NewEncoder(outputFile2)
	encoder2.SetIndent(2)
	if err2 := encoder2.Encode(&yamlData2); err2 != nil {
		fmt.Println("Error encoding YAML data:", err2)
		return
	}
	encoder2.Close()
	fmt.Println("YAML file updated successfully with comments preserved!")
}

// updateYAMLScalar updates a scalar value directly within a YAML node tree
func updateYAMLScalar(node *yaml.Node, targetKey string, newValue string) {
	if node.Kind == yaml.DocumentNode {
		node = node.Content[0] // Assuming the root is a document node, we step into it
	}

	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 { // Nodes are stored as pairs in MappingNode
			keyNode := node.Content[i]
			valueNode := node.Content[i+1]
			if keyNode.Value == targetKey {
				valueNode.Value = newValue
				return
			}
		}
	}

	fmt.Printf("Key '%s' not found for updating\n", targetKey)
}
