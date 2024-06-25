# Instruction


## How to run:
### Run app on local machine:
- make sure you have go and docker installed in your computer
- pull git repository: `git clone https://github.com/dtpthao/Sample-Go-Server.git && cd Sample-Go-Server`
- Run database and kafka services on docker service: `docker-compose up -d`
- Seed admin user: `go run seeds/main.go [newusername] [password]`
- Run program `go run main.go`


## APIs

### Login
```
POST /accounts/login
```
Request:
```json
{
  "username": [username],
  "password": [password]
}
```

Response: status OK (200)
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTkwOTIyOTUsImlhdCI6MTcxOTA4ODY5NSwidXNlcm5hbWUiOiJhZG1pbiJ9.mTHMQS_OQC1pbKTMecN0FrIFMxgRnWZzfRBMOoMNVDs"
}
```


### User APIs with admin privilege
As a admin, I can create/update/view list/view detail/delete contracts and staffs.

- Create new user:
```
POST /api/staffs/ (create)
```
Header authorization: `Bearer [token]`

Request:
```json
{
  "username": "newuser",
  "password": "12345",
  "is_admin": false
}
```

Response: status OK (200)

- Get user list:
```
GET /api/staffs/ (get list)
```
Header authorization: `Bearer [token]`

Request: None

Response: status OK (200)
```json
[
    {
        "uuid": "386069f6-72e1-4300-b7a4-a212e728ba5a",
        "username": "admin",
        "password": "12345",
        "is_admin": true,
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
    },
    {
        "uuid": "eddc9b4a-d9b5-4189-b291-93874d218805",
        "username": "staff",
        "password": "12345",
        "is_admin": false,
        "created_at": "2024-06-22T16:24:58.743+07:00",
        "updated_at": "2024-06-22T16:24:58.743+07:00"
    }
]
```

- Get user's details:
```
GET /api/staffs/<id>/ (get detail)
```

Header authorization: `Bearer [token]`

Request: `<id>` - user's uuid

Response: status OK (200)
```json
{
  "uuid": "eddc9b4a-d9b5-4189-b291-93874d218805",
  "username": "staff",
  "password": "12345",
  "is_admin": false,
  "created_at": "2024-06-22T16:24:58.743+07:00",
  "updated_at": "2024-06-22T16:24:58.743+07:00"
}
```

- Update user:
```
PUT /api/staffs/<id>/ (update)
```

Header authorization: `Bearer [token]`

Request: `<id>` - user's uuid
```json
{
    "username": "staff",
    "password": "newpass",
    "is_admin": true
}
```

Response: status OK (200)

- Delete user:
```
/api/staffs/<id>/ (delete)
```
Header authorization: `Bearer [token]`

Request: `<id>` - user's uuid

Response status: NoContent (204)
Response header: `"message": {"success": true}`



### Contract APIs for logged in user

As a staff, I can create/update/view list/view detail/delete contracts.

- Create new contract:
```
POST /api/contracts/ (create)
```
Header authorization: `Bearer [token]`

Request:
```json
{
  "username": "newuser",
  "password": "12345",
  "is_admin": false
}
```

Response: status OK (200)

- Get user's contracts list:
```
GET /api/contracts/ (get list)
```
Header authorization: `Bearer [token]`

Request: None

Response: status OK (200)
```json
[
  {
    "uuid": "4fb66a7c-f23a-4f4f-9b4e-528c77863df0",
    "name": "admincontract2",
    "user_uuid": "123456",
    "description": "alksjfsdfsaasslkf2453252jlskjflksjflksdaaaaaaaaaflskj",
    "created_at": "2024-06-22T16:30:57.345+07:00",
    "updated_at": "2024-06-22T16:30:57.345+07:00",
    "deleted_at": true
  },
  {
    "uuid": "7e431d60-e00c-43d2-af49-6ffaf807f4cc",
    "name": "admincontract",
    "user_uuid": "123456",
    "description": "alksjfsdfsslkfjlskjflksjflksdaaaaaaaaaflskj",
    "created_at": "2024-06-22T16:30:39.745+07:00",
    "updated_at": "2024-06-22T16:30:39.745+07:00"
  },
  {
    "uuid": "d314f579-bf2b-4121-b371-bc9f3b39c85e",
    "name": "admincontract1",
    "user_uuid": "123456",
    "description": "alksjfsdfsslkf2453252jlskjflksjflksdaaaaaaaaaflskj",
    "created_at": "2024-06-22T16:30:51.511+07:00",
    "updated_at": "2024-06-22T16:30:51.511+07:00"
  }
]
```

- Get user's contract details:
```
GET /api/contracts/<id>/ (get detail)
```

Header authorization: `Bearer [token]`

Request: `<id>` - contract's uuid

Response: status OK (200)
```json
{
  "uuid": "7e431d60-e00c-43d2-af49-6ffaf807f4cc",
  "name": "admincontract",
  "user_uuid": "123456",
  "description": "alksjfsdfsslkfjlskjflksjflksdaaaaaaaaaflskj",
  "created_at": "2024-06-22T16:30:39.745+07:00",
  "updated_at": "2024-06-22T16:30:39.745+07:00"
}
```

- Update contract:
```
PUT /api/contracts/<id>/ (update)
```

Header authorization: `Bearer [token]`

Request: `<id>` - contract's uuid
```json
{
  "name": "new contract",
  "description": "cccccccccccccc"
}
```

Response: status OK (200)

- Delete user:
```
/api/contracts/<id>/ (delete)
```
Header authorization: `Bearer [token]`

Request: `<id>` - contract's uuid

Response status: NoContent (204)
Response header: `"message": {"success": true}`


