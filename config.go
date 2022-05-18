package ion_df

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ApiKey string `json:"api_key"`
	Host   string `json:"host"`
}

func (c *Config) GetApiKey() string {
	return c.Get().ApiKey
}

func (c *Config) GetHost() string {
	return c.Get().Host
}

func (c *Config) Create() os.File {
	file, err := os.Create("ion_df_config.json")

	file.Write([]byte(`{
    "api_key": "endpoints-key",
    "host": "localhost:9090"
}`))
	if err != nil {
		panic(err)
	}

	return *file
}

func (c *Config) Exists() bool {
	if _, err := os.Stat("ion_df_config.json"); os.IsNotExist(err) {
		return false
	}

	return true
}

func (c *Config) Get() Config {
	var configBuffer Config

	file, err := os.Open("ion_df_config.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	decodedJson := json.NewDecoder(file)
	err = decodedJson.Decode(&configBuffer)
	if err != nil {
		fmt.Println("JSON Decode Err: ", err)
	}

	return configBuffer
}
