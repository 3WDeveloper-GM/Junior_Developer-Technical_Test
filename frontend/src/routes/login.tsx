import { createFileRoute } from "@tanstack/react-router";
import { LoginForm } from "../src/components/pages/login";

export const Route = createFileRoute("/login")({
  component: () => <LoginForm />,
});
