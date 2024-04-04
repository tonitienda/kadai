import { getAccessToken } from "@auth0/nextjs-auth0";

export function buildUrl(path: string) {
  return `${process.env.BACKEND_BASE_URL}${path}`;
}

export async function getAuthHeader() {
  try {
    const accessTokenResponse = await getAccessToken({
      authorizationParams: {
        audience: process.env.AUTH0_AUDIENCE || "",
      },
    });

    return "Bearer " + accessTokenResponse.accessToken || "";
  } catch (error) {
    return "";
  }
}
