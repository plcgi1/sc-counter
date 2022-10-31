package listeners

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/plcgi1/address-caller-stack/go-listener/internal/config"
	"github.com/plcgi1/address-caller-stack/go-listener/internal/contracts"
	"github.com/sirupsen/logrus"
)

type IAddressCountedRequest interface {
	ListenAddressCounted()
	IsError(err error, errLabel string, logger *logrus.Entry) bool
	ProcessEvent(value *contracts.AddressCallerCounterUserIncremented, logger *logrus.Entry)
}

type AddressCountedRequestListener struct {
	contractClient   *contracts.AddressCallerCounter
	logger           *logrus.Entry
	config           *config.Config
	targetEthAddress common.Address
}

func NewAddressCountedService(cfg *config.Config, contractClient *contracts.AddressCallerCounter, logger *logrus.Entry) IAddressCountedRequest {
	return &AddressCountedRequestListener{
		logger:         logger,
		config:         cfg,
		contractClient: contractClient,
	}
}

func (l *AddressCountedRequestListener) ListenAddressCounted() {
	logger := l.logger.
		WithField("function", "ListenAddressCounted")

	logger.Info("Started")

	addressCountedChan := make(chan *contracts.AddressCallerCounterUserIncremented)

	//// websocket: close 1006 (abnormal closure): unexpected EOF issue -
	////              https://github.com/ethereum/go-ethereum/issues/22266
	addressCountedIterator := event.Resubscribe(l.config.WebsocketTimeout, func(ctx context.Context) (event.Subscription, error) {
		return l.contractClient.WatchUserIncremented(
			&bind.WatchOpts{Start: nil, Context: context.Background()},
			addressCountedChan,
			nil,
			nil,
		)
	})
	for {
		select {
		case errChan := <-addressCountedIterator.Err():
			{
				notOk := l.IsError(errChan, "failed to get channel addressCountedIterator", logger)
				if notOk {
					return
				}
			}
		case value := <-addressCountedChan:
			l.ProcessEvent(value, logger)
		} // END select {
	} // END for {
}

func (l *AddressCountedRequestListener) IsError(err error, errLabel string, logger *logrus.Entry) bool {
	if err != nil {
		logger.WithError(err).Error(errLabel)
		return true
	}
	return false
}

func (l *AddressCountedRequestListener) ProcessEvent(value *contracts.AddressCallerCounterUserIncremented, logger *logrus.Entry) {
	logger.Debug("Event fired")

	logger.
		WithField("Event address", value.Account.Hex()).
		WithField("Event count", value.Count).
		Debug("Event fired")
}
