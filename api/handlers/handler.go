package handlers

import (
	"goApiStartetProject/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Handler struct {
	Env  *config.Env
	EthClient 	*ethclient.Client
}

func NewHandler(ethClient *ethclient.Client, env *config.Env) *Handler{
	return &Handler{
		Env: env,
		EthClient: ethClient,
	}
}