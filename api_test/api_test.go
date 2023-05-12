package api_test

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupTestContainers(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "postgres:13.3-alpine",
		Env: map[string]string{
			"POSTGRES_PASSWORD": cfg.Password,
			"POSTGRES_USER":     cfg.User,
			"POSTGRES_DB":       cfg.DBName,
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := postgresC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

}
func APITest(t *testing.T) {
	setupTestContainers(&testing.T{})

}
