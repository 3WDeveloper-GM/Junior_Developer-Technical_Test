import api from "@/api/api";
import { user } from "../type/user";

export const useUser = () => {
  const setUserParams = (
    userName: string,
    userEmail: string,
    token: string,
  ) => {
    localStorage.setItem("token", token);
    localStorage.setItem("userEmail", userEmail);
    localStorage.setItem("userName", userName);
  };

  const getUser = () => {
    const User: Omit<user, "password"> = {
      email: localStorage.getItem("userEmail") || "",
      authToken: localStorage.getItem("token") || "",
    };

    return User;
  };

  const validate = () => {
    const { authToken } = getUser();

    return api.get("/public/whoAmI", {
      headers: { Authorization: "Bearer " + authToken },
    });
  };

  return { setUserParams, getUser, validate };
};

export type userContext = ReturnType<typeof useUser>;
