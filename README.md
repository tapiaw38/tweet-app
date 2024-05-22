# Tweet App

## Description

Tweet App is a microblogging application that allows users to post and share tweets. The app is built following the hexagonal architecture, which ensures a clear separation between business logic and technical details such as data persistence and user interface.

## Architecture

The hexagonal architecture, also known as ports and adapters, is used to create a highly decoupled and maintainable application. In Tweet App, the architecture is divided into the following layers:

- `domain`: Contains business entities and business logic.
- `adapters`: Contains the adapters used to connect the application with external services.
- `usecases`: Contains the application's use cases, representing the operations that users can perform.

## Running with Docker

To run the application with Docker, you first need to build the Docker image with the following command:

```sh
docker-compose build
```

Then, you can start the application with the following command:

```sh
docker-compose up
```

## Endpoints

The application provides the following endpoints:

### Create Users Request

**Method**: POST
**Endpoint**: /api/users
**Body**:

```json
{
  "first_name": "User",
  "last_name": "One",
  "username": "user1",
  "email": "user1@example.com",
  "password": "password"
}
```

### Create Users Response

```json
{
  "data": {
    "id": 1,
    "first_name": "User",
    "last_name": "One",
    "username": "user1",
    "email": "user1@example.com",
    "created_at": "2021-10-10T00:00:00Z",
    "updated_at": "2021-10-10T00:00:00Z"
  }
}
```

### Create Follow Request

**Method**: POST
**Endpoint**: /api/users/follow/{user_id}
**Body**:

```json
{
  "follow_id": 2
}
```

### Create Follow Response

```json
{
  "message": "User followed successfully"
}
```

### Create Tweets Request

**Method**: POST
**Endpoint**: /api/tweets
**Body**:

```json
{
  "user_id": 1,
  "content": "This is a tweet"
}
```

### Create Tweets Response

```json
{
  "data": {
    "id": 1,
    "user_id": 1,
    "content": "This is a tweet",
    "created_at": "2021-10-10T00:00:00Z",
    "updated_at": "2021-10-10T00:00:00Z"
  }
}
```

### List Tweets Request

**Method**: GET
**Endpoint**: /api/tweets

The request might be done using the following query parameters:

| Query      | Description                                                                 | Example |
| :--------- | :-------------------------------------------------------------------------- | :------ |
| `username` | _Filter the tweets and get those of those followed by the user's username._ | `user1` |

### List Tweets Response

```json
{
  "data": [
    {
      "id": 1,
      "user_id": 1,
      "content": "This is a tweet",
      "created_at": "2021-10-10T00:00:00Z",
      "updated_at": "2021-10-10T00:00:00Z"
    }
  ]
}
```
