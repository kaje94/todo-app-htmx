# Server Side of the Todo App

This is the server side of the Todo App. It is a REST API built with Express.js and uses SQLite as the database.

## Getting Started

To get started, you need to install the dependencies. Run the following command in the `server` directory:

```
npm install
```

This will install all the necessary dependencies listed in the `package.json` file.

## Running the Server

To start the server, run the following command in the `server` directory:

```
npm start
```

This will start the server on the port specified in the `.env` file. If no port is specified, it will default to port 5000.

## API Endpoints

The server has the following endpoints:

- `GET /todos`: Returns a list of all todos.
- `POST /todos`: Creates a new todo. The body of the request should be a JSON object with a `title` and `completed` property.
- `GET /todos/:id`: Returns the todo with the specified id.
- `PUT /todos/:id`: Updates the todo with the specified id. The body of the request should be a JSON object with the properties to be updated.
- `DELETE /todos/:id`: Deletes the todo with the specified id.

## Environment Variables

The server uses the following environment variables:

- `PORT`: The port on which the server runs.

These variables are stored in the `.env` file.

## Database

The server uses SQLite as the database. The database file is `database.sqlite`.

The `TodoModel` class in `src/models/todoModel.js` represents a todo item in the database. It has methods to interact with the database.

## Routes and Controllers

The routes for the todos resource are set up in `src/routes/todos.js`. These routes use the `TodosController` class in `src/controllers/todosController.js` to handle requests.

## Testing

To run tests, use the following command:

```
npm test
```

This will run all tests defined in the `test` directory.