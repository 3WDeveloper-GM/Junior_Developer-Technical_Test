import { useMutation } from "@tanstack/react-query";
import {  updateBill } from "@/api/bills";
import { AlertComp } from "@/components/common/AlertComponent";
import { useContext } from "react";
import { CurrentBillContext } from "@/contexts/billCtx";
import { UpdateBillForm } from "@/components/billUpdateForm";

export const BillUpdateView = () => {
  const { mutateAsync, isError, isPending, isSuccess } = useMutation({
    mutationFn: updateBill,
  });

  const ctx = useContext(CurrentBillContext)
  const data = ctx?.bill

  return (
    <div>
      <div className="p-20  w-full lg:grid lg:min-h-[600px] lg:grid-cols-[2fr_1fr] xl:min-h-[800px] place-items-center">
        <div className="flex items-center justify-center py-12">
          <div className="mx-auto grid gap-6">
            <div className="grid gap-2 text-center place-items-center">
              <h1 className="text-3xl font-bold">Update Bill</h1>
              <p className="text-balance text-muted-foreground">
                Update the Bill in your database
              </p>
              <UpdateBillForm mutate={mutateAsync} previousData={data}/>
            </div>
          </div>
        </div>
        <div className="w-64">
          {isError ? (
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
          ) : null}
        </div>
      </div>
    </div>
  );
}
