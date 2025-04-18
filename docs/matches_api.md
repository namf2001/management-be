# Matches API

## List Matches
Get all matches with optional filters.

### Request
```http
GET /matches
Authorization: Bearer <token>

Query Parameters:
- status (optional): Filter by match status (scheduled, completed, cancelled)
- start_date (optional): Filter by start date
- end_date (optional): Filter by end date
- opponent_team_id (optional): Filter by opponent team
```

### Response
```json
{
    "success": true,
    "data": {
        "matches": [
            {
                "id": "integer",
                "opponent_team": {
                    "id": "integer",
                    "name": "string",
                    "company_name": "string"
                },
                "match_date": "datetime",
                "venue": "string",
                "is_home_game": "boolean",
                "our_score": "integer",
                "opponent_score": "integer",
                "status": "string",
                "notes": "string"
            }
        ]
    }
}
```

## Get Match
Get match by ID with detailed information.

### Request
```http
GET /matches/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "opponent_team": {
            "id": "integer",
            "name": "string",
            "company_name": "string",
            "contact_person": "string",
            "contact_phone": "string"
        },
        "match_date": "datetime",
        "venue": "string",
        "is_home_game": "boolean",
        "our_score": "integer",
        "opponent_score": "integer",
        "status": "string",
        "notes": "string",
        "players": [
            {
                "player_id": "integer",
                "player_name": "string",
                "jersey_number": "integer",
                "position": "string",
                "minutes_played": "integer",
                "goals_scored": "integer",
                "assists": "integer",
                "yellow_cards": "integer",
                "red_card": "boolean"
            }
        ]
    }
}
```

## Create Match
Create new match.

### Request
```http
POST /matches
Content-Type: application/json
Authorization: Bearer <token>

{
    "opponent_team_id": "integer",
    "match_date": "datetime",
    "venue": "string",
    "is_home_game": "boolean",
    "notes": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "opponent_team_id": "integer",
        "match_date": "datetime",
        "venue": "string",
        "is_home_game": "boolean",
        "status": "scheduled",
        "notes": "string"
    }
}
```

## Update Match
Update existing match.

### Request
```http
PUT /matches/{id}
Content-Type: application/json
Authorization: Bearer <token>

{
    "opponent_team_id": "integer",
    "match_date": "datetime",
    "venue": "string",
    "is_home_game": "boolean",
    "our_score": "integer",
    "opponent_score": "integer",
    "status": "string",
    "notes": "string"
}
```

### Response
```json
{
    "success": true,
    "data": {
        "id": "integer",
        "opponent_team_id": "integer",
        "match_date": "datetime",
        "venue": "string",
        "is_home_game": "boolean",
        "our_score": "integer",
        "opponent_score": "integer",
        "status": "string",
        "notes": "string"
    }
}
```

## Delete Match
Delete match by ID.

### Request
```http
DELETE /matches/{id}
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "message": "Match deleted successfully"
}
```

## Update Match Players
Update player participation in a match.

### Request
```http
PUT /matches/{id}/players
Content-Type: application/json
Authorization: Bearer <token>

{
    "players": [
        {
            "player_id": "integer",
            "minutes_played": "integer",
            "goals_scored": "integer",
            "assists": "integer",
            "yellow_cards": "integer",
            "red_card": "boolean"
        }
    ]
}
```

### Response
```json
{
    "success": true,
    "data": {
        "match_id": "integer",
        "players": [
            {
                "player_id": "integer",
                "player_name": "string",
                "minutes_played": "integer",
                "goals_scored": "integer",
                "assists": "integer",
                "yellow_cards": "integer",
                "red_card": "boolean"
            }
        ]
    }
}
```

## Get Match Statistics
Get match statistics and summary.

### Request
```http
GET /matches/{id}/statistics
Authorization: Bearer <token>
```

### Response
```json
{
    "success": true,
    "data": {
        "match_summary": {
            "total_players": "integer",
            "total_minutes_played": "integer",
            "total_goals": "integer",
            "total_assists": "integer",
            "total_yellow_cards": "integer",
            "total_red_cards": "integer"
        },
        "player_performance": [
            {
                "player_id": "integer",
                "player_name": "string",
                "position": "string",
                "minutes_played": "integer",
                "goals_scored": "integer",
                "assists": "integer",
                "yellow_cards": "integer",
                "red_card": "boolean"
            }
        ],
        "position_summary": {
            "position": "count"
        }
    }
}
``` 