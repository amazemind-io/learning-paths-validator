validate:
	cd src && go run . path ../backend/python.yaml ../domains/index.yaml
.PHONY: validate