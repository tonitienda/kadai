const express = require("express");
const { v4: uuidv4 } = require("uuid");

const ownerId = "6c8d5572-a815-49aa-9e6d-7fee79ddd59d";

function makeTaskRouter(datasource) {
  const router = express.Router();

  async function addTask(req, res) {
    const task = req.body;
    // TODO - Add validation
    console.log("task: ", task);
    task.id = uuidv4();
    task.ownerId = ownerId;
    task.status = "pending";
    await datasource.addTask(task);

    res.json(task);
  }

  router.get("/", async (req, res) => {
    const tasks = await datasource.getTasks(ownerId);
    res.json(tasks);
  });

  router.get("", async (req, res) => {
    const tasks = await datasource.getTasks(ownerId);
    res.json(tasks);
  });

  router.get("/:taskId", async (req, res) => {
    const task = await datasource.getTaskById(req.params.taskId);
    res.json(task);
  });

  router.post("", async (req, res) => {
    console.log("POST ''");
    addTask(req, res);
  });

  router.post("/", async (req, res) => {
    console.log("POST /");
    addTask(req, res);
  });

  router.delete("/:taskId", async (req, res) => {
    await datasource.deleteTask(req.params.taskId);
    res.status(201).end();
  });
  return router;
}

module.exports = makeTaskRouter;
