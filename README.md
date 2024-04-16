# student_marks

## How to run this project

1. Please install https://github.com/golang-migrate/migrate

2. set below variables as environment variables on your system (Reference: [Windows](https://phoenixnap.com/kb/windows-set-environment-variable), [Mac](https://phoenixnap.com/kb/set-environment-variable-mac), [Linux](https://phoenixnap.com/kb/linux-set-environment-variable))

```
export DB_NAME=postgres
export DB_PORT=5432
export DB_HOST=localhost
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export APP_PORT=8080
```

3. docker-compose up

4. migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

5. hit API add record
```
curl --location 'http://localhost:8080/api/v1/records/add' \
--header 'Content-Type: application/json' \
--data '{
    "name":"test-1",
    "marks":[1,2,3,4,5]
}'
```

6. hit API get record
```
curl --location 'http://localhost:8080/api/v1/records/find' \
--header 'Content-Type: application/json' \
--data '{
"startDate": "2024-01-26",
"endDate": "2024-05-02",
"minCount": 1,
"maxCount": 300
}'
```

