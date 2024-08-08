# Task Management API

The **Task Management API** provides a comprehensive set of tools and resources that enable you to efficiently manage tasks within your application.

## Endpoints

### 1. Retrieve All Tasks

**GET /tasks**

This endpoint retrieves a list of tasks.

#### Request

No request body is required for this endpoint.

#### Response

The response will be a JSON array containing task objects with the following properties:

- `id` (string): The unique identifier for the task.
- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `due_date` (string): The due date of the task.
- `status` (string): The status of the task.

Example response body:

```json
[
    {
        "id": "23",
        "title": "hello world",
        "description": "world is nice",
        "due_date": "0001-01-01T00:00:00Z",
        "status": "done"
    }
]
```

### 2. Retrieve a Task by ID

**GET /tasks/5**

This endpoint retrieves the details of a specific task identified by the ID `5`.

#### Request

No request body is required for this endpoint.

#### Response

The response will be in JSON format with a status code of 201. It will contain an array of objects, where each object represents a task with the following properties:

- `id` (string): The unique identifier of the task.
- `title` (string): The title or name of the task.
- `description` (string): The description or details of the task.
- `due_date` (string): The due date for the task.
- `status` (string): The status of the task.

Example response body:

```json
[
    {
        "id": "5",
        "title": "Task 5",
        "description": "Fifth task",
        "due_date": "0001-01-01T00:00:00Z",
        "status": "In Progress"
    }
]
```

### 3. Add a New Task

**POST /tasks**

This endpoint allows the client to add a new task by sending a POST request to the specified URL.

#### Request Body

- `id` (text, optional): The unique identifier for the task.
- `title` (text, required): The title of the task.
- `description` (text, required): The description of the task.
- `dueDate` (text, optional): The due date for the task.
- `status` (text, optional): The status of the task.

Example request body:

```json
{
    "id": "10",
    "title": "Task 10",
    "description": "Tenth task",
    "dueDate": "2024-08-16T16:00:00.000000Z",
    "status": "Pending"
}
```

#### Response

The response will include nothing with a status code of 204, or an error message if the task addition was unsuccessful.

### 4. Update a Task by ID

**PUT /tasks/5**

This endpoint is used to update a specific task by its ID.

#### Request Body

- `id` (string): The ID of the task.
- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `dueDate` (string): The due date of the task.
- `status` (string): The status of the task.

Example request body:

```json
{
    "id": "10",
    "title": "Task 10",
    "description": "Tenth task",
    "dueDate": "2024-08-16T16:00:00.000000Z",
    "status": "done"
}
```

#### Response

The response schema for this request is as follows:

```json
{
  "type": "object",
  "properties": {
    "message": {
      "type": "string"
    }
  }
}
```

### 5. Delete a Task by ID

**DELETE /tasks/5**

This endpoint is used to delete a specific task.

#### Request

This request does not require a request body.

#### Response

The response will not contain a response body. The status code will indicate the success or failure of the deletion operation.
