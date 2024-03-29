import { Form } from "@remix-run/react";

export default function Register() {
  return (
    <Form action="/auth/auth0?screen_hint=signup" method="post">
      <button>Register with Auth0</button>
    </Form>
  );
}
