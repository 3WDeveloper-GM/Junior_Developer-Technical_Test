import { Bill } from "@/type/bills";
import { Checkbox } from "@/components/ui/checkbox";
import { createColumnHelper } from "@tanstack/react-table";

const colHelper = createColumnHelper<Bill>();

export const FetchColumns = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => {
          row.toggleSelected(!!value);
        }}
        aria-label="Select row"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  colHelper.accessor((row) => row.numeroRegistro, {
    id: "Registro",
    cell: (info) => info.getValue(),
    header: () => <span> Register </span>
  }),
  colHelper.accessor((row) => row.idFactura, {
    id: "billID",
    cell: (info) => info.getValue(),
    header: () => <span>Bill ID</span>,
  }),
  colHelper.accessor((row) => row.proveedor.identificacion, {
    id: "Provider ID",
    cell: (info) => info.getValue(),
    header: () => <span>Provider ID</span>,
  }),
  colHelper.accessor((row) => row.proveedor.nombre, {
    id: "Provider Name",
    cell: (info) => info.getValue(),
    header: () => <span>Provider Name</span>,
  }),
  colHelper.accessor((row) => row.montoTotal, {
    id: "Total Amount",
    cell: (info) => info.getValue(),
    header: () => <span>Total Amount</span>,
    enableSorting: true,
  }),
  colHelper.accessor((row) => row.fechaEmision, {
    id: "Emission Date",
    cell: (info) => info.getValue(),
    header: () => <span>Emission Date</span>,
  }),
];
