package config

import (
	"context"
	"encoding/json"

	// "fmt"
	vault "github.com/hashicorp/vault/api"
	// "github.com/hashicorp/vault/api/auth/approle"
	"log"
)

func VaultSecrets(vaultAdd, vaultToken, secretPath string) (*Config, error) {
	vaultConfig := vault.DefaultConfig()

	vaultConfig.Address = vaultAdd
	client, err := vault.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}

	client.SetToken(vaultToken)

	// Read a secret from the default mount path for KV v2 in dev mode, "secret"
	secret, err := client.KVv2("secret").Get(context.Background(), secretPath)

	// fmt.Printf("from vault:: ,%v \n", secret.Data)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	j, err := json.Marshal(secret.Data)

	if err != nil {
		log.Fatalf("unable to marshal secrets: %v", err)
	}

	config := &Config{}

	err = json.Unmarshal(j, config)
	if err != nil {
		log.Fatalf("unable to parse secrets: %v", err)
	}

	return config, nil
}
