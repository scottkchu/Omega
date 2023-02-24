package common

import (
	"fmt"
	"os"

	cohere "github.com/cohere-ai/cohere-go"
)

// var (
// 	cohereToken string
// )

// func init() {
// 	configs, err := GetConfigs()
// 	if err != nil {
// 		fmt.Println("Error getting configs: ", err)
// 		return
// 	}

// 	cohereToken = configs.CohereToken
// }

func Generate(prompt string) (string, error) {
	co, err := cohere.CreateClient(os.Getenv("COHERE_TOKEN"))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	response, err := co.Generate(cohere.GenerateOptions{
		Model:             "command-xlarge-nightly",
		Prompt:            prompt,
		MaxTokens:         600,
		Temperature:       0.9,
		K:                 0,
		P:                 0.75,
		FrequencyPenalty:  0,
		PresencePenalty:   0,
		StopSequences:     []string{},
		ReturnLikelihoods: "NONE",
	  })
	  if err != nil {
		fmt.Println(err)
		return "", err
	  }

	fmt.Println("Prompt:", prompt)
	fmt.Println("Response:", response.Generations[0])
	return response.Generations[0].Text, nil
}
