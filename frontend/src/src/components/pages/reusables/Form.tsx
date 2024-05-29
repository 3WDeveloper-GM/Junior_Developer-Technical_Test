"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "../../ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../../ui/form";
import { Input } from "../../ui/input";
import axios from "axios";
import { useAuth } from "@/hooks/useAuth";

interface FormProps {
  alert: string;
  setAlert: React.Dispatch<React.SetStateAction<string>>;
  setError: React.Dispatch<React.SetStateAction<boolean>>;
}

const FormSchema = z.object({
  username: z.string().email().min(2, {
    message: "Username must be at least 2 characters.",
  }),
  password: z.string().min(8, {
    message: "Password must be at least 8 characters",
  }),
});

const InputForm: React.FC<FormProps> = (props): JSX.Element => {

  const context = useAuth()

  const { setAlert, setError } = props;
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  function onSubmit(data: z.infer<typeof FormSchema>) {
    const authKeys = {
      correoUsuario: data.username,
      contraUsuario: data.password,
    };
    axios
      .post("http://localhost:4040/public/login", authKeys)
      .then((response) => {
        const {id, usuario} = response.data.resultado

        console.log(context)

        context.logIn(usuario, id)

        localStorage.setItem("username",usuario)
        localStorage.setItem("userID",id)

        setAlert(response.data.message);
        setError(false)

        console.log(context)

      })
      .catch((reject) => {
        console.error("unsuccessful");
        setAlert(reject.response.data.error);
        setError(true)
      });
  }

  return (
    <>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="w-4/5 space-y-2"
        >
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <>
                <FormItem className="my-4">
                  <FormLabel className="text-gray-400">Username</FormLabel>
                  <FormControl>
                    <Input placeholder="john.doe@xyz.com" {...field} />
                  </FormControl>
                  <FormDescription></FormDescription>
                  <FormMessage />
                </FormItem>
              </>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <>
                <FormItem className="my-8">
                  <FormLabel className="text-gray-400">Password</FormLabel>
                  <FormControl>
                    <Input
                      type="password"
                      placeholder=""
                      className="my-8"
                      {...field}
                    />
                  </FormControl>
                  <FormDescription></FormDescription>
                  <FormMessage />
                </FormItem>
              </>
            )}
          />
          <div className="my-8">
            <Button
              variant="outline"
              className="bg-gray-900 text-gray-400 outline-gray-400"
              type="submit"
            >
              Submit
            </Button>
          </div>
        </form>
      </Form>
    </>
  );
};

export default InputForm;
