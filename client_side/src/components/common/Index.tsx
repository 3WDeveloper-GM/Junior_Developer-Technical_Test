import { useQuery } from "@tanstack/react-query";
import { whoAmI } from "../../api/auth";
import { useUser } from "@/hooks/userHooks";

export const Index = () => {
  const { getUser } = useUser();
  const { authToken } = getUser();
  useQuery({
    queryKey: ["Iam", { header: authToken }],
    queryFn: whoAmI,
  });

  return <></>;
};
