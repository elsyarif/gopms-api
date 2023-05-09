# go-auth-api
REST API implementing Hexagonal Architecture, including authentication

### Feature
- [x] User Register
- [x] Authentication
- [x] Refresh Token

### Endpoint

| Method  | path            | Description                               |
|---------|-----------------|-------------------------------------------|
| POST    | /users          | Create a new user                         |
| GET     | /user/profile   | Get user information with protected route |
| POST    | /authentication | Login with credential                     |
| PUT     | /authentication | Create new access token                   |
| Delete  | /authentication | logout                                    |

## Stack
- Gin
- Postgres
- Sqlx
- Logrus
- Dotenv