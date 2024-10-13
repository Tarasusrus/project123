# Установка переменных
SWAGGER_CMD := swag init -g cmd/api/main.go -o ./docs --parseInternal

# Определение цели как .PHONY, чтобы make выполнял её каждый раз
.PHONY: docs install-swag generate-docs

# Цель для генерации документации
docs:
	$(SWAGGER_CMD)