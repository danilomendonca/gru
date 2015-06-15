package action

import (
	"crypto/rand"
	"encoding/hex"

	log "github.com/Sirupsen/logrus"
)

type Start struct{}

func (p *Start) Name() string {
	return "start"
}

func (p *Start) Initialize() error {
	return nil
}

func (p *Start) Run(config *GruActionConfig) error {
	var err error = nil
	var uuid string
	if config.Target == "container" {
		err = config.Client.StartContainer(config.Target, config.HostConfig)
	} else {
		uuid, err = generateUUID()
		name := config.Service + uuid

		log.WithFields(log.Fields{
			"name": name,
			"id":   "TODO",
		}).Infoln("Starting new container")
		//config.Client.CreateContainer(config.HostConfig, name)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"id":    config.Target,
			"error": err,
		}).Errorln("Error starting container")
		return err
	}

	return nil
}

func generateUUID() (string, error) {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return "", err
	}

	u[8] = (u[8] | 0x80) & 0xBF // what does this do?
	u[6] = (u[6] | 0x40) & 0x4F // what does this do?

	return hex.EncodeToString(u), nil
}
