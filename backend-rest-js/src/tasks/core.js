export async function deleteTask(datasource, taskId, userId) {
  const task = await datasource.getTaskById(taskId);

  // TODO - Validate ownership and task status
  task.deletedAt = new Date();
  task.deletedBy = userId;

  return datasource.updateTask(task);
}

export async function undeleteTask(datasource, taskId, userId) {
  console.log("Undeleting ", taskId);
  const task = await datasource.getTaskById(taskId);

  // TODO - Validate ownership and task status
  task.deletedAt = null;
  task.deletedBy = null;

  return datasource.updateTask(task);
}
