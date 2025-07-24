package models

import "fmt"

type Role string

const (
	RoleAdmin    Role = "ims:admin"
	RoleUser     Role = "ims:reader"
	RoleScraper  Role = "ims:scraper"
	RoleExporter Role = "ims:exporter"
)

func TryParseRole(role string) (Role, error) {
	switch role {
	case string(RoleAdmin):
		return RoleAdmin, nil
	case string(RoleUser):
		return RoleUser, nil
	case string(RoleScraper):
		return RoleScraper, nil
	default:
		if len(role) >= len(string(RoleExporter)) && role[:len(string(RoleExporter))] == string(RoleExporter) {
			return RoleExporter, nil
		}

		return "", fmt.Errorf("invalid role: %s", role)
	}
}
