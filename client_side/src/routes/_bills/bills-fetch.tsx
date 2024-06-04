import { BillsFetchView } from "@/layout/billFetch";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_bills/bills-fetch")({
  component: () => <BillsFetchView />,
});
