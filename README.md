# Тестовое задание 

## Запуск приложения

Для запуска приложения воспользуйтесь двумя командами:
- `make run` - запускает приложение
- `make build` - компилирует приложение

Порт задается через переменное окружение, стандартно работает по порту :8085

## Примеры использования

### Создание конфига 
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"cmd":"df -h","os":"linux","stdin":""}' \
   -k https://localhost:8085/api/v1/remote-execution
```

#### Пример команды для Linux
```
{
    "cmd": "df -h", 
    "os": "linux", 
    "stdin":""
}
```

#### Пример команды для Windows 

```
{
    "cmd": "ping.exe https://google.com", 
    "os": "windows", "
    stdin":""
}
```

