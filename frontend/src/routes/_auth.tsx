import { InitializeAuth } from "@/hooks/useAuth";
import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/_auth")({
  beforeLoad: async ({ location, context }) => {
    const {authenticated} = context.authentication
    await InitializeAuth().then(
      (response) => {
        const {ususario, userid} = response.data
        context.authentication.logIn(ususario, userid)
      }
    )
    console.log(authenticated)
    if (!authenticated) {
      throw redirect({
        to: "/login",
        search: {
          redirect: location.href,
        },
      });
    }
  },
});
