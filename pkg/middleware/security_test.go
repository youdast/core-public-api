package middleware

import (
	"net/http/httptest"
	"testing"
	"youdast/core-public-api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSecurityMiddleware(t *testing.T) {
	// Configure Fiber to trust X-Forwarded-For for testing
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0/0"}, // Trust all for testing
		ProxyHeader:             "X-Forwarded-For",
	})

	cfg := &config.Config{
		AllowedIPs:   "192.168.1.100",
		AllowedHosts: "api.example.com,localhost",
	}

	app.Use(SecurityMiddleware(cfg))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	tests := []struct {
		name       string
		ip         string
		host       string
		statusCode int
	}{
		{
			name:       "Allow Localhost IP",
			ip:         "127.0.0.1",
			host:       "random.com",
			statusCode: 200,
		},
		{
			name:       "Allow Whitelisted IP",
			ip:         "192.168.1.100",
			host:       "random.com",
			statusCode: 200,
		},
		{
			name:       "Allow Whitelisted Host",
			ip:         "10.0.0.1", // Unknown IP
			host:       "api.example.com",
			statusCode: 200,
		},
		{
			name:       "Allow Whitelisted Host Localhost",
			ip:         "10.0.0.1",
			host:       "localhost",
			statusCode: 200,
		},
		{
			name:       "Block Unknown IP and Host",
			ip:         "10.0.0.1",
			host:       "evil.com",
			statusCode: 403,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			req.Header.Set("X-Forwarded-For", tt.ip)
			req.Host = tt.host

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.statusCode, resp.StatusCode)
		})
	}
}
