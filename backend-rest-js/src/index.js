const express = require("express");
const tasks = require("./tasks");

const app = express();
const v0 = express.Router();

app.use(express.json());

app.use((req, res, next) => {
  console.log(`[${req.method}] ${req.url}`);
  next();
});

app.get("/healthz", (req, res) => {
  res.json(`{"message": "OK"}`);
});

app.use("/v0", v0);
v0.use("/tasks", tasks);

app.listen(8080, () => {
  console.log("Server running on port 8080");
});
