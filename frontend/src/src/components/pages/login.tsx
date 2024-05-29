import { useState } from "react";

import CardForm from "./reusables/card";
import AlertCard from "./reusables/alertCardModal";

export function LoginForm() {


  const [alertText, setAlert] = useState("");
  const [isError, setIsError] = useState(false);
  return (
    <div className="grid grid-cols-[12fr_29fr_12fr] grid-rows-[29fr_70fr_29fr] h-[100vh] gap-3 place-items-center">
      <div></div>
      <div>
        <AlertCard text={alertText} error={isError} />
      </div>
      <div></div>

      <div></div>
      <div>
        <CardForm alert={alertText} setAlert={setAlert} setError={setIsError} />
      </div>
      <div></div>

      <div></div>
      <div></div>
      <div></div>
    </div>
  );
}
