import { Link } from "react-router-dom";

const Navigation = () => {
  return (
    <nav className={"h-screen overflow-y-auto border-r-2 border-r-white p-2"}>
      <aside className="flex flex-col overflow-y-auto">
        <ul className={"flex flex-col space-y-2"}>
          <li>
            <Link to="/" className={"nav-item"}>
              Home
            </Link>
          </li>
          <li>
            <Link to="/settings" className={"nav-item"}>
              Settings
            </Link>
          </li>
          <li>
            <Link to="/caching" className={"nav-item"}>
              Caching
            </Link>
          </li>
          <li>
            <Link to="/tutorial" className={"nav-item"}>
              Tutorial
            </Link>
          </li>
          <li>
            <Link to="/resources" className={"nav-item"}>
              Resources
            </Link>
          </li>
        </ul>
      </aside>
    </nav>
  );
};

export default Navigation;
