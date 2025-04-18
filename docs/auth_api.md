# Authentication API

## Login
Authenticate user and get JWT token.

### Request
```http
POST /auth/login
Content-Type: application/json

{
    "username": "string",
    "password": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "token": "string",
        "user": {
            "id": "integer",
            "username": "string",
            "email": "string",
            "full_name": "string"
        }
    }
}
```

## Register
Register new admin user.

### Request
```http
POST /auth/register
Content-Type: application/json

{
    "username": "string",
    "password": "string",
    "email": "string",
    "full_name": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "username": "string",
        "email": "string",
        "full_name": "string"
    }
}
```

## Change Password
Change user password.

### Request
```http
POST /auth/change-password
Content-Type: application/json
Authorization: Bearer <token>

{
    "current_password": "string",
    "new_password": "string"
}
```

### Response
```json
{
    "success": true,
    "message": "Password changed successfully"
} 