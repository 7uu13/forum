# GoForum Web Application

GoForum is a web application that allows users to engage in discussions, create and categorize posts, like or dislike content, and apply various filters.

## Table of Contents
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [User Registration](#user-registration)
  - [Communication](#communication)
  - [Likes and Dislikes](#likes-and-dislikes)
  - [Filtering](#filtering)
- [Docker](#docker)
- [Contributing](#contributing)

## Features

GoForum offers the following key features:

- **User Registration:**
  - Users can register by providing their email, username, and password.
  - Duplicate email registrations are handled gracefully.

- **Communication:**
  - Registered users can create posts and comments.
  - Posts can be associated with one or more categories.
  - Posts and comments are visible to all users, even non-registered ones.

- **Likes and Dislikes:**
  - Only registered users can like or dislike posts and comments.
  - The number of likes and dislikes is visible to all users.

- **Filtering:**
  - A filtering mechanism allows users to:
    - Filter posts by categories.
    - Filter posts created by them.
    - Filter posts they've liked.

- **Database:**
  - GoForum uses SQLite as the database, ensuring efficient data storage and retrieval.

- **Docker:**
  - The application is containerized using Docker for easy deployment and management.
  - Use Docker to containerize the application.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following prerequisites:

- Go: Install the Go programming language from https://golang.org/dl/
- Docker: Install Docker: https://docs.docker.com/get-docker/

### Installation

1. Clone the GoForum repository:
   ```sh
   git clone https://01.kood.tech/git/Karl-Thomas/forum.git
   cd forum
   ```

2. Build and run the GoForum web application:
   ```sh
   go run .
   ```

GoForum should now be running locally at [http://localhost:8080](http://localhost:8080).

## Usage

### User Registration

1. Access the GoForum application at [http://localhost:8080](http://localhost:8080).

2. Register as a new user by providing your email, username, and password.

3. If the email is already registered, you will receive an error message.

### Communication

1. Registered users can create posts and comments.

2. Posts can be categorized and will be visible to all users, registered or not.

### Likes and Dislikes

1. Registered users can like or dislike posts and comments.

2. The number of likes and dislikes is visible to all users.

### Filtering

1. Use the filtering mechanism to:
   - Filter posts by categories.
   - Filter posts created by you.
   - Filter posts you've liked.

## Docker

GoForum can be containerized using Docker. Refer to the Docker documentation for more information on containerizing and deploying the application.

1. Open the project folder and run this command: 
```sh
docker build -t forum .    
```

2. Then run this command to start server on port 8080:
```sh
docker run -p 8080:8080 forum 
```
3. Open browser and write localhost:8080 to see if webpage and docker works.
```sh
docker run -p 8080:8080 forum 
```
## Contributing

- Kveber
- Karl-Thomas
- Ktuule



