import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  beforeLoad: ({ context }) => {
    console.log(context.User.getUser());
    throw redirect({
      to: "/login",
    });
  },
  component: () => <div>Hello /!</div>,
});
