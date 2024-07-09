package handler

import (
	"go.uber.org/zap"

	"github.com/nautilusgames/demo/auth/internal/ent"
	"github.com/nautilusgames/demo/auth/token"
	"github.com/nautilusgames/demo/config/pb"
)

type Handler struct {
	logger *zap.Logger
	
	cfg    *pb.Config
	entClient *ent.Client

	accessToken       token.Maker
	playerTenantToken token.Maker
}

func New(
	logger *zap.Logger,
	cfg *pb.Config,
	entClient *ent.Client,
	accessToken token.Maker,
	playerTenantToken token.Maker,
) Handler {
	return Handler{
		logger:            logger,
		cfg:               cfg,
		entClient:         entClient,
		accessToken:       accessToken,
		playerTenantToken: playerTenantToken,
	}
}
