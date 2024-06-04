import { whoAmI } from "@/api/auth";
import { NavBar } from "@/components/common/navBar";
import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/_bills")({
  beforeLoad: async ({ context }) => {
    const { User, Client } = context;
    const { authToken } = User.getUser();
    try {
      console.log("tried")
      Client.getQueryData(["whoAmI", { header: authToken }]) ??
        (await Client.fetchQuery({
          queryKey: ["whoAmI", { header: authToken }],
          queryFn: whoAmI,
          retry: 1,
        }));
    } catch {
      console.log("catched")
      throw redirect({ to: "/login" })
    }
  },
  component: () => (
    <div>
      <NavBar />
      <Outlet />
    </div>
  ),
});
