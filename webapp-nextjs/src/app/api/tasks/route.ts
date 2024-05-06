import { buildUrl, getAuthHeader } from "../tools";

export const GET = async function getTasks() {
  const authHeader = await getAuthHeader();

  if (!authHeader) {
    return Response.json([]);
  }

  const res = await fetch(buildUrl("/v0/tasks"), {
    cache: "no-store",
    headers: {
      Authorization: authHeader,
    },
  });

  if (!res.ok) {
    const b = await res.json();
    throw new Error("Failed to fetch data");
  }

  const data = await res.json();
  return Response.json(data);
};

export const POST = async function addTask(req: Request) {
  const reqData = await req.json();
  console.log("adding a new task");
  const res = await fetch(buildUrl("/v0/tasks"), {
    method: "POST",
    body: JSON.stringify({
      title: reqData.title,
      description: reqData.description,
    }),

    headers: {
      Authorization: await getAuthHeader(),
      "Content-Type": "application/json",
    },
  });

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    console.log("failed to add a new task");
    console.error(await res.status);
    console.error(await res.text());
    throw new Error("Failed to add task");
  }

  console.log("task added successfully");

  const resData = res.json();
  return Response.json(resData);
};
