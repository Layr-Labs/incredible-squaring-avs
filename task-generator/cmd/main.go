package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	taskgenerator "github.com/Layr-Labs/incredible-squaring-avs/task-generator"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {

	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "credible-squaring-task-generator"
	app.Usage = "Credible Squaring Task Generator"
	app.Description = "Service that generates tasks and sends them to Task Manager."

	app.Action = taskGeneratorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func taskGeneratorMain(ctx *cli.Context) error {

	log.Println("Initializing Task Generator...")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	fmt.Println("Config:", string(configJson))

	taskGen, err := taskgenerator.BuildTaskGenerator(config)
	if err != nil {
		return err
	}

	err = taskGen.Start(context.Background())
	if err != nil {
		return err
	}

	return nil

}
