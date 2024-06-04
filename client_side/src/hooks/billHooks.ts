import { Bill, emptyBill } from "@/type/bills";
import { useState } from "react";

export const useBill = () => {
  const [currentBill, setBill] = useState(emptyBill);

  const setBillInfo = (bill: (Bill | undefined)) => {
    if (typeof bill === "undefined") {
      setBill(emptyBill)
      return
    }
    console.log("setting bill info")
    setBill(bill);
  };

  const getCurrentBill = () => {
    return currentBill;
  };

  return { setBillInfo, getCurrentBill };
};

export type BillContext = ReturnType<typeof useBill>;
