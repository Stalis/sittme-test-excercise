# Stalis Stream Service
[![Go Report Card](https://goreportcard.com/badge/github.com/Stalis/sittme-test-excercise)](https://goreportcard.com/report/github.com/Stalis/sittme-test-excercise)

## Описание
SSS - заготовка микросервиса, предназначенного для потоковых трансляций данных на любые устройства и сервисы

## Запуск
Для простого запуска необходимо перейти в корневую директорию репозитория и ввести команду
```
go run . <run-args>
```
Для конфигурирования используются следующие параметры запуска:
```
-host string
        Host IP address or DNS name (default "localhost")
-path string
        Path for binding service (default "/")
-port int
        Http server port (default 8080)
-read-timeout duration
        Timeout for reading data
-stream-timeout duration
        Timeout for finish stream after interrupted (default 5s)
-tls-cert string
        Path to TLS certificate (default "cert.pem")
-tls-key string
        Path to TLS key file (default "key.pem")
-use-tls
        Enables HTTPS by TLS certificate
-write-timeout duration
        Timeout for writing data
```

## API
В качестве API используется формат JSON REST по протоколу http
### Создание трансляции
Для создания трансляции необходимо создать пустой POST-запрос на адрес сервиса. Ответ будет в таком формате:
```json
{
    "type": "stream",
    "id": "<id трансляции>",
    "attributes": {
        "created": "<дата и время создания трансляции>",
        "state": "created",
    }
}
```
### Получение информации о трансляции
Для этого нужно сделать GET-запрос с одним аргументом `id`, в котором нужно передать id созданной трансляции. Ответ будет такой же, как в предыдущем пункте
### Изменение состояния трансляции
Для изменения состояния нужно сделать PUT-запрос с двумя параметрами:

- `id` - id трансляции
- `state` - желаемое состояние

Состояния могут указываться как текстом(независимо то регистра), так и цифрами:

- `"Created"` == 0
- `"Active"` == 1
- `"Interrupted"` == 2
- `"Finished"` == 3

При этом невозможно установить состояние `Created`, но из него можно установить состояние `Active`, а из него уже либо `Finished`, либо `Interrupted`. Из `Interrupted` можно перейти в `Active` в течение `-stream-timeout` или в `Finished`. По истечению `-stream-timeout` состояние меняется на `Finished`

## Удаление трансляции
Для удаления трансляции необходимо сделать DELETE-запрос с одним аргументом `id` - id трансляции, которую необходимо удалить. Особого формата ответа не предусмотрено

## Ошибки
При ошибках в работе сервера над запросом клиента, сервер возвращает следующий JSON-объект:

```json
{
    "type": "error",
    "error": "<текст с информацией по ошибке>"
}
```