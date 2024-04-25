import { buildUrl, getAuthHeader } from "../../tools";

export const POST = async function undoAction(req: Request) {
  const body = await req.json();

  console.log(body);

  const res = await fetch(buildUrl(body.url), {
    method: body.method,
    headers: {
      Authorization: await getAuthHeader(),
    },
  });

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to undo action");
  }

  const resData = res.json();
  return Response.json(resData);
};
