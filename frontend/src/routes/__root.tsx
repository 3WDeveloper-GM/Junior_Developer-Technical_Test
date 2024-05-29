import { AuthenticationContext } from "@/hooks/useAuth";
import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";

type RouterContext = {
  authentication: AuthenticationContext;
};

export const Route = createRootRouteWithContext<RouterContext>()({
  component: () => <Outlet />,
});
