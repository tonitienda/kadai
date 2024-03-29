import { NextApiRequest } from "next";
import { buildUrl, getAuthHeader } from "../../tools";

export const DELETE = async function deleteTask(
  req: NextApiRequest,
  route: { params: { id: string } }
) {
  const taskId = route.params.id;

  const res = await fetch(buildUrl("/v0/tasks/" + taskId), {
    method: "DELETE",
    headers: {
      Authorization: await getAuthHeader(),
    },
  });

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to delete task");
  }

  const resData = res.json();
  return Response.json(resData);
};
