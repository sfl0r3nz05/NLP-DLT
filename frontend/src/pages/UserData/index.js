import React, { useState } from "react";
import {
  Col,
  Icon,
  Input,
  Button,
  Form,
  notification,
  Row,
  Spin,
} from "antd";
import { EyeInvisibleOutlined, EyeTwoTone } from '@ant-design/icons';
import axios from "axios";

notification.config({
  placement: "topRight",
  bottom: 50,
  duration: 1.5,
});

const UserData = () => {

  const initialFormState = {
    pass_old: "",
    pass_new: "",
    pass_new2: "",
  };

  const [userinfo, setUserinfo] = useState(initialFormState);
  const [loading, setLoading] = useState(false);

  function handleChange(event) {
    const value = event.target.value;
    setUserinfo({
      ...userinfo,
      [event.target.name]: value
    });
  }

  //Create a news products
  const handleSubmit = (e) => {
    //clearState();
    e.preventDefault();
    setLoading(true);

    const jwtToken = localStorage.getItem("token"); //Set POST request
    axios
      .post(`https://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/userData`, { userinfo }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "CONTRASEÑA CAMBIADA SATISFACTORIAMENTE"
          );
        }
        if (res.status === 201) {
          openNotificationWithIcon(
            "error",
            "FALTAN VALORES PARA REALIZAR LOS CAMBIOS"
          );
        }
        if (res.status === 202) {
          openNotificationWithIcon(
            "error",
            "CONTRASEÑA ANTERIOR ERRÓNEA"
          );
        }
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "LA NUEVA CONTRASEÑA DEBE SER DIFERENTE A LA ANTERIOR"
          );
        }
        if (res.status === 204) {
          openNotificationWithIcon(
            "error",
            "LA NUEVA CONTRASEÑA NO COINCIDE CON LA REPETIDA"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "ERROR AL CAMBIAR CONTRASEÑA",
        )
      )
      .finally(() => setLoading(false));
  };

  const openNotificationWithIcon = (type, title, description) => {
    notification[type]({
      message: title,
      description: description,
    });
  };

  const formItemLayout = {};

  return (
    <section className="CommentsWrapper">
      <h2> INFORMACIÓN DE USUARIO</h2>
      <Row gutter={[16, 16]}>
        <Col xl={24} lg={24} md={24}>
          <Form
            {...formItemLayout}
            name="basic"
            initialvalues={{
              remember: true
            }}
            //onChange={handleSubmit}
            onSubmit={handleSubmit}
          >
            <Spin spinning={loading}>
              <Col lg={8} md={24}>

                <Form.Item
                  label="CONTRASEÑA ANTERIOR:"
                  name="pass_old"
                  rules={[
                    {
                      required: true,
                      message: "Introducir la contraseña anterior"
                    }
                  ]}
                  hasFeedback
                >
                  <Input.Password
                    size="large"
                    prefix={
                      <Icon type="lock" style={{ color: "rgba(0,0,0,.25)" }} />
                    }
                    placeholder={"fGQzOlL2Cv"}
                    iconRender={visible => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
                    type="text"
                    name="pass_old"
                    value={userinfo.pass_old}
                    onChange={handleChange}
                    style={{ width: '55%' }}
                  />
                </Form.Item>
                <Form.Item
                  label="NUEVA CONTRASEÑA:"
                  name="pass_new"
                  rules={[
                    {
                      required: true,
                      message: "Introducir la cantidad de ethers en wei"
                    }
                  ]}
                >
                  <Input.Password
                    size="large"
                    prefix={
                      <Icon type="lock" style={{ color: "rgba(0,0,0,.25)" }} />
                    }
                    placeholder={"mQgN0pfEse"}
                    iconRender={visible => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
                    type="text"
                    name="pass_new"
                    value={userinfo.pass_new}
                    onChange={handleChange}
                    style={{ width: '55%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="REPETIR NUEVA CONTRASEÑA:"
                  name="pass_new2"
                  rules={[
                    {
                      required: true,
                      message: "Repetir la nueva contraseña"
                    }
                  ]}
                >
                  <Input.Password
                    size="large"
                    prefix={
                      <Icon type="lock" style={{ color: "rgba(0,0,0,.25)" }} />
                    }
                    placeholder={"mQgN0pfEse"}
                    iconRender={visible => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
                    type="text"
                    name="pass_new2"
                    value={userinfo.pass_new2}
                    onChange={handleChange}
                    style={{ width: '55%' }}
                  />
                </Form.Item>
                <br />
                <br />
                <br />
                <Form.Item >
                  <Button
                    size="large"
                    type="primary"
                    htmlType="submit"
                    block
                    style={{ width: '55%' }}
                  >
                    CAMBIAR CONTRASEÑA
                  </Button>
                </Form.Item>
              </Col>
              <Col lg={16} md={24}></Col>
            </Spin>
          </Form>
        </Col>
      </Row>
    </section>
  );
};

export default UserData;