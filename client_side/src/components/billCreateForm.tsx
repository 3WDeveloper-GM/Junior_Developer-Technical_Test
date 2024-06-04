import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "./ui/button";
import { UseMutateAsyncFunction } from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import { FC } from "react";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { Input } from "./ui/input";
import { Bill } from "@/type/bills";
import { createBillParams } from "@/api/bills";
import { useUser } from "@/hooks/userHooks";
import { DateTime } from "luxon";
import { Link } from "@tanstack/react-router";

const FormSchema = z.object({
  billID: z.coerce.string().uuid(),
  emissionDate: z.coerce.string().date(),
  totalAmmount: z.coerce
    .number()
    .min(0, { message: "must be a positive number" }),
  details: z.object({
    revenueTax: z.coerce
      .number()
      .min(0, { message: "must be a positive number" }),
    clientType: z.coerce.string(),
    address: z.coerce.string(),
    clientName: z
      .string()
      .min(8, { message: "must be longer than 8 characters" }),
  }),
  miscelaneous: z.object({
    energyConsumption: z.coerce
      .number()
      .min(0, { message: "must be a positive number" }),
    tariffType: z.coerce.string(),
  }),
});

interface BillFormProps {
  mutate: UseMutateAsyncFunction<
    AxiosResponse<unknown, any>,
    Error,
    createBillParams,
    unknown
  >;
  previousData?: Bill;
}

export const BillForm: FC<BillFormProps> = (props) => {
  const { previousData } = props;

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      billID: crypto.randomUUID() || previousData?.idFactura,
      details: {
        revenueTax: 0 || previousData?.detalles.revenueTax,
        clientType: "" || previousData?.detalles.clientType,
        address: "" || previousData?.detalles.address,
        clientName: "" || previousData?.detalles.clientName,
      },
      miscelaneous: {
        energyConsumption: 0 || previousData?.miscelaneo.energyConsumption,
        tariffType: "" || previousData?.miscelaneo.tariffType,
      },
      totalAmmount: 0 || previousData?.montoTotal,
      emissionDate:
        localStorage.getItem("billDate") ||
        DateTime.now().toFormat("yyyy-MM-dd"),
    },
  });

  const { authToken } = useUser().getUser();
  const submit = async (data: z.infer<typeof FormSchema>) => {
    const bill: Omit<Bill, "proveedor"> = {
      idFactura: data.billID,
      fechaEmision: DateTime.fromISO(data.emissionDate).toISO(),
      montoTotal: data.totalAmmount,
      detalles: data.details,
      miscelaneo: data.miscelaneous,
    };

    console.log("sent the following", bill);

    props.mutate({ bill: bill, token: authToken });
  };

  const { errors } = form.formState;
  const values = form.getValues();
  const formInputTags = [
    {
      id: 1,
      name: "billID",
      placeHolder: crypto.randomUUID(),
      error: errors.billID,
      type: "text",
      label: "Bill ID",
      value: values.billID,
    },
    {
      id: 2,
      name: "totalAmmount",
      placeHolder: 0,
      error: errors.totalAmmount,
      type: "number",
      label: "Total Amount",
      value: values.totalAmmount,
    },
    {
      id: 3,
      name: "emissionDate",
      placeHolder: "",
      error: errors.emissionDate,
      type: "date",
      label: "Date",
      value: values.emissionDate,
    },
    {
      id: 4,
      name: "details.clientName",
      placeHolder: "",
      error: errors.details?.clientName,
      type: "text",
      label: "Client Name",
      value: values.details.clientName,
    },
    {
      id: 5,
      name: "details.revenueTax",
      placeHolder: 0,
      error: errors.details?.revenueTax,
      type: "number",
      label: "Revenue Tax",
      value: values.details.revenueTax,
    },
    {
      id: 6,
      name: "details.clientType",
      placeHolder: "",
      error: errors.details?.clientType,
      type: "text",
      label: "Client Type",
      value: values.details.clientType,
    },
    {
      id: 7,
      name: "details.address",
      placeHolder: "",
      error: errors.details?.address,
      type: "text",
      label: "Address",
      value: values.details.address,
    },
    {
      id: 8,
      name: "miscelaneous.energyConsumption",
      placeHolder: 0,
      error: errors.miscelaneous?.energyConsumption,
      type: "number",
      label: "Energy Consumption",
      value: values.miscelaneous.energyConsumption,
    },
    {
      id: 9,
      name: "miscelaneous.tariffType",
      placeHolder: "",
      error: errors.miscelaneous?.tariffType,
      type: "text",
      label: "Tariff Type",
      value: values.miscelaneous.tariffType,
    },
  ];

  return (
    <Form {...form}>
      <form
        className="grid grid-cols-3 place-items-center gap-16 w-full "
        onSubmit={form.handleSubmit(submit)}
      >
        {formInputTags.map((item) => (
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
                      placeholder={item.placeHolder}
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
        <div></div>
        <div className="grid grid-cols-2 place-items-center gap-5">
          <div>
            <Button type="submit">Submit Bill</Button>
          </div>
          <div>
            <Button type="button">
              <Link to="/bills-fetch">Back to Fetch</Link>
            </Button>
          </div>
        </div>
        <div></div>
      </form>
    </Form>
  );
};
