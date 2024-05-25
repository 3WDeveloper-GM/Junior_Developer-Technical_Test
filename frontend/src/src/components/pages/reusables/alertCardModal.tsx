import { AlertCircle, Terminal } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "../../ui/alert";

interface alertCardProps {
  text: string;
  error: boolean;
}

const AlertCard: React.FC<alertCardProps> = (props): JSX.Element => {
  const { text, error } = props;
  switch (true) {
    case text != "":
      if (error == true) {
        return (
          <>
            <Alert variant={"destructive"}>
              <AlertCircle className="w-4 h-4" />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>{text}</AlertDescription>
            </Alert>
          </>
        );
      } else {
        return (
          <>
            <Alert>
              <Terminal className="w-4 h-4" />
              <AlertTitle>Success!</AlertTitle>
              <AlertDescription>{text}</AlertDescription>
            </Alert>
          </>
        );
      }
    default:
      return <></>;
  }
};

export default AlertCard;
