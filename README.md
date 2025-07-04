# Приложение по управлению кошельком

## Описание

Это проект на Go, который реализует API для управления кошельком. Он использует библиотеку `gin` для упрощенной работы с HTTP, а также `gorm` для работы с PostgreSQL.


## Используемые технологии

*   **Golang:** Язык программирования.
*   **Gin:** HTTP-фреймворк.
*   **PostgreSQL:** БД для хранения задач.
*   **Docker и Docker Compose:** Для контейнеризации.

## Инструкция по запуску приложения

1.  **Убедитесь, что у вас установлен и запущен Docker и Docker Compose.**

2.  **Склонируйте репозиторий:**

    ```bash
    git clone https://github.com/mrMaks2/test_task_ITK.git
    ```

3.  **Запустите проект с помощью команды:**

    ```bash
    docker-compose --env-file ./config.env up -d
    ```
4. **Для запуска теста перейдите в папку tests с помощью команды:**

    ```bash
    cd tests
    ```

5. **И запустите тесты с помощью команды:**

    ```bash
    go test
    ```

4.  **Для остановки проекта используйте команду:**

    ```bash
    docker-compose --env-file ./config.env down
    ```


## Переменные окружения

В файле .env имеются следующие переменные (желательно в файле .env переменные окружения заменить на свои, но и без этого будет работать, так как файл не стал добавлять в .gitignore):

*   **POSTGRES_USER:**  Имя пользователя для PostgreSQL.
*   **POSTGRES_PASSWORD:** Пароль пользователя для PostgreSQL.
*   **POSTGRES_PORT:** Порт для PostgreSQL.
*   **POSTGRES_NAME:** Имя базы данных PostgreSQL.
*   **POSTGRES_USE_SSL:** Использовать ли SSL для PostgreSQL.
*   **POSTGRES_HOST:** Хост для PostgreSQL.
*   **CONN_HOST:** Порт для сервиса приложения.

## API

*   **POST /api/v1/wallet:** Выполняет логику изменения баланса кошелька.
*   **GET /api/v1/wallets/:walletId** Получить баланс кошелька по ID.