# Blog Post API

This is a simple blog post API written in Golang using the Gorilla Mux framework. The API provides basic CRUD operations for managing blog posts. It runs on port 9098.

## Setup

1. Clone the repository:

```
git clone <repository-url>
```

2. Install Golang (if not already installed) from the official Golang website: https://golang.org/

3. Install project dependencies:

```
go mod download
```

4. Configure the environment variables:

   - Create a `.env` file in the project root directory.
   - Add the following environment variable to the `.env` file:

   ```plaintext
   PORT=9098
   ```

5. Run the application:

```
go run main.go
```

The application will start running on port 9098 as specified in the `PORT` environment variable.

## API Endpoints

The API provides the following endpoints for managing blog posts:

- **GET /posts**

  Fetches all blog posts.

- **GET /posts/{id}**

  Fetches a single blog post by ID.

- **POST /posts**

  Creates a new blog post.

- **PUT /posts/{id}**

  Updates a blog post by ID.

- **DELETE /posts/{id}**

  Deletes a blog post by ID.

## Request and Response Formats

### Request Format

For creating and updating blog posts, the request body should be in JSON format and include the following fields:

```json
{
  "title": "My Blog Post",
  "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam ac consectetur elit."
}
```

### Response Format

The API responds with JSON-formatted data for most requests. For example, when fetching all blog posts, the response will be an array of blog post objects:

```json
[
  {
    "id": 1,
    "title": "My Blog Post",
    "desc":"Description of the post",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam ac consectetur elit."  
  },
  {
    "id": 2,
    "title": "Another Blog Post",
    "desc": "Another description of the blog post",
    "content": "Sed vel eros eget nisi vestibulum porttitor vitae sit amet nunc."
  }
]
```

## Database Configuration

This API does not use any database by default. But you can create the database file (`blog.db`) in the project root directory when the application runs.

## Contributing

Contributions are welcome! If you find any issues or want to add new features, please create a pull request with a detailed description of the changes.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as per the terms of the license.