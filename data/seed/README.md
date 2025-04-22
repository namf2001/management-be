# Database Seeding

This directory contains scripts to seed the database with fake data for development and testing purposes.

## Overview

The seed scripts generate fake data for the following entities:
- Departments
- Users
- Teams
- Players
- Matches
- Match Players
- Team Fees
- Player Statistics

## Prerequisites

Before running the seed scripts, you need to install the required dependencies:

```bash
go get github.com/bxcodec/faker/v3
```

## Running the Seed Scripts

To seed the database with fake data, run the following command from the project root:

```bash
go run data/seed/main.go
```

This will execute all the seed functions in the correct order to respect foreign key constraints.

## Customization

You can modify the seed files to change the amount or type of data generated:

- `departments.go`: Modify the department names or add more departments
- `users.go`: Change the number of users or password generation
- `teams.go`: Adjust the number of teams or team attributes
- `players.go`: Change the number of players or player attributes
- `matches.go`: Modify the number of matches or match attributes
- `match_players.go`: Adjust how players are assigned to matches
- `team_fees.go`: Change the number or amount of team fees
- `player_statistics.go`: Adjust the statistics generation logic

## Notes

- The seed scripts are designed to be idempotent, meaning they can be run multiple times without causing duplicate data issues.
- Some entities have unique constraints, so running the scripts multiple times may result in some entities being skipped if they would violate these constraints.