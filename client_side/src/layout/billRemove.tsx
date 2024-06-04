import { deleteBill } from "@/api/bills";
import { BillRemoveLayout } from "@/components/billRemoveForm";
import { AlertComp } from "@/components/common/AlertComponent";
import { CurrentBillContext } from "@/contexts/billCtx";
import { useUser } from "@/hooks/userHooks";
import { emptyBill } from "@/type/bills";
import { useQuery } from "@tanstack/react-query";
import { useContext } from "react";

export const BillRemoveView = () => {
  const currentBill = useContext(CurrentBillContext);
  const currBill = currentBill?.bill;

  const { authToken } = useUser().getUser();
  const { refetch, error, isLoading, isSuccess, isError } = useQuery({
    queryKey: ["delBill", { header: authToken, BillID: currBill?.idFactura }],
    queryFn: deleteBill,
    enabled: false,
    retry: 0,
  });


  return (
    <div className="grid grid-cols-[2fr_1fr] place-items-center h-[600px] p-8">
      <div>
        {!currBill ? (
          <BillRemoveLayout content={emptyBill} refetch={refetch} />
        ) : (
          <BillRemoveLayout content={currBill} refetch={refetch} />
        )}
      </div>
      <div>
        {isError ? (
          <AlertComp
            error={isError}
            message={error.response.data.error}
            classification="Error"
          />
        ) : null}
        {isLoading ? (
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
            classification="Bill Erased!"
          />
        ) : null}
      </div>
    </div>
  );
};
