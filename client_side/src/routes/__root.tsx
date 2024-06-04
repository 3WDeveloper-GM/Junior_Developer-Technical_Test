import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";
import { Index } from "../components/common/Index";
import { userContext } from "@/hooks/userHooks";
import { QueryClient } from "@tanstack/react-query";

interface RouterContext {
  User: userContext;
  Client: QueryClient;
}

export const Route = createRootRouteWithContext<RouterContext>()({
  component: () => (
    <>
      <Index />
      <Outlet />
    </>
  ),
});
