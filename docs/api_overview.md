# API Documentation - Company Football Team Management System

## Overview
This API provides endpoints for managing a company football team, including players, matches, departments, and team maintenance.

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
All endpoints except login/register require JWT authentication.
Include the token in the Authorization header:
```
Authorization: Bearer <token>
```

## Common Response Format
```json
{
    "success": true,
    "data": {},
    "message": "Operation successful"
}
```

## Error Response Format
```json
{
    "success": false,
    "error": {
        "code": "ERROR_CODE",
        "message": "Error description"
    }
}
```

## Rate Limiting
- 100 requests per minute per IP
- 1000 requests per hour per user

## API Versioning
Current version: v1
Version is included in the URL path: `/api/v1/...` 