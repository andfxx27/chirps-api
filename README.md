# chirps-api

REST API for chatting app chirps

## Data Dictionary

### User

| Field      | Data Type    |
| ---------- | ------------ |
| id         | serial       |
| first_name | varchar(100) |
| last_name  | varchar(100) |
| username   | varchar(100) |
| email      | varchar(255) |
| password   | varchar(255) |
| status     | varchar(255) |
| created_at | timestamp    |
| updated_at | timestamp    |
| deleted_at | timestamp    |

### Follow

| Field       | Data Type    |
| ----------- | ------------ |
| follower_id | varchar(255) |
| followed_id | varchar(255) |
| created_at  | timestamp    |
| updated_at  | timestamp    |
| deleted_at  | timestamp    |

To view the ERD, follow [this](https://drive.google.com/file/d/1BDLShHIfMx1AXGY9u8bqFa0Pf4yNrxQ8/view?usp=sharing) link.

## Important Notes

## Todo

Update this README to include everything needed for this project to get up and running.
