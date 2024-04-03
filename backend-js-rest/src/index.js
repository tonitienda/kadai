const express = require("express");
const tasks = require("./tasks");

const app = express();
const router = express.Router();
const v0 = express.Router();

v0.use("/tasks", tasks);
router.use("/v0", v0);

app.listen(8081, () => {
  console.log("Server running on port 8081");
});
