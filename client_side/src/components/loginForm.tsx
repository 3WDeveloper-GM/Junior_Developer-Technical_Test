import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import FormField from "./common/formInput";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { UseMutateFunction } from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import { user } from "@/type/user";
import { FC, useState } from "react";

const FormSchema = z.object({
  email: z.string().email(),
  password: z
    .string()
    .min(8, { message: "must be a longer than 8 characters" }),
});

interface LogFormProps {
  mutate: UseMutateFunction<
    AxiosResponse<unknown, any>,
    Error,
    Omit<user, "authToken">,
    unknown
  >;
}

export const LogForm: FC<LogFormProps> = (props) => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const [formPassword, setPass] = useState("");
  const [formEmail, setEmail] = useState("");

  const { mutate } = props;

  const submit = () => {
    console.log(formPassword, formEmail)
    mutate({ password: formPassword, email: formEmail });
  };

  return (
    <form className="space-y-6" onSubmit={form.handleSubmit(submit)}>
      <div className="grid col-auto">
        <div className="grid gap-4">
          <div className="grid gap-2 my-3">
            <Label htmlFor="email">Email</Label>
            <FormField
              Setter={setEmail}
              type="text"
              name="email"
              placeholder="shadcn@gmail.com"
              register={form.register}
              error={form.formState.errors.email}
            />
          </div>
        </div>
        <div className="grid gap-4">
          <div className="grid gap-2 my-3">
            <Label htmlFor="password">Password</Label>
            <FormField
              Setter={setPass}
              type="password"
              name="password"
              placeholder=""
              register={form.register}
              error={form.formState.errors.password}
            />
          </div>
        </div>
      </div>
      <div className="grid gap-4">
        <Button type="submit">Log in</Button>
      </div>
    </form>
  );
};
