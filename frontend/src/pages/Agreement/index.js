import React, { useState } from "react";
import {
  AutoComplete,
  Row,
  Col,
  Form,
  Input,
  Button,
  notification,
  Spin
} from "antd";
import "./../../App.css";
import axios from "axios";
import * as lerData from "./../../data/LER.json";

notification.config({
  placement: "topRight",
  bottom: 50,
  duration: 1.5,
});

const Agreement = () => {

  const initialFormState = {
    mno1: "",
    mno2: "",
    nameRA: "",
  };

  const formItemLayout = {};
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const [createAgreement, setcreateAgreement] = useState(initialFormState);
  const [loading, setLoading] = useState(false);

  function onChange(value) {
    setcreateAgreement(prevValue => ({ ...prevValue, mno2: value }))
  }

  function handleChange(event) {
    const value = event.target.value;
    setcreateAgreement({
      ...createAgreement,
      [event.target.name]: value
    });
  }

  const openNotificationWithIcon = (type, title, description) => {
    notification[type]({
      message: title,
      description: description,
    });
  };

  //Create a news products
  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)
    const jwtToken = localStorage.getItem("token");
    //Set POST request
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/proposeAgreementInitiation`, { createAgreement, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "SUCCESSFULLY REGISTERED AGREEMENT"
          );
        }
        if (res.status === 201) {
          openNotificationWithIcon(
            "error",
            "MISSING VALUES TO CREATE THE AGREEMENT"
          );
        }
        if (res.status === 202) {
          openNotificationWithIcon(
            "error",
            "ROAMING AGREEMENT MUST BE CREATED BETWEEN TWO MNOs"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "UNREGISTERED ROAMING AGREEMENT",
        )
      )
      .finally(() => setLoading(false));
  };

  return (
    <section className="CommentsWrapper">
      <Row gutter={[16, 16]} type="flex">
        <h2> ROAMING AGREEMENT CREATION</h2>
        <Col xl={24} lg={24} md={24}>
          <Form
            {...formItemLayout}
            name="basic"
            initialvalues={{
              remember: true
            }}
          >
            <Spin spinning={loading}>
              <Col lg={1} md={24}></Col>
              <Col lg={12} md={24}>
                <Form.Item
                  label="NAME OF THE SECOND MOBILE NETWORK OPERATOR:"
                  name="mno2"
                  rules={[
                    {
                      required: true,
                      message: "Introducir identificador único del lote"
                    }
                  ]}
                >
                  <AutoComplete
                    size="large"
                    dataSource={lerData.LER.map(data => data.name)}
                    placeholder={"E.g.: ORANGE"}
                    style={{ width: '80%' }}
                    onSelect={(data) => data}
                    onChange={onChange}
                  >
                  </AutoComplete>
                </Form.Item>

                <Form.Item
                  label="NAME OF THE ROAMING AGREEMENT"
                  name="nameRA"
                  rules={[
                    {
                      required: true,
                      message: "Introducir correctamente la métrica"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder={"E.g.: RA001"}
                    type="text"
                    name="nameRA"
                    value={createAgreement.nameRA}
                    onChange={handleChange}
                    style={{ width: '80%' }}
                  />
                </Form.Item>
                <br />
                <br />
                <br />
                <Form.Item>
                  <Button
                    size="large"
                    type="primary"
                    htmlType="submit"
                    block
                    style={{ width: '80%' }}
                    onClick={handleSubmit}
                  >
                    CREATE ROAMING AGREEMENT
                  </Button>
                </Form.Item>
              </Col>
            </Spin>
          </Form>
        </Col>
      </Row>
    </section>
  );
};

export default Agreement;