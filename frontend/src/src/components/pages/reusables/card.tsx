import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "../../ui/card";
import InputForm from "./Form";

export interface CardFormProps {
  alert: string;
  setAlert: React.Dispatch<React.SetStateAction<string>>;
  setError: React.Dispatch<React.SetStateAction<boolean>>;
}

const CardForm: React.FC<CardFormProps> = (props): JSX.Element => {
  const { alert, setAlert, setError } = props;
  return (
    <Card className="w-96 bg-gray-900 h-[32rem]">
      <CardHeader>
        <CardTitle className="text-gray-400 text-center p-2 my-2">
          Login
        </CardTitle>
        <CardDescription className="text-gray-400 text-center p-2 my-2">
          Please input your name and password
        </CardDescription>
        <CardContent>
          <InputForm alert={alert} setAlert={setAlert} setError={setError} />
        </CardContent>
      </CardHeader>
    </Card>
  );
};

export default CardForm;
