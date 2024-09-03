# Annotation API

<div align="center">
  <br>
  <a href="https://echo.labstack.com"><img height="80" src="https://cdn.labstack.com/images/echo-logo.svg" alt=""></a>
  <br>
</div>

## Table of Contents

1. [Overview](#overview)
2. [Endpoints](#endpoints)
    - [Retrieve All Annotations](#1-retrieve-all-annotations)
    - [Create a New Annotation](#2-create-a-new-annotation)
    - [Update an Existing Annotation](#3-update-an-existing-annotation)
    - [Delete an Annotation](#4-delete-an-annotation)
3. [Request and Response Formats](#request-and-response-formats)
4. [Error Handling](#error-handling)
5. [Examples](#examples)
    - [Retrieve All Annotations](#retrieve-all-annotations-example)
    - [Create a New Annotation](#create-a-new-annotation-example)
    - [Update an Existing Annotation](#update-an-existing-annotation-example)
    - [Delete an Annotation](#delete-an-annotation-example)

## Overview

The Annotation API allows you to manage annotations in your system. You can perform operations to retrieve, create, update, and delete annotations. Each annotation includes text (in Arabic) and optional metadata.

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



