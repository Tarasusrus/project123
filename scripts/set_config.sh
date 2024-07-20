#!/bin/bash

# Переход в каталог со скриптом
cd "$(dirname "$0")" || exit 1
cd ..

# Функция для загрузки переменных из файла .env без использования export
load_env() {
    local env_file="$1"
    if test -f "$env_file"; then
        while IFS= read -r line || [ -n "$line" ]; do
            if [[ "$line" =~ ^[a-zA-Z_][a-zA-Z0-9_]*= ]]; then
                export "$line"
            fi
        done < "$env_file"
    fi
}

# Функция для генерации случайной строки
generate_random_string() {
    local length=$1
    openssl rand -base64 "$((length * 3 / 4))" | tr -dc 'a-zA-Z0-9' | head -c "$length" || true
}

# Функция для проверки и добавления SERVICE_API_KEY в .env
add_integration_api_key() {
    local env_file="$1"
    local key_name="SERVICE_API_KEY"
    local key_value=$(generate_random_string 256)

    if ! grep -q "^${key_name}=" "$env_file"; then
        echo "${key_name}=${key_value}" >> "$env_file"
    fi
}

# Загрузка переменных из файла .env
load_env ".env"

# Добавление SERVICE_API_KEY в .env, если его нет
add_integration_api_key ".env"

# Объединение шаблона YAML с переменными окружения и запись результата в файл конфигурации
envsubst < configs/template.config.yml > configs/config.yml
