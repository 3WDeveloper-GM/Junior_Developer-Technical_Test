import { FC } from "react";
import { Alert, AlertDescription, AlertTitle } from "../ui/alert";
import { AlertCircle, Terminal } from "lucide-react";

interface alertProps {
  classification: string;
  message: string;
  error: boolean;
}

export const AlertComp: FC<alertProps> = (props) => {
  return (
    <>
      <Alert variant={props.error ? "destructive" : "default"}>
        {props.error ? (
          <AlertCircle className="h-4 w-4" />
        ) : (
          <Terminal className="h-4 w-4" />
        )}
        <AlertTitle>{props.classification}</AlertTitle>
        <AlertDescription>{props.message}</AlertDescription>
      </Alert>
    </>
  );
};
