import { Link, useRouterState } from "@tanstack/react-router";
import { FC } from "react";

interface navBarItemProps {
  url: string;
  urlName: string;
}

export const NavBarItem: FC<navBarItemProps> = (props) => {
  const { pathname } = useRouterState().location;

  return (
    <>
      {pathname === props.url ? (
        <div className="bg-zinc-500 text-primary-foreground rounded-md px-2 py-2">
          <Link to={props.url}>{props.urlName}</Link>
        </div>
      ) : null}
      {pathname !== props.url ? (
        <div className="px-2 py-2">
          <Link to={props.url}>{props.urlName}</Link>
        </div>
      ) : null}
    </>
  );
};
