# Team Fees API

## List Team Fees
Get all team maintenance fees with optional filters.

### Request
```http
GET /team-fees

Query Parameters:
- start_date (optional): Filter by start date
- end_date (optional): Filter by end date
```

### Response
```json
{
    "success": true,
    "data": {
        "fees": [
            {
                "id": "integer",
                "amount": "decimal",
                "payment_date": "date",
                "description": "string",
                "created_at": "datetime",
                "updated_at": "datetime"
            }
        ],
        "summary": {
            "total_amount": "decimal",
            "total_payments": "integer",
            "average_amount": "decimal"
        }
    }
}
```

## Get Team Fee
Get team fee by ID.

### Request
```http
GET /team-fees/{id}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "amount": "decimal",
        "payment_date": "date",
        "description": "string",
        "created_at": "datetime",
        "updated_at": "datetime"
    }
}
```

## Create Team Fee
Create new team maintenance fee.

### Request
```http
POST /team-fees
Content-Type: application/json
Authorization: Bearer <token>

{
    "amount": "decimal",
    "payment_date": "date",
    "description": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "amount": "decimal",
        "payment_date": "date",
        "description": "string",
        "created_at": "datetime",
        "updated_at": "datetime"
    }
}
```

## Update Team Fee
Update existing team fee.

### Request
```http
PUT /team-fees/{id}
Content-Type: application/json
Authorization: Bearer <token>

{
    "amount": "decimal",
    "payment_date": "date",
    "description": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "amount": "decimal",
        "payment_date": "date",
        "description": "string",
        "created_at": "datetime",
        "updated_at": "datetime"
    }
}
```

## Delete Team Fee
Delete team fee by ID.

### Request
```http
DELETE /team-fees/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "message": "Team fee deleted successfully"
}
```

## Get Team Fees Statistics
Get detailed statistics about team fees.

### Request
```http
GET /team-fees/statistics

Query Parameters:
- year (optional): Filter by year
```

### Response
```json
{
    "success": true,
    "data": {
        "summary": {
            "total_amount": "decimal",
            "total_payments": "integer",
            "average_amount": "decimal"
        },
        "monthly_summary": [
            {
                "month": "string",
                "total_amount": "decimal",
                "number_of_payments": "integer"
            }
        ],
        "yearly_summary": [
            {
                "year": "integer",
                "total_amount": "decimal",
                "number_of_payments": "integer"
            }
        ]
    }
}
``` 