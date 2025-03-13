BACKEND_GEN_DIR=../backend/internal/interface/oapi
FRONTEND_GEN_DIR=../frontend/src/api

generate-backend:
	oapi-codegen -package oapi -generate types,gin-server -o $(BACKEND_GEN_DIR)/openapi.gen.go open-api.yaml

generate-frontend:
	npx openapi-typescript open-api.yaml --output $(FRONTEND_GEN_DIR)/openapi.gen.ts

generate: generate-backend generate-frontend
