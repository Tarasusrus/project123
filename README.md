# project123

## Как запустить сервис?

1. Копируем .env.example в .env и меняем в .env настройки.
```
cp .env.example .env
```

2. Билдим и поднимаем контейнеры в docker-compose (app + postgreSQL containers):

### Команды

Сбилдить и поднять контейнеры:

```
make up
```

Завершить работу сервиса:

```
make down
```

