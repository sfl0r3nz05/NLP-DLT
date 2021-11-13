import React, { useContext, useEffect, useState } from "react";
import { BrowserRouter as Router, Route, Switch, Link } from "react-router-dom";
import { AuthContext } from "../App";
import { routes } from "./routes";
import MyMenu from "../components/menu";
import { Layout, Menu, Icon, Button } from "antd";
const { Header, Content, Footer } = Layout;

const Rutes = ({ history }) => {
  const { state, dispatch } = useContext(AuthContext);
  const selectedUser = state.user;
  const [selectedKey, setSelectedKey] = useState(null)

  useEffect(() => {
    if (!selectedUser.path.includes(history.location.pathname)) onChangeRoute(selectedUser.path[0])
    else onChangeRoute(history.location.pathname);
  }, [])

  //Change route with history
  const onChangeRoute = (path) => {
    history.push(path);
    let new_route = routes.find((route) => route.path === path);
    if (new_route && new_route.key) setSelectedKey(new_route.key);
  }

  const logout = e => {
    e.preventDefault();
    dispatch({ type: "LOGOUT" });
  };

  return (
    <Router>
      <Layout className="app">
        <MyMenu Link={Link} onChangeRoute={onChangeRoute} selectedKey={[selectedKey]} />
        <Layout>
          <Header className="header" style={{ background: "#fff", padding: "0 0%" }}>
            <Menu mode="horizontal" style={{ lineHeight: "64px", textAlign: "right" }}>
              <Menu.Item key="10">
                <Link to={"/user-data"} onClick={() => onChangeRoute("/user-data")}>
                  {selectedUser && selectedUser.username} <Icon type="user" />
                </Link>
              </Menu.Item>
              <Menu.Item key="7">
                <Button type="danger" size="small" ghost onClick={e => logout(e)}>
                  <Icon type="logout" /> Close session
                </Button>
              </Menu.Item>
            </Menu>
          </Header>
          <Content style={{ padding: "0 2%" }}>
            <Layout style={{ padding: "24px 0", margin: "70px 100px", background: "#fff" }}>
              <Content
                style={{
                  padding: "0 45px",
                  minHeight: 280
                }}
              >
                <Switch location={history.location}>
                  {routes.map(route => selectedUser.path.includes(route.path) && selectedKey && <Route key={route.key} exact={route.exact} path={route.path} component={route.body} />)}
                </Switch>
              </Content>
            </Layout>
          </Content>
          <Footer style={{ textAlign: "center", padding: 3, minHeight: 48, marginTop: 20 }}>&#9400; Linux Foundation 2021 </Footer>
        </Layout>
      </Layout>
    </Router>
  );
};
export default Rutes;
