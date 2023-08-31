# Avito-task

#### !!! Для запуска проекта в корневой папке должен лежать файл config.yaml!!!

## Установка

    go mod download

## Старт приложения

     go run .\cmd\main.go    

## Применение миграций

     Для миграций был использован инструмент goose. Чтобы его использовать рекомендуется установить переменные среды:
           GOOSE_DBSTRING=host=0.0.0.0 user=postgres password=postgres dbname=postgres sslmode=disable (с вашими данными бд)
           GOOSE_DRIVER=postgres
           
     Далее можно применять основные команды:
    
           goose up (применение миграций)
           goose down (откат последней миграции)
           gooose create <название_миграции> sql (создание новой миграции)
           goose status (просмотр примененных миграциий)

## Запуск в Docker

     docker-compose build
     docker-compose up

# HTTP API

## Handlers

### URL `/add-segment` Method `POST`

    {
    "PercentUsers":0.8,
    "name":"avito_message"
    }

### reply

    {
	"name": "avito_message"
    }

### `/delete-segment` Method `DELETE`

    {
    "name":"avito_message"
    }

### reply

    {
	"name": "avito_message"
    }

## URL `/add-segment-for-user` Method `POST`

    {
	"userId":1,
	"segmentName": ["avito_voice", "avito_message"]

}

### reply

    {
    "name": [
    "avito_voice",
    "avito_message"
    ]
    }

### URL `/get-user-segments` Method `GET``

    {
	"userId":1
    }

### reply

    {
	"name": {
		"UserId": 1,
		"Segments": [
			{
				"Name": "avito_message",
				"PercentUsers": 0 -> 
                -> # при скольки процентах юзер попал в сигмент, 
                   # добавлен в ручную если процент 0
        
			},
			{
				"Name": "avito_voice",
				"PercentUsers": 0
			}
		    ]
	    }
    }

## URL `/delete-user-segments` Method `PUT`

    {
	"userId":1,
	"segmentName":["avito_message","avito_voice"]
    }
### reply
    {
	"name": "avito_message"
    }

### URL `/get-user-history` Method `POST`

    {
	"userId":1,
	"dateStart":"2021.01.02",
	"dateEnd":"2026-09-31"
    }

### reply

    {
	"url": "http://localhost:8080/save-history/1693506444.csv"
    }
Вопросы с которыми столкнулся:
#### 1) Нужно ли реально удалять сегменты из таблиц?
    Я решил оставлять записи для просмотра истории и не удалять их из таблиц.
    Вместо удаления реализовал доп поле is_removed, чтобы отслеживать удален сигмент или нет.
    Аналогично для таблицы, которая реализует связь "Многие ко многим" user_segment.
    Решение было такое, потому что нужно всегда суметь достать информацию о добавлении или удалении юзера из сегментов.
#### 2) Как должен отдаваться файл .csv для истории и нужно ли перезаписывать его каждый раз?
    Я сделал динамическую ссылку на файл, а сам файл никогда не перезаписывается.
    Под каждый запрос создается новый .csv с timestamp в названии
#### С миграциями удобней:)
#### 3) при добавлении сегмента с указанием процентов(сколько пользователей попадут в этот сегмент) куда округлять? 
    При реализации 3 доп задания, наткнулся на такую штуку.
    Если пользователей 7 а процент попадания в сегмент указан 0.5, то это будет 3 юзера, а не 4
    решил оставить так:)
