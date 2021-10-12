//import jwt_decode from "jwt-decode";
//import { BehaviorSubject } from "rxjs";
import React, { useContext, useState } from "react";
import { Form, Icon, Input, Button, Layout, Row, Col, Alert, Spin } from "antd";
import { AuthContext } from "../../App";
import logoHyperledger from "./../../assets/img/hyperledger.png";

const { Header, Content, Footer } = Layout;

const Login = () => {
  const { dispatch } = useContext(AuthContext);
  const initialState = {
    name: "",
    password: "",
    isSubmitting: false,
    errorMessage: null,
  };

  const [data, setData] = useState(initialState);

  const handleInputChange = (event) => {
    setData({
      ...data,
      [event.target.name]: event.target.value,
    });
  };

  const handleFormSubmit = (event) => {
    event.preventDefault();
    setData({
      ...data,
      isSubmitting: true,
      errorMessage: null,
    });

    fetch(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/authentication`, {
      method: "post",
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/x-www-form-urlencoded',
        'Content-Length': 1108
      },
      body: JSON.stringify({
        username: data.username,
        password: data.password,
      }),
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        }
        throw res;
      })
      .then((resJson) => {
        dispatch({
          type: "LOGIN",
          payload: resJson,
        });
      })
      .catch((error) => {
        let errorText =
          error.status === 404
            ? "Usuario incorrecto"
            : error.status === 401
              ? "Contraseña incorrecta"
              : "No se pudo realizar la petición";
        setData({
          ...data,
          isSubmitting: false,
          errorMessage: errorText,
        });
      });
  };

  return (
    <div>
      <Layout
        className="layout"
        style={{ minHeight: "100vh", background: "#fff" }}
      >
        <Header className="header-login" style={{ padding: 0 }}>
          {/*<div className="logo-login">
            <img src={logoCmobileWhite} alt="Logo" style={{ height: "100%", marginLeft: "10px" }} />
          </div>*/}
        </Header>
        <Content>
          <Row
            style={{ background: "#fff", minHeight: "calc(100vh - 140px)" }}
            type="flex"
            justify="space-around"
            align="middle"
          >
            <Col sm={24} md={12} lg={6} xl={6} style={{ textAlign: "center" }}>
              <Form
                onSubmit={handleFormSubmit}
                className="login-form"
                style={{ padding: 24 }}
              >
                <img
                  src={logoHyperledger}
                  alt="Logo"
                  style={{ height: "20vh", marginBottom: "20px" }}
                />
                <Form.Item>
                  <Input
                    size="large"
                    prefix={
                      <Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />
                    }
                    placeholder="Username"
                    name="username"
                    onChange={handleInputChange}
                  />
                </Form.Item>
                <Form.Item>
                  <Input
                    size="large"
                    prefix={
                      <Icon type="lock" style={{ color: "rgba(0,0,0,.25)" }} />
                    }
                    type="password"
                    placeholder="Password"
                    name="password"
                    onChange={handleInputChange}
                  />
                </Form.Item>
                <Form.Item>
                  <Spin spinning={data.isSubmitting}>
                    <Button
                      size="large"
                      type="primary"
                      htmlType="submit"
                      className="login-form-button"
                    >
                      Entrar
                    </Button>
                  </Spin>
                  {data.errorMessage && (
                    <Alert message={data.errorMessage} type="error" showIcon />
                  )}
                </Form.Item>
              </Form>
            </Col>
          </Row>
        </Content>
        <Footer style={{ textAlign: "center" }}>
          Linux Foundation 2021 Created by sfigueroa
        </Footer>
      </Layout>
    </div>
  );
};

export default Login;
