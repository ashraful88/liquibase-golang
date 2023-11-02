package cmd

import (
	"fmt"
	"os/exec"
)

// CreateConnectionString creates connection string for flyway
func LiquibaseCreateConnectionString(dbHost, dbPort, dbName, dbUser, dbPass, changeFileLoc string) string {
	if dbPort == "" {
		dbPort = "5432"
	}
	if changeFileLoc == "" {
		changeFileLoc = "changelog.xml"
	}
	return fmt.Sprintf(`--url=jdbc:postgresql://%s:%s/%s --username=%s --password=%s --changelog-file=%s`,
		dbHost, dbPort, dbName, dbUser, dbPass, changeFileLoc)
}

// LiquibaseValidate runs database schema validation
func LiquibaseValidate(dbConStr string) (string, error) {
	cmd := fmt.Sprintf("liquibase validate %s", dbConStr)
	outInfo, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	return string(outInfo), err
}

// LiquibaseRollbackDry generate rollback sql
func LiquibaseRollbackDry(dbConStr, versionTag string) (string, error) {
	cmd := fmt.Sprintf("liquibase rollback-sql --tag=%s %s", versionTag, dbConStr)
	cleanSchema, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	return string(cleanSchema), err
}

// LiquibaseMigrate runs database schema migration
func LiquibaseMigrate(dbConStr string) (string, error) {
	cmd := fmt.Sprintf("liquibase update %s", dbConStr)
	out, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	return string(out), err
}
