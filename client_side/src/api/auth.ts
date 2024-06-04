import { Navigate } from "@tanstack/react-router";
import { user } from "../type/user";
import api from "./api";

export async function whoAmI({ queryKey }: { queryKey: unknown[] }) {
  const [_key, { header }] = queryKey;
  return api.get("/public/whoAmI", {
    headers: { Authorization: "Bearer " + header },
    withCredentials: true,
  });
}

export const login = (user: Omit<user, "authToken">) =>
  api.post("/public/login", user);

export const protectRoute = async (user: Omit<user, "password">) => {
  try {
    await api.get("/public/whoAmI", {
      headers: { Authorization: "Bearer " + user.authToken },
    });
  } catch {
    Navigate({ to: "/login" });
  }
};
