# Players API

## List Players
Get all players with optional filters.

### Request
```http
GET /players

Query Parameters:
- department_id (optional): Filter by department
- is_active (optional): Filter by active status
- position (optional): Filter by position
```

### Response
```json
{
    "success": true,
    "data": {
        "players": [
            {
                "id": "integer",
                "department_id": "integer",
                "department_name": "string",
                "full_name": "string",
                "jersey_number": "integer",
                "position": "string",
                "date_of_birth": "date",
                "height_cm": "integer",
                "weight_kg": "integer",
                "phone": "string",
                "email": "string",
                "is_active": "boolean",
                "statistics": {
                    "total_matches": "integer",
                    "total_minutes_played": "integer",
                    "total_goals": "integer",
                    "total_assists": "integer"
                }
            }
        ]
    }
}
```

## Get Player
Get player by ID with detailed statistics.

### Request
```http
GET /players/{id}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "department_id": "integer",
        "department_name": "string",
        "full_name": "string",
        "jersey_number": "integer",
        "position": "string",
        "date_of_birth": "date",
        "height_cm": "integer",
        "weight_kg": "integer",
        "phone": "string",
        "email": "string",
        "is_active": "boolean",
        "statistics": {
            "total_matches": "integer",
            "total_minutes_played": "integer",
            "total_goals": "integer",
            "total_assists": "integer",
            "total_yellow_cards": "integer",
            "total_red_cards": "integer"
        },
        "recent_matches": [
            {
                "match_id": "integer",
                "match_date": "datetime",
                "opponent": "string",
                "minutes_played": "integer",
                "goals_scored": "integer",
                "assists": "integer"
            }
        ]
    }
}
```

## Create Player
Create new player.

### Request
```http
POST /players
Content-Type: application/json
Authorization: Bearer <token>

{
    "department_id": "integer",
    "full_name": "string",
    "jersey_number": "integer",
    "position": "string",
    "date_of_birth": "date",
    "height_cm": "integer",
    "weight_kg": "integer",
    "phone": "string",
    "email": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "department_id": "integer",
        "full_name": "string",
        "jersey_number": "integer",
        "position": "string",
        "date_of_birth": "date",
        "height_cm": "integer",
        "weight_kg": "integer",
        "phone": "string",
        "email": "string",
        "is_active": true
    }
}
```

## Update Player
Update existing player.

### Request
```http
PUT /players/{id}
Content-Type: application/json
Authorization: Bearer <token>

{
    "department_id": "integer",
    "full_name": "string",
    "jersey_number": "integer",
    "position": "string",
    "date_of_birth": "date",
    "height_cm": "integer",
    "weight_kg": "integer",
    "phone": "string",
    "email": "string",
    "is_active": "boolean"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "department_id": "integer",
        "full_name": "string",
        "jersey_number": "integer",
        "position": "string",
        "date_of_birth": "date",
        "height_cm": "integer",
        "weight_kg": "integer",
        "phone": "string",
        "email": "string",
        "is_active": "boolean"
    }
}
```

## Delete Player
Delete player by ID.

### Request
```http
DELETE /players/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "message": "Player deleted successfully"
}
```

## Get Player Statistics
Get detailed statistics for a player.

### Request
```http
GET /players/{id}/statistics
```

### Response
```json
{
    "success": true,
    "data": {
        "total_matches": "integer",
        "total_minutes_played": "integer",
        "total_goals": "integer",
        "total_assists": "integer",
        "total_yellow_cards": "integer",
        "total_red_cards": "integer",
        "matches_by_position": {
            "position": "count"
        },
        "goals_by_match": [
            {
                "match_date": "datetime",
                "opponent": "string",
                "goals": "integer"
            }
        ]
    }
}
``` 