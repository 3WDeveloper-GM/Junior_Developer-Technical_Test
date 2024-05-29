import { ReactNode } from "@tanstack/react-router";
import { createContext, useState } from "react";

interface user {
  name: string;
  id: string;
  authenticated: boolean;
  logOut: () => void;
  logIn: (name: string, id: string) => void;
}

export const AuthContext = createContext<user | null>(null);

export function useProvideAuth(): user {
  const [name, setName] = useState("");
  const [id, setId] = useState("");
  const [authenticated, setAuthenticated] = useState(false);

  const logOut = (): void => {
    setName("");
    setId("");
    setAuthenticated(false);
  };

  const logIn = (name: string, id: string) => {
    console.log("logging user");
    setName(name)
    setId(id);
    setAuthenticated(true);
  };

  return {
    name: name,
    id: id,
    authenticated: authenticated,
    logIn: logIn,
    logOut: logOut,
  };
}
export default function AuthProvider({ children }: { children: ReactNode }) {
  const auth = useProvideAuth();
  return <AuthContext.Provider value={auth}> {children} </AuthContext.Provider>;
}
