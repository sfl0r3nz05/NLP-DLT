import React, { useState } from "react";
import {
  AutoComplete,
  Row,
  Col,
  Form,
  Input,
  Button,
  Spin,
  notification,
} from "antd";
import "./../../App.css";
import axios from "axios";
import * as lerData from "../../data/COUNTRY.json";

notification.config({
  placement: "topRight",
  bottom: 50,
  duration: 1.5,
});

const initialFormState = {
  mno_name: "",
  mno_country: "",
  mno_network: "",
};

const RegisterMNO = () => {

  const formItemLayout = {};

  const [loading, setLoading] = useState(false);
  const [feature, setFeature] = useState(initialFormState);
  let userDetails = JSON.parse(localStorage.getItem('user'));

  function onChange(value) {
    setFeature(prevValue => ({ ...prevValue, mno_country: value }))
  }

  function handleChange(event) {
    const value = event.target.value;
    setFeature({
      ...feature,
      [event.target.name]: value
    });
  }

  function handleSubmit(event) {
    event.preventDefault();
    setLoading(true)
    const jwtToken = localStorage.getItem("token");
    //Set POST request
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/addOrg`, { feature, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "MNO SUCCESSFULLY REGISTERED"
          );
        }
        if (res.status === 201) {
          openNotificationWithIcon(
            "error",
            "MISSING VALUES TO REGISTER THE MNO"
          );
        }
        if (res.status === 202) {
          openNotificationWithIcon(
            "error",
            "CONTAINER NOT DEPLOYED FOR THIS ENTITY"
          );
        }
        if (res.status === 403) {
          openNotificationWithIcon(
            "error",
            "IDENTITY ALREADY REGISTERED"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "IDENTITY ALREADY REGISTERED",
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

  return (
    <section className="CommentsWrapper">
      <h2> REGISTER MOBILE NETWORK OPERATOR</h2>
      <Row gutter={[16, 16]} type="flex">
        <Col xl={24} lg={24} md={24}>
          <Form
            {...formItemLayout}
            name="basic"
            initialvalues={{
              remember: true
            }}
          //onSubmit={handleSubmit}          
          >
            <Spin spinning={loading}>
              <Col lg={1} md={24}></Col>
              <Col lg={12} md={24}>
                <Form.Item
                  label="MOBILE NETWORK OPERATOR NAME:"
                  name="mno_name"
                  rules={[
                    {
                      required: true,
                      message: "MOBILE NETWORK OPERATOR NAME"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder="E.g.: TELEFONICA"
                    type="text"
                    name="mno_name"
                    style={{ width: '80%' }}
                    value={feature.mno_name}
                    onChange={handleChange}
                  />
                </Form.Item>

                <Form.Item
                  label="COUNTRY OF MOBILE NETWORK OPERATOR:"
                  name="mno_country"
                  rules={[
                    {
                      required: true,
                      message: "COUNTRY OF MOBILE NETWORK OPERATOR"
                    }
                  ]}
                >
                  <AutoComplete
                    size="large"
                    dataSource={lerData.LER.map(data => data.name)}
                    placeholder={"SELECT THE COUNTRY OF THE MNO"}
                    style={{ width: '80%' }}
                    onSelect={(data) => data}
                    onChange={onChange}
                  />
                </Form.Item>

                <Form.Item
                  label="NETWORK INFORMATION:"
                  name="mno_network"
                  rules={[
                    {
                      required: true,
                      message: "NETWORK INFORMATION"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder="E.g.: MOVISTAR"
                    type="text"
                    name="mno_network"
                    style={{ width: '80%' }}
                    value={feature.mno_network}
                    onChange={handleChange}
                  />
                </Form.Item>

                <br />
                <br />

                <Form.Item>
                  <Button
                    size="large"
                    type="primary"
                    htmlType="submit"
                    block
                    onClick={e => {
                      handleSubmit(e)
                    }}
                    style={{ width: '80%' }}
                  >
                    REGISTER MNO
                  </Button>
                </Form.Item>
              </Col>
              <Col lg={11} md={24}>
              </Col>
            </Spin>
          </Form>
        </Col>
      </Row>
    </section>
  );
};

export default RegisterMNO;