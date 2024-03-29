import type { ActionFunctionArgs } from "@remix-run/node";

import { redirect } from "@remix-run/node";

import { destroySession, getSession } from "~/services/session.server";

export const action = async ({ request }: ActionFunctionArgs) => {
  const session = await getSession(request.headers.get("Cookie"));
  const logoutURL = new URL(process.env.AUTH0_LOGOUT_URL || ""); // i.e https://YOUR_TENANT.us.auth0.com/v2/logout

  logoutURL.searchParams.set(
    "client_id",
    // FIXME: set client id in env
    process.env.AUTH0_CLIENT_ID || ""
  );
  logoutURL.searchParams.set("returnTo", "http://localhost:3000" || "");

  return redirect(logoutURL.toString(), {
    headers: {
      "Set-Cookie": await destroySession(session),
    },
  });
};
