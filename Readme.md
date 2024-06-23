# Draft instruction


### How to run:
1. Create docker database service: `docker-compose up -d`
2. Attach database container and create tables from `db/init.sql`
3. Run program `go run main.go`


### APIs

#### Login
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


#### User APIs with admin privilege
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


#### Contract APIs for logged in user

	//- As a staff, I can create/update/view list/view detail/delete contracts.
	//* POST /api/contracts/ (create)
	//* GET /api/contracts/ (get list)
	//* PATCH/PUT /api/contracts/<id>/ (update)
	//* GET /api/contracts/<id>/ (get detail)
	//* DELETE /api/contracts/<id>/ (delete)