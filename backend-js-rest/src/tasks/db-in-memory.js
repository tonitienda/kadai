const TasksByOwner = {};
const TasksById = {};

function getTasks(ownerId) {
  if (!TasksByOwner[ownerId]) {
    TasksByOwner[ownerId] = [];
  }

  return Promise.resolve(TasksByOwner[ownerId]);
}

function getTaskById(taskId) {
  return Promise.resolve(TasksById[taskId]);
}

function addTask(task) {
  if (!TasksByOwner[task.ownerId]) {
    TasksByOwner[task.ownerId] = [];
  }

  TasksByOwner[task.ownerId].push(task);
  TasksById[task.id] = task;
}

function deleteTask(taskId) {
  const task = TasksById[taskId];
  if (!task) {
    return;
  }

  delete TasksById[taskId];

  const tasks = TasksByOwner[task.ownerId];
  if (!tasks) {
    return;
  }

  const index = tasks.indexOf(task);
  if (index === -1) {
    return;
  }

  tasks.splice(index, 1);
}

module.exports = {
  getTasks,
  getTaskById,
  addTask,
  deleteTask,
};
