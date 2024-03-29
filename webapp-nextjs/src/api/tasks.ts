export async function getTasks() {
  const res = await fetch("/api/tasks");

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    const b = await res.json();
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export async function addTask(title: string, description: string) {
  const res = await fetch("/api/tasks", {
    method: "POST",
    body: JSON.stringify({
      title,
      description,
    }),
  });
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to add task");
  }

  return res.json();
}

export async function deleteTask(taskID: string) {
  const res = await fetch(`/api/tasks/` + taskID, {
    method: "DELETE",
  });
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    alert(res.status);
    throw new Error("Failed to delete task");
  }

  return res.json();
}
