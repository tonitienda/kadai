import { getAccessToken } from "@auth0/nextjs-auth0";

export async function getTasks() {
  const accessTokenResponse = await getAccessToken({});

  const res = await fetch("http://backend:8080/v0/tasks", {
    cache: "no-store",
    headers: {
      // FIXME: user id should be taken from the access token.
      // With this setting and the current implementation, the user ID will
      // be the same for all users.
      "X-User-ID": "b78fe6be-0642-4cfa-9f19-9cc8e53b129d",
      Authorization: "Bearer " + accessTokenResponse.accessToken || "",
    },
  });
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    const b = await res.json();
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export async function addTask() {
  const res = await fetch("http://backend:8080/v0/tasks", {
    method: "POST",
    body: JSON.stringify({
      title: "New task",
      description: "This is a new task",
    }),

    headers: {
      "X-User-ID": "b78fe6be-0642-4cfa-9f19-9cc8e53b129d",
      "Content-Type": "application/json",
    },
  });
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to add task");
  }

  return res.json();
}
