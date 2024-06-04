import { Bill } from "@/type/bills";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { FC } from "react";

interface BillReadProps {
  content: Bill | undefined;
}

export const BillReadLayout: FC<BillReadProps> = (props) => {
  const { content } = props;

  const mainTableParams = [
    { id: 0, name: "Bill ID", value: content?.idFactura },
    { id: 1, name: "Total Amount", value: content?.montoTotal },
    { id: 2, name: "Emission Date", value: content?.fechaEmision },
    { id: 3, name: "Provider", value: content?.proveedor.nombre },
    { id: 4, name: "Provider ID", value: content?.proveedor.identificacion },
  ];

  const keyValueDetailEntries = Object.entries(content?.detalles);
  const detailTableParams = keyValueDetailEntries.map(
    ([key, value], index) => ({ key: index, name: key, value: value }),
  );

  const keyValueMiscEntries = Object.entries(content?.miscelaneo);
  const miscTableParams = keyValueMiscEntries.map(([key, value], index) => ({
    key: index,
    name: key,
    value: value,
  }));


  return (
    <div className="p-8 grid auto-rows-auto gap-4 rounded-md w-full place-items-center">
      <div className="bg-accent text-accent-foreground w-full text-center">
        <h5 className="text-xl"> Main Field Information</h5>
      </div>
      <div className="w-11/12">
        <Table>
          <TableHeader>
            <TableRow>
              {mainTableParams.map((item) => (
                <TableHead key={item.id}>{item.name}</TableHead>
              ))}
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              {mainTableParams.map((item) => (
                <TableCell key={item.id + 7}> {item.value} </TableCell>
              ))}
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <div className="bg-accent text-accent-foreground w-full text-center">
        <h3 className="text-xl"> Detail Information </h3>
      </div>

      <div className="w-11/12">
        <Table>
          <TableHeader>
            <TableRow>
              {detailTableParams.map((item) => (
                <TableHead key={item.key}>{item.name}</TableHead>
              ))}
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              {detailTableParams.map((item) => (
                <TableCell key={item.key}> {item.value} </TableCell>
              ))}
            </TableRow>
          </TableBody>
        </Table>
      </div>

      <div className="bg-accent text-accent-foreground w-full text-center">
        <h3 className="text-xl"> Miscelaneous Information </h3>
      </div>

      <div className="w-11/12">
        <Table>
          <TableHeader>
            <TableRow>
              {miscTableParams.map((item) => (
                <TableHead key={item.key}>{item.name}</TableHead>
              ))}
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              {miscTableParams.map((item) => (
                <TableCell key={item.key}> {item.value} </TableCell>
              ))}
            </TableRow>
          </TableBody>
        </Table>
      </div>
    </div>
  );
};
