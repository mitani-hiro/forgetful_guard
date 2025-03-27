BACKEND_GEN_DIR=../backend/internal/interface/oapi
FRONTEND_GEN_DIR=../frontend/src/api

generate-backend:
	swagger-cli bundle open-api.yaml -o merged-api.yaml
	oapi-codegen -package oapi -generate types,gin-server -o $(BACKEND_GEN_DIR)/openapi.gen.go merged-api.yaml
	rm merged-api.yaml

generate-frontend:
	swagger-cli bundle open-api.yaml -o merged-api.yaml
	npx openapi-typescript merged-api.yaml --output $(FRONTEND_GEN_DIR)/openapi.gen.ts
	rm merged-api.yaml

generate: generate-backend generate-frontend

generate-html:
	redoc-cli bundle open-api.yaml -o docs/open-api.html
