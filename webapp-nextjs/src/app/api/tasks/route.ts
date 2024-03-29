import { buildUrl, getAuthHeader } from "../tools";

export const GET = async function getTasks() {
  const res = await fetch(buildUrl("/v0/tasks"), {
    cache: "no-store",
    headers: {
      Authorization: await getAuthHeader(),
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

  const res = await fetch(buildUrl("/v0/tasks"), {
    method: "POST",
    body: JSON.stringify({
      title: reqData.title,
      description: reqData.description,
    }),

    headers: {
      Authorization: await getAuthHeader(),
    },
  });

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to add task");
  }

  const resData = res.json();
  return Response.json(resData);
};
