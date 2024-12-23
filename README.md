# YandexAcademyOfGolang

Проект за Спринт 1, калькулятор, работающий через сервер. 

Работа проекта заключается в том, что пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.

Перед началом работы склонируйте к себе в IDE репозиторий: https://github.com/BorislavTimoshin/YandexAcademyOfGolang.git
 
 Файл cmd/main.go - главный файл для запуска программы.

Файл internal/application/application.go - файл-посредник для диалога пользователем с приложением.

Файл pkg/calculator/calculator.go - код алгоритма работы калькулятора

Файл pkg/calculator/calculator_test.go - код с некоторыми тестами.

Файл pkg/calculator/calculator_handler.go - код работы с серверной частью программы.

Ознакомьтесь с проектом и запустите его:
go run ./cmd/main.go

ТЗ Яндекса:
У сервиса 1 endpoint с url-ом /api/v1/calculate. Пользователь отправляет на этот url POST-запрос с телом:

{ "expression": "выражение, которое ввёл пользователь" } В ответ пользователь получает HTTP-ответ с телом:

{ "result": "результат выражения" } и кодом 200, если выражение вычислено успешно, либо HTTP-ответ с телом:

{ "error": "Expression is not valid" } и кодом 422, если входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций пользователь ввёл символ английского алфавита.

Ещё один вариант HTTP-ответа:

{ "error": "Internal server error" } и код 500 в случае какой-либо иной ошибки («Что-то пошло не так»).

Чтобы получить ответ, нужно отправить запрос через Postman или cURL:

URL: http://localhost:8080/api/v1/calculate Метод: POST Заголовок: Content-Type: application/json И сам пример, который будет приведен ниже. Для этого нужен сайт Postman - https://www.postman.com/. Выберите метод POST, введите URL: http://localhost:8080/api/v1/calculate. Перейдите на вкладку Headers, и добавьте новый ключ (key): Content-Type, и в нем значение (value): application/json. Потом перейдите во вкладку Body, выберите raw и выберите JSON данные. Далее вводите тестовые данные в формате JSON: { "expression": "5 + 3 * 2" } В итоге мы получим следующее:

{"result": 11.000000} Если мы попробуем вставить

{ "expression": "5 * (5 + 7))" } То получим следующее:

{"error": "Expression is not valid"} То есть ошибку 422

 Или например если вставить вот такое выражение:

{ "expression": "5 / 0" }, то выдаст ошибку 500: Internal server error