import { Link, useLocation } from "react-router-dom";
import { Sidebar, Menu, MenuItem } from "react-pro-sidebar";
import { useEffect, useState } from "react";
import { FaHome, FaRegQuestionCircle } from "react-icons/fa";
import { FaGear } from "react-icons/fa6";
import { MdOutlineStorage } from "react-icons/md";
import { IoDocumentTextOutline, IoMenu } from "react-icons/io5";

const activeTab = {
  home: false,
  settings: false,
  caching: false,
  tutorial: false,
  resources: false,
};

const Navigation = () => {
  const [collapsed, setCollapsed] = useState(false);
  const [toggled, setToggled] = useState(false);

  const collapseSidebar = () => {
    setCollapsed(!collapsed);
  };

  let location = useLocation();
  useEffect(() => {
    console.log(location.pathname);
    switch (location.pathname) {
      case "/":
        activeTab.home = true;
        break;
      case "/settings":
        activeTab.settings = true;
        break;
      case "/caching":
        activeTab.caching = true;
        break;
      case "/tutorial":
        activeTab.tutorial = true;
        break;
      case "/Resources":
        activeTab.resources = true;
        break;
      default:
        break;
    }
  }, [location]);

  return (
    <Sidebar
      width={"250px"}
      collapsedWidth={"60px"}
      collapsed={collapsed}
      rootStyles={{
        height: "100vh",
        zIndex: 1000,
      }}
      backgroundColor={"rgba(0, 0, 0, 0.10)"}
      className={
        "bg-black bg-opacity-30 border-r-2 border-gray-300 text-gray-900"
      }
      style={{ height: "100vh" }}
    >
      <Menu
        iconShape="square"
        menuItemStyles={{
          button: ({ active }) => {
            return {
              backgroundColor: active ? "white" : "transparent",
              border: "none",
              padding: "0.75rem",
              width: "100%",
            };
          },
        }}
      >
        <div className={"font-semibold text-xl flex justify-center my-2"}>
          {collapsed ? (
            <IoMenu
              className={
                "text-4xl transition hover:bg-black hover:bg-opacity-10 rounded-md p-0.5"
              }
              onClick={collapseSidebar}
            />
          ) : (
            <div className={"flex gap-2"}>
              <h3>P2P Web Cache</h3>
              <IoMenu
                className={
                  "text-4xl transition hover:bg-black hover:bg-opacity-10 rounded-md p-0.5"
                }
                onClick={collapseSidebar}
              />
            </div>
          )}
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
