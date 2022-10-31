package main

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/plcgi1/address-caller-stack/go-listener/internal/config"
	"github.com/plcgi1/address-caller-stack/go-listener/internal/contracts"
	"github.com/plcgi1/address-caller-stack/go-listener/internal/listeners"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

func main() {
	//parse config
	cfg, err := config.GetConfig(afero.NewOsFs())
	if err != nil {
		panic(errors.Wrap(err, "failed to load config"))
	}
	//setup logger
	baseLog := log.New()
	if cfg.DebugLog {
		baseLog.SetLevel(log.DebugLevel)
	}
	logger := log.NewEntry(baseLog)
	// ethereum setup
	ethClient, err := ethclient.Dial(cfg.WssRpcUrl)
	if err != nil {
		errLabel := "failed to create eth client dialing"
		logger.WithField("WssRpcUrl", cfg.WssRpcUrl).Error(errors.Wrap(err, errLabel))
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	ethAddress := common.HexToAddress(cfg.ContractAddress)
	contractClient, err := contracts.NewAddressCallerCounter(ethAddress, ethClient)

	if err != nil {
		logger.
			WithField("ethAddress", ethAddress).
			WithError(err).
			Error("failed to get contractClient")
		return
	}
	listener := listeners.NewAddressCountedService(
		&cfg,
		contractClient,
		logger,
	)
	go listener.ListenAddressCounted()

	wg.Wait()
}
