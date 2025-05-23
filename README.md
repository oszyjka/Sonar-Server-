# Sonar-Server-
Add liter:
1. Init(bash): **curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2**
2. Run: **golangci-lint run ./...**
3. Make hook executable : **chmod +x .git/hooks/pre-commit**

[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=oszyjka_Sonar-Server-&metric=bugs)](https://sonarcloud.io/summary/new_code?id=oszyjka_Sonar-Server-) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=oszyjka_Sonar-Server-&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=oszyjka_Sonar-Server-) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=oszyjka_Sonar-Server-&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=oszyjka_Sonar-Server-) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=oszyjka_Sonar-Server-&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=oszyjka_Sonar-Server-)