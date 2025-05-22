# Sonar-Server-
Add liter:
1. Init(bash): **curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2**
2. Run: **golangci-lint run ./...**
3. Make hook executable : **chmod +x .git/hooks/pre-commit**

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=oszyjka_Sonar-Server-&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=oszyjka_Sonar-Server-)