import axios from "axios";
import { LoginForm } from "./src/components/pages/login";

const App = () => {
  axios.defaults.withCredentials = true;
  setTimeout(
    () =>
      axios
        .get("http://localhost:4040/public/healthCheck", {
          withCredentials: true,
        })
        .then((response) => console.log(response))
        .catch((error) => console.log(error.response.data.error)),
    1000,
  );
  return <LoginForm />;
};

export default App;
