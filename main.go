package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	BUILD_VERSION        = "0.1.0" // x-release-please-version
	INPUT_GITHUB_SECRETS = "GITHUB_SECRETS"
)

var versionFlag bool

func init() {
	flag.BoolVar(&versionFlag, "V", false, "show version")
	flag.BoolVar(&versionFlag, "version", false, "show version")
}

func main() {
	// Parse input parameters
	flag.Parse()

	if versionFlag {
		fmt.Println(BUILD_VERSION)
		return
	}

	secretId := flag.Arg(0)
	if secretId == "" {
		log.Fatal("secret ID is empty")
	}

	// Find secret value
	secret, err := findSecret(secretId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(secret)
}

func findSecret(secretId string) (secret string, err error) {
	var secrets map[string]string

	githubSecrets := os.Getenv(INPUT_GITHUB_SECRETS)
	if githubSecrets == "" {
		return secret, fmt.Errorf("environment variable %s is empty", INPUT_GITHUB_SECRETS)
	}

	// Read JSON into a map
	err = json.Unmarshal([]byte(githubSecrets), &secrets)
	if err != nil {
		return
	}

	// Find the secret in the map
	secret, ok := secrets[secretId]
	if !ok {
		return secret, fmt.Errorf("entry '%s' not found", secretId)
	}

	return
}
