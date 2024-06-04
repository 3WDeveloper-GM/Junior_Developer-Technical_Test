import { NavBarItem } from "./navbarURL";

export const NavBar = () => {
  return (
    <div className="grid grid-cols-[2fr_1fr_1fr] bg-accent">
      <div className="flex justify-evenly my-3 py-2">
       <NavBarItem url="/bills-fetch" urlName={"Fetch"} />
       <NavBarItem url="/bills-read" urlName={"Read"} />
       <NavBarItem url="/bills-remove" urlName={"Remove"} />
       <NavBarItem url="/bills-update" urlName={"Update"} />
       <NavBarItem url="/bills-create" urlName={"Create"} />
      </div>
      <div></div>
      <div></div>
    </div>
  );
};
