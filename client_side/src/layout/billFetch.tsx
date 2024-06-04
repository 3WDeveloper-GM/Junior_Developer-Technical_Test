import { fetchDate } from "@/api/bills";
import { FetchForm } from "@/components/billFetchForm";
import { useUser } from "@/hooks/userHooks";
import { zodResolver } from "@hookform/resolvers/zod";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { add } from "date-fns";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { BillsFetchSchema } from "@/type/formSchemas";
import { BillTable } from "@/components/fetchtable/billFetchTable";
import { Bill, emptyBill } from "@/type/bills";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Link } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { useContext } from "react";
import { CurrentBillContext } from "@/contexts/billCtx";

export const BillsFetchView = () => {
  const form = useForm<z.infer<typeof BillsFetchSchema>>({
    resolver: zodResolver(BillsFetchSchema),
    defaultValues: {
      startDate: localStorage.getItem("startDate") || Date.now().toString(),
      endDate:
        localStorage.getItem("endDate") ||
        add(Date.now(), { days: 1 }).toString(),
    },
  });

  const { endDate, startDate } = form.getValues();
  const { authToken } = useUser().getUser();
  const client = useQueryClient();
  const bill = useContext(CurrentBillContext)?.bill || emptyBill;

  const { data, refetch, isSuccess, isStale } = useQuery({
    queryKey: [
      "fetchBills",
      { header: authToken, endDate: endDate, startDate: startDate },
    ],
    queryFn: fetchDate,
    enabled: false,
    retry: 0,
    staleTime: Infinity,
  });

  if (isSuccess) {
    client.setQueryData(
      [
        "fetchBills",
        { header: authToken, endDate: endDate, startDate: startDate },
      ],
      data,
    );
  }

  return (
    <div>
      <div className="w-full lg:grid lg:min-h-[600px] lg:grid-cols-[2fr_5fr] xl:min-h-[900px] place-items-center">
        <div className="flex items-center justify-center py-12">
          <div className="mx-auto grid gap-6">
            <div className="grid gap-2 text-center place-items-center">
              <h1 className="text-3xl font-bold">Fetch Bill</h1>
              <p className="text-balance text-muted-foreground">
                Fetch a New Bill from the Database.
              </p>
              <FetchForm form={form} refetch={refetch} />
            </div>
          </div>
        </div>
        <div className="w-[77rem] grid grid-rows-[3fr_1fr] place-items-center">
          <div className="w-full p-6">
            <ScrollArea className="h-[500px] rounded-md border">
              {isSuccess ? (
                <BillTable data={data.data.resultado as Bill[]} />
              ) : null}
              {isStale ? (
                <BillTable data={data.data.resultado as Bill[]} />
              ) : null}
            </ScrollArea>
          </div>
          <div className="grid grid-cols-3 place-items-center gap-5">
            <div>
              <Button>
                <Link to="/bills-read">Read bill</Link>
              </Button>
            </div>
            <div>
              <Button>
                <Link to="/bills-update">Update Bill</Link>
              </Button>
            </div>
            <div>
              <Button variant="destructive">
                <Link to="/bills-remove">Delete bill</Link>
              </Button>
            </div>

          </div>
          {/* {isError ? (
            <AlertComp
              error={isError}
              message={"error while processing request"}
              classification="Error"
            />
          ) : null}
          {isPending ? (
            <AlertComp
              error={false}
              message="Loading..."
              classification="Request in transit"
            />
          ) : null}
          {isSuccess ? (
            <AlertComp
              error={false}
              message="Success!"
              classification="Bill sent!"
            />
          ) : null} */}
        </div>
      </div>
    </div>
  );
};
