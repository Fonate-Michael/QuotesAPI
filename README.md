# QuotesAPI

A RESTful API for managing and retrieving quotes, built with Go and Gin framework. This API allows you to perform CRUD operations on quotes and manage comments on them.

## Features

- **Quote Management**
  - Get all quotes with pagination
  - Get a specific quote by ID
  - Search quotes by content
  - Get a random quote
  - Add new quotes
  - Update existing quotes
  - Delete quotes

- **Comment System**
  - Get all comments for a specific quote
  - Add comments to quotes

## Tech Stack

- **Backend**: Go (Golang)
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: Standard library `database/sql`

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Fonate-Michael/QuotesAPI.git
   cd QuotesAPI
   ```

2. Set up environment variables:
   Create a `.env` file in the root directory with the following variables:
   ```
   DB_HOST=your_db_host
   DB_PORT=your_db_port
   DB_USER=your_db_username
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run database migrations:
   ```sql
   CREATE TABLE IF NOT EXISTS quotes (
       id SERIAL PRIMARY KEY,
       message TEXT NOT NULL,
       author VARCHAR(255) NOT NULL
   );

   CREATE TABLE IF NOT EXISTS comments (
       id SERIAL PRIMARY KEY,
       quote_id INTEGER REFERENCES quotes(id) ON DELETE CASCADE,
       user_id INTEGER NOT NULL,
       comment TEXT NOT NULL
   );
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`

## API Endpoints

### Quotes

- **GET /quotes** - Get all quotes
  - Query Parameters:
    - `page` - Page number (default: 1)
    - `limit` - Number of items per page (default: 7)

- **GET /quotes/:id** - Get a specific quote by ID

- **GET /quotes/random** - Get a random quote

- **GET /quotes/search** - Search quotes
  - Query Parameters:
    - `q` - Search query (required)

- **POST /quotes** - Add a new quote
  - Request Body:
    ```json
    {
        "message": "Sample quote",
        "author": "Author Name"
    }
    ```

- **PUT /quotes/:id** - Update a quote
  - Request Body:
    ```json
    {
        "message": "Updated quote",
        "author": "Updated Author"
    }
    ```

- **DELETE /quotes/:id** - Delete a quote

### Comments

- **GET /quotes/:id/comments** - Get all comments for a quote

- **POST /quotes/:id/comments** - Add a comment to a quote
  - Request Body:
    ```json
    {
        "user_id": 1,
        "comment": "This is a great quote!"
    }
    ```

## Example Usage

### Get all quotes (paginated)
```bash
curl -X GET "http://localhost:8080/quotes?page=1&limit=5"
```

### Add a new quote
```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"message":"To be or not to be","author":"William Shakespeare"}'
```

### Add a comment to a quote
```bash
curl -X POST http://localhost:8080/quotes/1/comments \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "comment": "This is a classic!"}'
```

## Error Handling

The API returns appropriate HTTP status codes along with error messages in JSON format:

```json
{
    "error": "Quote not found"
}
```

## CORS

CORS is enabled with the following configuration:
- Allowed Origins: *
- Allowed Methods: GET, POST, PUT, PATCH, DELETE
- Allowed Headers: Origin, Content-Type, Authorization, Accept
- Exposed Headers: Content-Length
- Credentials: true

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Gin Web Framework](https://github.com/gin-gonic/gin)

