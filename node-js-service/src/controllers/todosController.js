const TodoModel = require("../models/todoModel");
const oauth = require("axios-oauth-client");
const axios = require("axios");

const serviceURL = process.env.SVC_URL;

const getClientCredentials = oauth.clientCredentials(
  axios.create(),
  process.env.TOKEN_URL,
  process.env.CONSUMER_KEY,
  process.env.CONSUMER_SECRET
);

class TodosController {
  // Get all todos
  async getAll(req, res) {
    try {
      console.log('Running "getAll" ');
      const todos = await TodoModel.findAll();

      const auth = await getClientCredentials();
      const accessToken = auth.access_token;
      const response = await axios.get(`${serviceURL}/todos`, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
      console.log("response:", response);

      res.status(200).json(todos);
    } catch (error) {
      res.status(500).json({ message: "Error getting todos" });
    }
  }

  // Get a single todo by id
  async getById(req, res) {
    try {
      console.log('Running "getById"');
      const todo = await TodoModel.findById(req.params.id);
      if (todo) {
        res.status(200).json(todo);
      } else {
        res.status(404).json({ message: "Todo not found" });
      }
    } catch (error) {
      res.status(500).json({ message: "Error getting todo" });
    }
  }

  // Create a new todo
  async create(req, res) {
    try {
      console.log('Running "create"');
      const todo = await TodoModel.create(req.body);
      res.status(201).json(todo);
    } catch (error) {
      res.status(500).json({ message: "Error creating todo" });
    }
  }

  // Update a todo
  async update(req, res) {
    try {
      console.log('Running "update"');
      const todo = await TodoModel.update(req.params.id, req.body);
      if (todo) {
        res.status(200).json(todo);
      } else {
        res.status(404).json({ message: "Todo not found" });
      }
    } catch (error) {
      res.status(500).json({ message: "Error updating todo" });
    }
  }

  // Delete a todo
  async delete(req, res) {
    try {
      console.log('Running "delete"');
      const todo = await TodoModel.destroy({ where: { id: req.params.id } });
      if (todo) {
        res.status(200).json({ message: "Todo deleted" });
      } else {
        res.status(404).json({ message: "Todo not found" });
      }
    } catch (error) {
      res
        .status(500)
        .json({ message: "Error deleting todo, " + error?.message });
    }
  }
}

module.exports = new TodosController();
