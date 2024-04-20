import { Link } from "react-router-dom";
import { Sidebar, Menu, MenuItem } from "react-pro-sidebar";
import { useState } from "react";
import { FaHome, FaRegQuestionCircle } from "react-icons/fa";
import { FaGear } from "react-icons/fa6";
import { MdOutlineStorage } from "react-icons/md";
import { IoDocumentTextOutline } from "react-icons/io5";

const Navigation = () => {
  const [collapsed, setCollapsed] = useState(false);
  const [toggled, setToggled] = useState(false);

  return (
    <Sidebar
      breakPoint={"md"}
      collapsed={collapsed}
      rootStyles={{
        height: "100vh",
        width: "200px",
        zIndex: 1000,
      }}
      backgroundColor={"rgba(0, 0, 0, 0.10)"}
      onToggle={() => setCollapsed(!collapsed)}
      toggled={toggled}
      toggle={setToggled}
      className={
        "bg-black bg-opacity-30 border-r-2 border-gray-300 text-gray-900"
      }
      style={{ height: "100vh" }}
    >
      <Menu iconShape="square">
        <div className={"font-semibold text-xl flex justify-center my-2"}>
          <h3>P2P Web Cache</h3>
        </div>
        <MenuItem
          icon={<FaHome />}
          component={<Link to={"/"} className={"nav-item"} />}
        >
          Home
        </MenuItem>
        <MenuItem
          icon={<FaGear />}
          component={<Link to={"/settings"} className={"nav-item"} />}
        >
          Settings
        </MenuItem>
        <MenuItem
          icon={<MdOutlineStorage />}
          component={<Link to={"/caching"} className={"nav-item"} />}
        >
          Cache
        </MenuItem>
        <MenuItem
          icon={<FaRegQuestionCircle />}
          component={<Link to={"/tutorial"} className={"nav-item"} />}
        >
          Tutorial
        </MenuItem>
        <MenuItem
          icon={<IoDocumentTextOutline />}
          component={<Link to={"/Resources"} className={"nav-item"} />}
        >
          Resources
        </MenuItem>
      </Menu>
    </Sidebar>
  );
};

export default Navigation;
