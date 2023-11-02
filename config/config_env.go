package config

type kv map[string]string

// write configs here, not for keeping credentials
var envConfigs = map[string]kv{
	"local": {
		"API_PORT":                 "8088",
		"LIQUIBASE_CHANGELOG_FILE": "changelog.xml",
	},
	"staging": {
		"API_PORT":                 "80",
		"LIQUIBASE_CHANGELOG_FILE": "changelog.xml",
	},
	"production": {
		"API_PORT":                 "80",
		"LIQUIBASE_CHANGELOG_FILE": "changelog.xml",
	},
}
