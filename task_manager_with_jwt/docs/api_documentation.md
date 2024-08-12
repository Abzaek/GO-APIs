# Task Management API

The Task Management API provides a comprehensive set of tools and resources that enable you to efficiently manage tasks within your application. This API is designed to support various operations, including creating, retrieving, updating, and deleting tasks, all while ensuring secure access control.

## Features

- **Task Management**: Perform CRUD (Create, Read, Update, Delete) operations on tasks.
- **Authorization Control**: Secure access to task management functionality based on user roles.

## Endpoints

### 1. Get All Tasks
- **Endpoint**: `/tasks`
- **Method**: `GET`
- **Description**: Retrieves a list of all tasks.
- **Response**:
  - A JSON array containing objects with properties like `id`, `title`, `description`, `due_date`, and `status`.

#### Example Request
```http
GET localhost:3000/tasks
```

#### Example Response
```json
[
  {
    "id": "1",
    "title": "Task 1",
    "description": "Description of Task 1",
    "due_date": "2024-08-16T16:00:00.000000Z",
    "status": "pending"
  },
  {
    "id": "2",
    "title": "Task 2",
    "description": "Description of Task 2",
    "due_date": "2024-08-17T16:00:00.000000Z",
    "status": "completed"
  }
]
```

### 2. Get Task by ID
- **Endpoint**: `/tasks/:id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific task by its ID.
- **Response**:
  - A JSON object representing the task with fields such as `id`, `title`, `description`, `due_date`, and `status`.

#### Example Request
```http
GET localhost:3000/tasks/2
```

#### Example Response
```json
{
  "id": "2",
  "title": "Task 2",
  "description": "Description of Task 2",
  "due_date": "2024-08-17T16:00:00.000000Z",
  "status": "completed"
}
```

### 3. Create a New Task
- **Endpoint**: `/tasks`
- **Method**: `POST`
- **Authorization**: **Admin Only**
- **Description**: Allows an admin user to create a new task.
- **Request Body**:
  - `id` (string): The unique identifier for the task.
  - `title` (string): The title of the task.
  - `description` (string): The description of the task.
  - `dueDate` (string): The due date of the task.
  - `status` (string): The status of the task.
- **Response**:
  - Status: `201 Created`
  - Message: `"Successfully created"`

#### Example Request
```http
POST localhost:3000/tasks
Content-Type: application/json

{
    "id": "10",
    "title": "Task 10",
    "description": "Tenth task",
    "dueDate": "2024-08-16T16:00:00.000000Z",
    "status": "done"
}
```

#### Example Response
```json
{
  "message": "Successfully created"
}
```

### 4. Update a Task
- **Endpoint**: `/tasks/:id`
- **Method**: `PUT`
- **Authorization**: **Admin Only**
- **Description**: Allows an admin user to update a specific task.
- **Request Body**:
  - Optional fields like `id`, `title`, `description`, `dueDate`, and `status`.
- **Response**:
  - Status: `201 Created`
  - Message: `"Updated successfully"`

#### Example Request
```http
PUT localhost:3000/tasks/3
Content-Type: application/json

{
    "id": "10",
    "title": "Task 10",
    "description": "Tenth task",
    "dueDate": "2024-08-16T16:00:00.000000Z",
    "status": "done"
}
```

#### Example Response
```json
{
  "message": "Updated successfully"
}
```

### 5. Delete a Task
- **Endpoint**: `/tasks/:id`
- **Method**: `DELETE`
- **Authorization**: **Admin Only**
- **Description**: Allows an admin user to delete a specific task.
- **Response**:
  - Status: `200 OK`
  - Message: `"Deleted successfully"`

#### Example Request
```http
DELETE localhost:3000/tasks/2
```

#### Example Response
```json
{
  "message": "Deleted successfully"
}
```

### 6. Promote a User to Admin
- **Endpoint**: `/promote/:id`
- **Method**: `PUT`
- **Authorization**: **Admin Only**
- **Description**: Allows an admin user to promote another user to the admin role.
- **Response**:
  - A JSON object containing an error message if the promotion fails.

#### Example Request
```http
PUT localhost:3000/promote/5
Authorization: Bearer <your-token>
```

#### Example Response
```json
{
  "error": "Unauthorized"
}
```

### 7. User Registration
- **Endpoint**: `/register`
- **Method**: `POST`
- **Description**: Registers a new user.
- **Request Body**:
  - `id` (string): The unique identifier for the user.
  - `role` (string): The role of the user.
  - `password` (string): The password for the user.
- **Response**:
  - A JSON object containing the user details and a token for authentication.

#### Example Request
```http
POST localhost:3000/register
Content-Type: application/json

{
    "id": "3",
    "role": "user",
    "password": "1234"
}
```

#### Example Response
```json
{
  "id": "3",
  "role": "user",
  "password": "1234",
  "token": "<your-token>"
}
```

### 8. User Login
- **Endpoint**: `/login`
- **Method**: `POST`
- **Description**: Authenticates a user and generates a token.
- **Request Body**:
  - `id` (string): The user's ID.
  - `role` (string): The user's role.
  - `password` (string): The user's password.
- **Response**:
  - A JSON object containing the user details and a token.

#### Example Request
```http
POST localhost:3000/login
Content-Type: application/json

{
    "id": "3",
    "role": "admin",
    "password": "1234"
}
```

#### Example Response
```json
{
  "id": "3",
  "role": "admin",
  "password": "1234",
  "token": "<your-token>"
}
```

## Authorization Rules

- Only users with the `admin` role can create, update, or delete tasks.
- Only users with the `admin` role can promote other users to the `admin` role.

