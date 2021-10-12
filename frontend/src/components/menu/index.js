import { AuthContext } from "../../App";
import React, { useContext } from "react";
import { Layout, Menu, Icon } from "antd";
import { routes } from "../../router/routes";
// Layout
const { Sider } = Layout;

const SiderMenu = ({ Link, onChangeRoute, selectedKey }) => {
  const { state } = useContext(AuthContext);
  const selectedUser = state.user;

  return (
    <Sider
      theme="default"
      breakpoint="md"
      collapsedWidth="0"
      width={216}
      style={{ backgroundColor: "white" }}
    >
      <div className="logo" />
      <Menu
        theme="dark"
        selectedKeys={selectedKey}
        style={{ lineHeight: "64px", height: "calc(100vh - 66px)", width: 215 }}
      >
        {routes.map(
          (route) =>
            route.menuShow && selectedUser.path.includes(route.path) &&(
              <Menu.Item key={route.key}>
                <Icon type={route.icon} />
                {route.text}
                <Link to={route.path} onClick={() => onChangeRoute(route.path)}>{route.text}</Link>
              </Menu.Item>
            )
        )}
      </Menu>
    </Sider>
  );
};

export default SiderMenu;
