# Teams API

## List Teams
Get all opponent teams.

### Request
```http
GET /teams/page={page_number}&limit={page_size}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "data": {
        "teams": [
            {
                "id": "integer",
                "name": "string",
                "company_name": "string",
                "contact_person": "string",
                "contact_phone": "string",
                "contact_email": "string",
                "match_history": {
                    "total_matches": "integer",
                    "wins": "integer",
                    "losses": "integer",
                    "draws": "integer"
                }
            }
        ],
        "pagination": {
            "current_page": "integer",
            "total_pages": "integer",
            "total_items": "integer",
            "items_per_page": "integer"
        }
    }
}
```

## Get Team
Get team by ID with match history.

### Request
```http
GET /teams/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "name": "string",
        "company_name": "string",
        "contact_person": "string",
        "contact_phone": "string",
        "contact_email": "string",
        "match_history": {
            "total_matches": "integer",
            "wins": "integer",
            "losses": "integer",
            "draws": "integer",
            "matches": [
                {
                    "match_id": "integer",
                    "match_date": "datetime",
                    "venue": "string",
                    "our_score": "integer",
                    "opponent_score": "integer",
                    "status": "string"
                }
            ]
        }
    }
}
```

## Create Team
Create new opponent team.

### Request
```http
POST /teams
Content-Type: application/json
Authorization: Bearer <token>

{
    "name": "string",
    "company_name": "string",
    "contact_person": "string",
    "contact_phone": "string",
    "contact_email": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "name": "string",
        "company_name": "string",
        "contact_person": "string",
        "contact_phone": "string",
        "contact_email": "string"
    }
}
```

## Update Team
Update existing team.

### Request
```http
PUT /teams/{id}
Content-Type: application/json
Authorization: Bearer <token>

{
    "name": "string",
    "company_name": "string",
    "contact_person": "string",
    "contact_phone": "string",
    "contact_email": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "name": "string",
        "company_name": "string",
        "contact_person": "string",
        "contact_phone": "string",
        "contact_email": "string"
    }
}
```

## Delete Team
Delete team by ID.

### Request
```http
DELETE /teams/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "message": "Team deleted successfully"
}
```

## Get Team Statistics
Get detailed statistics for a team.

### Request
```http
GET /teams/{id}/statistics
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "data": {
        "match_summary": {
            "total_matches": "integer",
            "wins": "integer",
            "losses": "integer",
            "draws": "integer",
            "goals_scored": "integer",
            "goals_conceded": "integer"
        },
        "recent_form": [
            {
                "match_date": "datetime",
                "result": "string",
                "score": "string"
            }
        ],
        "performance_by_venue": {
            "home": {
                "matches": "integer",
                "wins": "integer",
                "losses": "integer",
                "draws": "integer"
            },
            "away": {
                "matches": "integer",
                "wins": "integer",
                "losses": "integer",
                "draws": "integer"
            }
        }
    }
}
``` 