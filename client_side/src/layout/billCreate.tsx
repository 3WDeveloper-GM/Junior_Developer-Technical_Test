import { useMutation } from "@tanstack/react-query";
import { BillForm } from "@/components/billCreateForm";
import { createBill } from "@/api/bills";
import { AlertComp } from "@/components/common/AlertComponent";

export const BillCreateView = () => {
  const { mutateAsync, isError, isPending, isSuccess } = useMutation({
    mutationFn: createBill,
  });

  return (
    <div>
      <div className="p-20  w-full lg:grid lg:min-h-[600px] lg:grid-cols-[2fr_1fr] xl:min-h-[800px] place-items-center">
        <div className="flex items-center justify-center py-12">
          <div className="mx-auto grid gap-6">
            <div className="grid gap-2 text-center place-items-center">
              <h1 className="text-3xl font-bold">Create Bill</h1>
              <p className="text-balance text-muted-foreground">
                Add a new Bill to the Database
              </p>
              <BillForm mutate={mutateAsync} />
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
};
