package services

import (
	"context"
	"github.com/FleekHQ/space-poc/config"
	"github.com/FleekHQ/space-poc/core/env"
	"github.com/FleekHQ/space-poc/core/space/domain"
	"github.com/FleekHQ/space-poc/core/store"
	tc "github.com/FleekHQ/space-poc/core/textile/client"
)

// Implementation for space.Service
type Space struct {
	store store.Store
	cfg   config.Config
	env   env.SpaceEnv
	tc    tc.Client
}




func (s *Space) GetConfig(ctx context.Context) domain.AppConfig {
	return domain.AppConfig{
		FolderPath: s.cfg.GetString(config.SpaceFolderPath, ""),
		Port:       s.cfg.GetInt(config.SpaceServerPort, "-1"),
		AppPath:    s.env.WorkingFolder(),
	}

}

func NewSpace(st store.Store, tc tc.Client, cfg config.Config, env env.SpaceEnv) *Space {
	return &Space{
		store: st,
		cfg:   cfg,
		env:   env,
		tc:    tc,
	}
}