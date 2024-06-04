import { Bill } from "@/type/bills";
import api from "./api";

export async function fetchDate({ queryKey }: { queryKey: unknown[] }) {
  const [_key, { header, endDate, startDate }] = queryKey;
  return api.get("/public/bills/fetch", {
    params: {
      endDate: endDate,
      startDate: startDate,
    },
    headers: { Authorization: "Bearer " + header },
    withCredentials: true,
  });
}



export async function deleteBill({queryKey} : {queryKey: unknown[]}) {
  const [_key, {header, BillID}] = queryKey
  return api.delete(`public/bills/delete/${BillID}`, {
    headers: {Authorization: "Bearer " + header},
    withCredentials: true, 
  })
}


export type createBillParams = {
  bill: Omit<Bill, "proveedor">;
  token: string;
};

export const createBill = (Params: createBillParams) => {
  const { bill, token } = Params;
  return api.post("/public/bills/create", bill, {
    withCredentials: true,
    headers: { Authorization: "Bearer " + token },
  });
};

export const updateBill = (Params: createBillParams) => {
  const { bill, token } = Params;
  return api.put("/public/bills/update", bill, {
    withCredentials: true,
    headers: { Authorization: "Bearer " + token },
  });
};

