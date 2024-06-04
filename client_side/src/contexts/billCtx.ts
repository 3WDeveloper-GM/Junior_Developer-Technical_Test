import { Bill } from "@/type/bills";
import { createContext } from "react";

export interface BillContextType {
  bill: Bill;
  updateCurrentBill: (newBill: Bill) => void;
}

export const CurrentBillContext = createContext<BillContextType | null>(null);
