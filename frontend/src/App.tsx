import { RouterProvider, createRouter } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";
import AuthProvider, { useProvideAuth } from "./src/utils/authentication";

const router = createRouter({
  routeTree,
  context: { authentication: undefined! },
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const App = () => {
  const AuthenticationContext = useProvideAuth();

  return (
    <AuthProvider>
      <RouterProvider
        router={router}
        context={{ authentication: AuthenticationContext }}
      />
    </AuthProvider>
  );
};

export default App;
