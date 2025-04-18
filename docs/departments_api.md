# Departments API

## List Departments
Get all departments.

### Request
```http
GET /departments
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "data": {
        "departments": [
            {
                "id": "integer",
                "name": "string",
                "description": "string",
                "created_at": "datetime",
                "updated_at": "datetime"
            }
        ]
    }
}
```

## Get Department
Get department by ID.

### Request
```http
GET /departments/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "name": "string",
        "description": "string",
        "created_at": "datetime",
        "updated_at": "datetime"
    }
}
```

## Create Department
Create new department.

### Request
```http
POST /departments
Content-Type: application/json
Authorization: Bearer <token>

{
    "name": "string",
    "description": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "name": "string",
        "description": "string",
        "created_at": "datetime",
        "updated_at": "datetime"
    }
}
```

## Update Department
Update existing department.

### Request
```http
PUT /departments/{id}
Content-Type: application/json
Authorization: Bearer <token>

{
    "name": "string",
    "description": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "name": "string",
        "description": "string",
        "created_at": "datetime",
        "updated_at": "datetime"
    }
}
```

## Delete Department
Delete department by ID.

### Request
```http
DELETE /departments/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "message": "Department deleted successfully"
}
``` 