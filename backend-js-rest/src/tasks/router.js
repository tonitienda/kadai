const express = require("express");

function makeTaskRouter(datasource) {
  const router = express.Router();

  router.get("/", async (req, res) => {
    const tasks = await datasource.getTasks();
    res.json(tasks);
  });

  router.get("/:taskId", async (req, res) => {
    const task = await datasource.getTaskById(req.params.taskId);
    res.json(task);
  });

  router.post("/", async (req, res) => {
    const task = req.body;

    // TODO - Add validation
    const newTask = await datasource.addTask(task);
    res.status(201);
  });

  router.delete("/:taskId", async (req, res) => {
    await datasource.deleteTask(req.params.taskId);
    res.status(201);
  });
  return router;
}

module.exports = makeTaskRouter;
