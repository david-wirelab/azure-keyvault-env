package azure

import (
  "encoding/json"
  "context"
  "fmt"
  "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
  "github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
  "log"
)

func GetSecretData(name, keyvault string) (map[string]string, error) {
  var secrets map[string]string

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault client
  vaultUri := fmt.Sprintf("https://%s.vault.azure.net/", keyvault)
	client, err := azsecrets.NewClient(vaultUri, cred, nil)

	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	result, err := client.GetSecret(context.Background(), name, version, nil)
	if err != nil {
	   return nil, err
	}

  // Convert plaintext data received from Azure into a JSON type map of strings,
  // then read the key value pairs and pass them downstream.
  data_raw := map[string]string {
    name: *result.Value,
    }
  data, _ := json.Marshal(data_raw)
  err = json.Unmarshal(data, &secrets)
  if err != nil {
    return nil, fmt.Errorf("%s is not a key-pair secret", name)
  }

  return secrets, nil
}
