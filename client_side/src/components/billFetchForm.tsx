import { UseFormReturn } from "react-hook-form";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { FC } from "react";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import { QueryObserverResult, RefetchOptions } from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import { BillsFetchSchema } from "@/type/formSchemas";

interface FetchFormProps {
  form: UseFormReturn<
    {
      startDate: string;
      endDate: string;
    },
    any,
    undefined
  >;
  refetch: (
    options?: RefetchOptions | undefined,
  ) => Promise<QueryObserverResult<AxiosResponse<unknown, any>, Error>>;
}

export const FetchForm: FC<FetchFormProps> = (props) => {
  const { form, refetch } = props;

  const { errors } = form.formState;

  const { startDate, endDate } = form.getValues();

  const formFieldTags = [
    {
      id: 0,
      label: "Start Date",
      type: "date",
      placeholder: "",
      name: "startDate",
      error: errors.startDate,
      value: startDate,
    },
    {
      id: 1,
      label: "End Date",
      type: "date",
      placeholder: "",
      name: "endDate",
      error: errors.endDate,
      value: endDate,
    },
  ];

  const submit = async (data: z.infer<typeof BillsFetchSchema>) => {
    console.log("sent data:", data.startDate, data.endDate);
    localStorage.setItem("startDate", data.startDate);
    localStorage.setItem("endDate", data.endDate);
    refetch();
  };

  return (
    <Form {...form}>
      <form
        className="grid auto-rows-fr place-items-center gap-4 w-9/12 "
        onSubmit={form.handleSubmit(submit)}
      >
        {formFieldTags.map((item) => (
          <div key={item.id}>
            <FormField
              control={form.control}
              name={item.name}
              render={({ field }) => (
                <FormItem>
                  <FormLabel>{item.label}</FormLabel>
                  <FormControl>
                    <Input
                      type={item.type}
                      placeholder={item.placeholder}
                      {...field}
                      value={item.value}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        ))}
          <div>
            <Button type="submit">Fetch</Button>
          </div>
      </form>
    </Form>
  );
};
