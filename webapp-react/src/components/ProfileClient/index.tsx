"use client";

import { useUser } from "@auth0/nextjs-auth0/client";

export default function ProfileClient() {
  const { user, error, isLoading } = useUser();

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>{error.message}</div>;

  return (
    <>
      <a href="/api/auth/login">Login</a>
      <a href="/api/auth/logout">Logout</a>
      {user && (
        <div>
          <img width={50} src={user.picture || ""} alt={user.name || ""} />
          <h2>{user.name}</h2>
          <p>{user.email}</p>
        </div>
      )}{" "}
    </>
  );
}
