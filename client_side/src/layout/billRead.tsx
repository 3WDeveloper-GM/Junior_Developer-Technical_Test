import { BillReadLayout } from "@/components/billReadLayout";
import { CurrentBillContext } from "@/contexts/billCtx";
import { emptyBill } from "@/type/bills";
import { useContext } from "react";

export const BillsReadView = () => {
  const currentBill = useContext(CurrentBillContext);
  const currBill = currentBill?.bill;

  return (
    <div className="grid grid-cols-[3fr_1fr] place-items-center h-[600px] p-12">
      {!currBill ? (
        <BillReadLayout content={emptyBill} />
      ) : (
        <BillReadLayout content={currBill} />
      )}
    </div>
  );
};
