const express = require("express");
const { v4: uuidv4 } = require("uuid");
const core = require("./core");

// TODO - Implement auth and get user id by token
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
    if (!req.headers["authorization"]) {
      res.status(401).end();
      return;
    }

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

  router.post("/:taskId/undo-delete", async (req, res) => {
    console.log("POST /undo-delete");
    const taskId = req.params.taskId;
    await core.undeleteTask(datasource, taskId, ownerId);

    res.status(202).end();
  });

  router.delete("/:taskId", async (req, res) => {
    const taskId = req.params.taskId;
    await core.deleteTask(datasource, taskId, ownerId);

    res.status(202).json({
      url: `/v0/tasks/${taskId}/undo-delete`,
      method: "POST",
    });
  });
  return router;
}

module.exports = makeTaskRouter;
