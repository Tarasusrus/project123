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

# Загрузка переменных из файла .env
load_env ".env"

# Объединение шаблона YAML с переменными окружения и запись результата в файл конфигурации
envsubst < configs/template.config.yml > configs/config.yml
