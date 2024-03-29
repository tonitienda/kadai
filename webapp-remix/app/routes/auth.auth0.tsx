import { redirect, ActionFunctionArgs } from "@remix-run/node";

import { authenticator } from "~/services/auth.server";

export let loader = async ({ request }: ActionFunctionArgs) => {
  console.log("Hello from auth.auth0.tsx");
  return redirect("/");
};

export let action = async ({ request }: ActionFunctionArgs) => {
  console.log("Hello from auth.auth0.tsx");
  await authenticator.authenticate("auth0", request);
  redirect("/");
};
