# Go Weekly Database Golang

Perform integration for Tickitz database using golang technology and frameworks such as gin and implement MVC design patterns, perform database connections using pgx and also postman to help interaction between endpoints.

## How to Install

Make sure you already install Go to run this project

## How to Run this Project

1. Create a new empty directory for the project
2. Clone this project into the empty current directory:

```
git clone https://github.com/HN721/fgo24-be-Tickitz.git .
```

3. Install dependencies

```
go mod tidy
```

4. Run the project

```
go run main.go
```

## ERD Tickitz

```mermaid
erDiagram
direction LR
   users ||--o{ transactions : makes
    users ||--|| profile : has
    transactions ||--o{ transaction_detail : contains
    transactions ||--o{ history_transaction : contains
    transactions ||--o{ cinema : held_at
    transactions ||--o{ payment_method : paid_with
    movies ||--o{ transactions : involved_in

    movies ||--o{ movie_genre : categorized_as
    genres ||--o{ movie_genre : classifies

    movies ||--o{ movie_director : directed_by
    directors ||--o{ movie_director : directs

    movies ||--o{ movie_actors : includes
    actors ||--o{ movie_actors : acts_in
    users {
        string id PK
        string username
        string email
        string password
        string role
        timestamp created_at
    }
    profile{
        int id PK
        string fullname
         string phone_number
        string profile_image
        string id_user FK
    }
    movies {
        string id PK "index"
        string title
        string synopsis
        string background
        string poster
        timestamp release_date
        int duration
        int price

    }

    transactions {
        string id PK
        timestamp time
        date date_booking
        int price_total
        string id_cinema FK
        string id_payment_method FK
        string id_user FK
        string id_movie FK
    }
    transaction_detail{
        string id PK
        string id_transaction FK
        string costumer_name
        string costumer_phone
        string seat

    }

    directors {
        string id PK
        string name
    }

    genres {
        string id PK
        string name
    }

    actors {
        string id PK
        string name
    }

    movie_actors {
        string id PK
        string id_movie FK
        string id_actor FK
    }
      movie_genre {
        string id PK
        string id_movie FK
        string id_genre FK
    }
      movie_director {
        string id PK
        string id_movie FK
        string id_director FK
    }
     history_transaction{   string id PK
        string id_transaction FK
        string status
        string note
    }
    cinema{
        string ID PK
        string name
    }
    payment_method{
        string ID PK
        string name
    }

```

## Depedencies

- Gin Gonic (github.com/gin-gonic/gin)
- JWT V5 (github.com/golang-jwt/jwt/v5)
- PGX (github.com/jackc/pgx/v5)
- Godotenv (github.com/joho/godotenv)

## Contributing

We welcome contributions! 🚀

If you would like to open a Pull Request (PR), please follow these steps:

1. Fork this repository.

2. Create a new branch:

```
git checkout -b feature/your-feature-name
```

3. Make your changes.

4. Commit your changes with a clear message:

```
git commit -m "Add: your featusre description"
Push your branch to your fork:
```

```
git push origin feature/your-feature-name
```

5. Open a Pull Request (PR) to the main branch of this repository.

6. Make sure your code is clean, well-tested, and consistent with the project style.

7. License
   This project is open-source and available under the MIT License.
