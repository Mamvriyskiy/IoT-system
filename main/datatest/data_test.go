package main

import (
    // "context"
    // "testing"

    // "github.com/testcontainers/testcontainers-go"
    // "github.com/testcontainers/testcontainers-go/wait"
)

// func TestWithRedis(t *testing.T) {
//     ctx := context.Background()
//     req := testcontainers.ContainerRequest{
//         Image:        "postgres:15.4-alpine",
//         ExposedPorts: []string{"6379/tcp"},
//         WaitingFor:   wait.ForLog("Ready to accept connections"),
//     }
//     redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
//         ContainerRequest: req,
//         Started:          true,
//     })
//     if err != nil {
//         t.Fatalf("Could not start redis: %s", err)
//     }
//     defer func() {
//         if err := redisC.Terminate(ctx); err != nil {
//             t.Fatalf("Could not stop redis: %s", err)
//         }
//     }()
// }
