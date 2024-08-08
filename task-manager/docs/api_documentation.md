# Task Management API

The Task Management API provides a comprehensive set of tools and resources that enable you to efficiently manage tasks within your application.

## API Overview

The API consists of several endpoints that allow you to retrieve, create, update, and delete tasks. Each endpoint is described below.

## Endpoints

### 1. Retrieve All Tasks

- **Endpoint**: `/tasks`
- **Method**: `GET`
- **Description**: Retrieves a list of all tasks from the server.
- **Response**: JSON array containing objects with the following properties:
  - `id` (string): The unique identifier of the task.
  - `title` (string): The title or name of the task.
  - `description` (string): A description or details about the task.
  - `due_date` (string): The due date or deadline for the task.
  - `status` (string): The status of the task, indicating whether it is completed, pending, or in progress.
  
**Example Response**:
```json
[
  {
    "id": "1",
    "title": "Task 1",
    "description": "Description of Task 1",
    "due_date": "2024-08-09",
    "status": "Pending"
  }
]
```

### 2. Retrieve a Specific Task

- **Endpoint**: `/task/:id`
- **Method**: `GET`
- **Description**: Retrieves a specific task by its ID.
- **Response**: JSON object containing the task's ID, title, description, due date, and status.

**Example Request**:
```http
GET localhost:3000/tasks/2
```

**Example Response**:
```json
{
  "id": "2",
  "title": "Task 2",
  "description": "Description of Task 2",
  "due_date": "2024-08-10",
  "status": "In Progress"
}
```

### 3. Create a New Task

- **Endpoint**: `/tasks`
- **Method**: `POST`
- **Description**: Creates a new task. The request should include a JSON payload with the task details.
- **Request Body**:
  - `id` (string): The unique identifier for the task.
  - `title` (string): The title of the task.
  - `description` (string): The description of the task.
  - `dueDate` (string): The due date of the task.
  - `status` (string): The status of the task.

**Example Request Body**:
```json
{
  "id": "10",
  "title": "Task 10",
  "description": "Tenth task",
  "dueDate": "2024-08-16T16:00:00.000Z",
  "status": "Done"
}
```

**Response**:
- **Status**: `201 Created`
- **Body**: `{ "message": "Successfully created" }`

### 4. Update a Task

- **Endpoint**: `/task/:id`
- **Method**: `PUT`
- **Description**: Updates a specific task. The request should include a JSON payload with the updated task details.
- **Request Body**:
  - `id` (string, optional): The unique identifier of the task.
  - `title` (string, optional): The title of the task.
  - `description` (string, optional): The description of the task.
  - `dueDate` (string, optional): The due date of the task.
  - `status` (string, optional): The status of the task.

**Example Request Body**:
```json
{
  "id": "10",
  "title": "Updated Task 10",
  "description": "Updated description",
  "dueDate": "2024-08-16T16:00:00.000Z",
  "status": "Done"
}
```

**Response**:
- **Status**: `201 Created`
- **Body**: `Updated successfully`

### 5. Delete a Task

- **Endpoint**: `/task/:id`
- **Method**: `DELETE`
- **Description**: Deletes a specific task by its ID.

**Example Request**:
```http
DELETE localhost:3000/tasks/2
```

**Response**:
- **Status**: `200 OK`
- **Body**: `Task deleted successfully`

## Installation and Setup

1. **Clone the Repository**:
   ```sh
   git clone <repository-url>
   ```
2. **Install Dependencies**:
   Navigate to the project directory and install the necessary dependencies:
   ```sh
   cd task_management
   go mod tidy
   ```
3. **Run the Server**:
   Start the server:
   ```sh
   go run main.go
   ```

## Usage

Use the provided endpoints to manage tasks within your application. You can interact with the API using tools like [Postman](https://www.postman.com) or [cURL](https://curl.se/).

