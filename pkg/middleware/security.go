package middleware

import (
	"strings"
	"youdast/core-public-api/config"

	"github.com/gofiber/fiber/v2"
)

// SecurityMiddleware handles IP whitelisting and Host Header validation
func SecurityMiddleware(cfg *config.Config) fiber.Handler {
	// Parse Allowed IPs
	allowedIPs := make(map[string]bool)
	// Always allow localhost
	allowedIPs["127.0.0.1"] = true
	allowedIPs["::1"] = true

	if cfg.AllowedIPs != "" {
		ips := strings.Split(cfg.AllowedIPs, ",")
		for _, ip := range ips {
			allowedIPs[strings.TrimSpace(ip)] = true
		}
	}

	// Parse Allowed Hosts
	allowedHosts := make(map[string]bool)
	if cfg.AllowedHosts != "" {
		hosts := strings.Split(cfg.AllowedHosts, ",")
		for _, host := range hosts {
			allowedHosts[strings.TrimSpace(host)] = true
		}
	}

	return func(c *fiber.Ctx) error {
		// 1. Check IP Whitelist
		clientIP := c.IP()
		if allowedIPs[clientIP] {
			return c.Next()
		}

		// 2. Check Host Header (if configured)
		// If AllowedHosts is empty, we skip this check (or we could block everything else, but let's be safe)
		// If AllowedHosts is configured, we MUST match one of them.
		if len(allowedHosts) > 0 {
			host := c.Hostname()
			if allowedHosts[host] {
				return c.Next()
			}
		} else {
			// If no hosts are configured, and IP didn't match, we might want to allow it?
			// Based on user request "block access hit to api this besides this server or localhost",
			// if they haven't configured hosts yet, they might lock themselves out if we block here.
			// However, the plan says "If neither passes, BLOCK".
			// Let's stick to the plan but be careful.
			// If AllowedHosts is empty, we only rely on IP whitelist.
			// So if IP is not in whitelist, and no hosts configured -> Block.
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access Denied",
		})
	}
}
