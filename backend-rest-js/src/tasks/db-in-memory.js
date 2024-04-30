const TasksByOwner = {};
const TasksById = {};

async function getTasks(ownerId) {
  console.log("Getting tasks");

  if (!TasksByOwner[ownerId]) {
    TasksByOwner[ownerId] = [];
  }

  const nonDeletedTasks = TasksByOwner[ownerId].filter((t) => !t.deletedAt);

  console.log("nonDeletedTasks", nonDeletedTasks);

  return nonDeletedTasks;
}

function getTaskById(taskId) {
  return Promise.resolve(TasksById[taskId]);
}

function addTask(task) {
  console.log("Inserting", task);

  if (!TasksByOwner[task.ownerId]) {
    TasksByOwner[task.ownerId] = [];
  }

  TasksByOwner[task.ownerId].push(task);
  TasksById[task.id] = task;
}

function updateTask(task) {
  console.log("Updating", task);
  if (!TasksByOwner[task.ownerId]) {
    TasksByOwner[task.ownerId] = [];
  }

  TasksById[task.id] = task;
  for (let idx = 0; idx < TasksByOwner[task.ownerId].length; idx++) {
    if (TasksByOwner[task.ownerId][idx].id == task.id) {
      TasksByOwner[task.ownerId][idx] = task;
      return;
    }
  }
}

module.exports = {
  getTasks,
  getTaskById,
  addTask,
  updateTask,
};
