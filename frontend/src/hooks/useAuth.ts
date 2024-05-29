import { AuthContext } from "@/src/utils/authentication";
import axios, { AxiosResponse } from "axios";
import { useContext } from "react";

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}

export const InitializeAuth = async (): Promise<AxiosResponse> => {
  return axios.get("http://localhost:4040/public/whoAmI", {
    withCredentials: true,
  });
};

export type AuthenticationContext = ReturnType<typeof useAuth>;
