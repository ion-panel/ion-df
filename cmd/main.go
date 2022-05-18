package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	ion_df "ion-df"
	"ion-df/endpoints"
	"os"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	config, err := readConfig()
	if err != nil {
		log.Fatalln(err)
	}

	srv := server.New(&config, log)

	srv.CloseOnProgramEnd()
	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}
	///////////////////////////////////////////////////////////

	go func() {
		conf := &ion_df.Config{}
		if !conf.Exists() {
			fmt.Println("ion-config.json file does not exist. Creating...")
			conf.Create()

			fmt.Println("Created config file. Please restart Ion.")
			return
		}

		serverApi := endpoints.ServerApi{Server: srv}

		router := gin.Default()
		router.GET("/"+conf.GetApiKey()+"/server/maxPlayers", serverApi.GetMaxPlayerCount)
		router.GET("/"+conf.GetApiKey()+"/server/playerCount", serverApi.GetPlayerCount)
		router.GET("/"+conf.GetApiKey()+"/server/open", serverApi.OpenServer)
		router.GET("/"+conf.GetApiKey()+"/server/close", serverApi.CloseServer)

		router.GET("/" + conf.GetApiKey() + "/server/close")

		err = router.Run(conf.GetHost())
		if err != nil {
			return
		}
	}()

	///////////////////////////////////////////////////////////
	for {
		if _, err := srv.Accept(); err != nil {
			return
		}
	}
}

// readConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
func readConfig() (server.Config, error) {
	c := server.DefaultConfig()
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("failed encoding default config: %v", err)
		}
		if err := ioutil.WriteFile("config.toml", data, 0644); err != nil {
			return c, fmt.Errorf("failed creating config: %v", err)
		}
		return c, nil
	}
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		return c, fmt.Errorf("error reading config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("error decoding config: %v", err)
	}
	return c, nil
}
