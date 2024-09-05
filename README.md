# Annotation API

<div align="center">
  <br>
  <div style="display: flex; justify-content: center; gap: 20px;">
    <a href="https://echo.labstack.com">
      <img height="60" src="https://cdn.labstack.com/images/echo-logo.svg" alt="Echo Logo">
    </a>
    <a href="https://www.mysql.com">
      <img height="80" src="https://www.mysql.com/common/logos/logo-mysql-170x115.png" alt="MySQL Logo">
    </a>
    <a href="https://golang.org">
      <img height="100" src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Logo">
    </a>
  </div>
  <br>
</div>

## Table of Contents

1. [Overview](#overview)
2. [Technologies User](#technologies-used)
3. [Setup Instructions](#setup-instructions)
    - [Prerequisites](#prerequisites)
    - [Installing Dependencies](#installing-dependencies)
    - [Running the server](#running-the-server-)
4. [Endpoints](#endpoints)
    - [Retrieve All Annotations](#1-retrieve-all-annotations)
    - [Create a New Annotation](#2-create-a-new-annotation)
    - [Update an Existing Annotation](#3-update-an-existing-annotation)
    - [Delete an Annotation](#4-delete-an-annotation)
5. [Request and Response Formats](#request-and-response-formats)
6. [Error Handling](#error-handling)
7. [Examples](#examples)
    - [Retrieve All Annotations](#retrieve-all-annotations-example)
    - [Create a New Annotation](#create-a-new-annotation-example)
    - [Update an Existing Annotation](#update-an-existing-annotation-example)
    - [Delete an Annotation](#delete-an-annotation-example)

## Overview

The Annotation API allows you to manage annotations in your system. You can perform operations to retrieve, create, update, and delete annotations. Each annotation includes text (in Arabic) and optional metadata.

## Technologies Used

- **Go**: The programming language used for the project.
- **Echo**: A high-performance, extensible web framework for Go. [Echo Documentation](https://echo.labstack.com/)
- **MySQL**: A relational database management system used to store and manage data.

## Setup Instructions

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version X.X or later)
- [MySQL](https://dev.mysql.com/downloads/) (version X.X or later)

### Installing Dependencies

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/your-repo.git
   cd your-repo
   ```

2. **Install Go Dependencies**

    Use Go modules to install the project dependencies:

    ```bash
    go mod download
    ```
3. **SetUp MySQL** 

    - Ensure that MySQL is running on your machine or accessible remotely.
    - Login into Mysql: 
      ```bash
      mysql -u your_username -p
      ```
    - Select your database name:
      ```bash
      USE your_database_name;
      ```
    - Navigate to the root of the repository and run the following command which will in turn create the required table for the project:
      ```bash
      SOURCE table-creation.sql
      ```
    - Replace config/config.json with your values.

### Running the server 

To run the server run the following command in the root of the repository 

```bash
go run cmd/api/main.go
```
   

## Endpoints

### 1. Retrieve All Annotations

- **Endpoint:** `GET /annotations`
- **Description:** Retrieve a list of all annotations.
- **Response:**
    - **Status Code:** 200 OK
    - **Content-Type:** `application/json`
    - **Body:**
      ```json
      [
        {
          "id": "int",
          "text": "string",
          "metadata": "object"
        }
      ]
      ```
    - **Possible Errors:**
        - **500 Internal Server Error:** An unexpected server error occurred.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```

### 2. Create a New Annotation

- **Endpoint:** `POST /annotations`
- **Description:** Create a new annotation with the specified text and optional metadata.
- **Request Body:**
    - **Content-Type:** `application/json`
    - **Body:**
      ```json
      {
        "text": "string (Arabic text)",
        "metadata": "object (optional)"
      }
      ```
- **Response:**
    - **Status Code:** 201 Created
    - **Content-Type:** `application/json`
    - **Body:**
      ```json
      {
        "id": "string",
        "text": "string",
        "metadata": "object"
      }
      ```
    - **Possible Errors:**
        - **400 Bad Request:** Invalid input data or missing fields.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```
        - **500 Internal Server Error:** An unexpected server error occurred.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```

### 3. Update an Existing Annotation

- **Endpoint:** `PUT /annotations/:id`
- **Description:** Update the annotation with the specified `id`.
- **Request Body:**
    - **Content-Type:** `application/json`
    - **Body:**
      ```json
      {
        "text": "string",
        "metadata": "object (optional)"
      }
      ```
- **Response:**
    - **Status Code:** 200 OK
    - **Content-Type:** `application/json`
    - **Body:**
      ```json
      {
        "id": "string",
        "text": "string",
        "metadata": "object"
      }
      ```
    - **Possible Errors:**
        - **400 Bad Request:** Invalid input data or missing fields.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```
        - **404 Not Found:** The specified `id` does not exist.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```
        - **500 Internal Server Error:** An unexpected server error occurred.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```

### 4. Delete an Annotation

- **Endpoint:** `DELETE /annotations/:id`
- **Description:** Delete the annotation with the specified `id`.
- **Response:**
    - **Status Code:** 200 OK
    - **Content-Type:** `application/json`
    - **Body:**
      ```json
      {
        "success": true,
        "message": "string (optional)"
      }
      ```
    - **Possible Errors:**
        - **404 Not Found:** The specified `id` does not exist.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```
        - **500 Internal Server Error:** An unexpected server error occurred.
            - **Response:**
              ```json
              {
                "message": "string"
              }
              ```

## Request and Response Formats

- **Request Content-Type:** `application/json`
- **Response Content-Type:** `application/json`

## Error Handling

The API responds with appropriate HTTP status codes and messages for errors:

- **400 Bad Request:** Invalid input data or missing fields.
- **404 Not Found:** The specified resource does not exist.
- **500 Internal Server Error:** An unexpected server error occurred.

## Examples

### Retrieve All Annotations Example

**Request:**

```http
GET /annotations HTTP/1.1
Host: api.example.com
```

**Response:**

```http
[
  {
    "id": "10",
    "text": "الولد",
    "metadata": {
        "comments": "this is a comment"
    }
  }
]
```

### Create a New Annotation Example

**Request:**

```http
POST /annotations HTTP/1.1
Host: api.example.com
Content-Type: application/json

{
  "text": "الولد",
  "metadata": {
    "comments": "this is a comment"
  }
}
```

**Response:**

```http
[
  {
    "id": "10",
    "text": "الولد",
    "metadata": {
        "comments": "this is a comment"
    }
  }
]
```

### Update an Existing Annotation Example

**Request:**

```http
POST /annotations/10 HTTP/1.1
Host: api.example.com
Content-Type: application/json

{
  "text": "الولد البسيط",
  "metadata": {
    "comments": "this is a comment"
  }
}
```

**Response:**

```http
{
  "id": 10, 
  "text": "الولد البسيط",
  "metadata": {
    "comments": "this is a comment"
  }
}
```

### Delete an Annotation Example

**Request:**

```http
DELETE /annotations/10 HTTP/1.1
Host: api.example.com
```

**Response:**

```http
{
    "success": true
}
```



