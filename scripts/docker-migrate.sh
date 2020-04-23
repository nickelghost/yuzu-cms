export APP_PORT=3000
export APP_SECRET=secret
export APP_USER_NAME=admin
export APP_USER_PASS=$2y$12$JbJmIwUnnXaaiZCYB8pNm.gt/KMdQ4/alJTrWWLkOasd3iNXZRpEu
export APP_THEME=default
export APP_FORWARD_WEBPACK=false

export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASS=password
export DB_NAME=yuzu
export DB_REQUIRE_SSL=false

go run . -migrate -seed
