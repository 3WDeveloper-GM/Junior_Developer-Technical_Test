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
import { AlertComp } from "./common/AlertComponent";
import { Button } from "./ui/button";

interface BillReadProps {
  content: Bill | undefined;
  refetch: (
    options?: RefetchOptions | undefined,
  ) => Promise<QueryObserverResult<AxiosResponse<unknown, any>, Error>>;
}

export const BillRemoveLayout: FC<BillReadProps> = (props) => {
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

  const handleClick = (e) => {
    e.preventDefault()
    props.refetch()
  }

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
          <TableBody className="text-balance">
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

      <div className="text-center my-8">
        <AlertComp
          classification="Warning"
          message="Are you sure that you want to delete this Bill?"
          error={true}
        />
      </div>
      <div className="text-center my-8">
        <Button
          variant="destructive"
          type="submit"
          onClick={handleClick}
        >
          Delete
        </Button>
      </div>
    </div>
  );
};
