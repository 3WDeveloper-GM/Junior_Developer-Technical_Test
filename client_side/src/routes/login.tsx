import { createFileRoute } from "@tanstack/react-router";
import { LoginView } from "../layout/login";

export const Route = createFileRoute("/login")({
  component: () => (
    <>
      <LoginView />
    </>
  ),
});
