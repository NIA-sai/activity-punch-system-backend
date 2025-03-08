package test

import (
	"activity-punch-system-backend/config"
	"activity-punch-system-backend/internal/global/database"
	"activity-punch-system-backend/tools"
	"context"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
)

const (
	EnvFileName   = "docker-compose.env.yaml"
	ConfigFilName = "config.example.yaml"
)

func IsTest() bool {
	return os.Getenv("ENV") == "test"
}

func SetupEnvironment(t *testing.T) {
	t.Setenv("ENV", "test")
	compose, err := tc.NewDockerCompose(tools.SearchFile(EnvFileName))
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal))
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	require.NoError(t,
		compose.WaitForService("mysql", wait.ForLog("port: 3306  MySQL Community Server")).Up(ctx, tc.Wait(true)),
	)

	config.Init(tools.SearchFile(ConfigFilName))
	database.Init()
}
