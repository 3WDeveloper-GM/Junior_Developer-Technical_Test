import { InitializeAuth } from "@/hooks/useAuth";
import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  beforeLoad: async ({ location, context }) => {
    const { authenticated } = context.authentication;
    await InitializeAuth().then((response) => {
      const { ususario, userid } = response.data;
      context.authentication.logIn(ususario, userid);
    }).catch((error) => console.log(error.response));
    console.log(authenticated);
    if (!authenticated) {
      throw redirect({
        to: "/login",
        search: {
          redirect: location.href,
        },
      });
    }
  },
  component: () => <div>Hello /profile!</div>,
});
