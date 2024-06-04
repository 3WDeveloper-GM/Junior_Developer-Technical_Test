import { RouterProvider, createRouter } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useUser } from "./hooks/userHooks";
import { useState } from "react";
import { emptyBill } from "./type/bills";
import { CurrentBillContext } from "./contexts/billCtx";

const client = new QueryClient({
  defaultOptions: { queries: { refetchOnWindowFocus: true } },
});

const router = createRouter({
  routeTree: routeTree,
  context: { User: undefined!, Client: undefined! },
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const App = () => {
  const user = useUser();
  const [bill, setBill] = useState(emptyBill);
  return (
    <QueryClientProvider client={client}>
      <CurrentBillContext.Provider
        value={{ bill: bill, updateCurrentBill: setBill }}
      >
        <RouterProvider
          router={router}
          context={{ User: user, Client: client }}
        />
      </CurrentBillContext.Provider>
    </QueryClientProvider>
  );
};

export default App;
