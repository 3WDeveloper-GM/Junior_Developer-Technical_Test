import { LogForm } from "../components/loginForm";
import { useMutation } from "@tanstack/react-query";
import { login } from "@/api/auth";
import { Navigate } from "@tanstack/react-router";

type loginResponse = {
  userName: string;
  token: string;
  userEmail: string;
};

export const LoginView = () => {

  const { data, mutateAsync, isError, isPending, isSuccess } = useMutation({
    mutationFn: login,
  });

  if (isSuccess) {
    const response: loginResponse = data.data as loginResponse;

    localStorage.setItem("token", response.token);
    localStorage.setItem("userName", response.userName);
    localStorage.setItem("userEmail", response.userEmail);
  }

  return (
    <div>
      {isSuccess ? <Navigate to="/bills-read" /> : null}
      <div className="w-full lg:grid lg:min-h-[600px] lg:grid-cols-2 xl:min-h-[800px] place-items-center h-[100vh]">
        <div className="flex items-center justify-center py-12">
          <div className="mx-auto grid w-[350px] gap-6">
            <div className="grid gap-2 text-center">
              <h1 className="text-3xl font-bold">Login</h1>
              <p className="text-balance text-muted-foreground">
                Enter your email below to login to your account
              </p>
              <LogForm mutate={mutateAsync} />
            </div>
          </div>
        </div>
        <div className="bg-muted-foreground">
          {isError ? <p>error</p> : null}
          {isPending ? <p>loading...</p> : null}
          {isSuccess ? <Navigate to="/bills-read" /> : null}
        </div>
      </div>
    </div>
  );
};
