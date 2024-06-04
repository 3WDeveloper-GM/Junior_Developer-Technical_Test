import { Bill, emptyBill } from "@/type/bills";
import {
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { FC, useContext, useEffect, useState } from "react";
import { FetchColumns } from "./columns";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../ui/table";
import { useBill } from "@/hooks/billHooks";
import { CurrentBillContext } from "@/contexts/billCtx";
import { DateTime } from "luxon";

interface BillTableProps {
  data: Bill[];
}

export const BillTable: FC<BillTableProps> = (props) => {
  const [rowSelection, setRowSelection] = useState({});
  const { data } = props;
  const table = useReactTable<Bill>({
    columns: FetchColumns,
    data: data,
    state: {
      rowSelection,
    },
    onRowSelectionChange: setRowSelection,
    enableMultiRowSelection: false,
    getRowId: (row) => row.idFactura,
    getCoreRowModel: getCoreRowModel(),
  });

  const currentBill = data.find(
    (element) => element.idFactura === Object.keys(rowSelection)[0],
  );

  const ctx = useContext(CurrentBillContext);
  useEffect(() => {
    if (typeof currentBill !== "undefined") {
      localStorage.setItem(
        "billDate",
        DateTime.fromISO(currentBill?.fechaEmision).toFormat("yyyy-MM-dd"),
      );

      ctx?.updateCurrentBill(currentBill);
    }
  });

  return (
    <Table>
      <TableHeader>
        {table.getHeaderGroups().map((headerGroup) => (
          <TableRow key={headerGroup.id}>
            {headerGroup.headers.map((header) => (
              <TableHead key={header.id}>
                {header.isPlaceholder
                  ? null
                  : flexRender(
                      header.column.columnDef.header,
                      header.getContext(),
                    )}
              </TableHead>
            ))}
          </TableRow>
        ))}
      </TableHeader>
      <TableBody>
        {table.getRowModel().rows.map((row) => (
          <TableRow key={row.id}>
            {row.getVisibleCells().map((cell) => (
              <TableCell key={cell.id}>
                {flexRender(cell.column.columnDef.cell, cell.getContext())}
              </TableCell>
            ))}
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
