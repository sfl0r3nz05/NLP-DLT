import React, { useState, useEffect } from "react";
import {
  Row,
  Col,
  Form,
  Input,
  Button,
  Spin,
  notification,
  Tooltip
} from "antd";
import "./../../App.css";
import axios from "axios";
import { Icon as NewIco } from "antd";
import Clipboard from 'react-clipboard.js';
import { useGlobal } from "reactn";
//Leaflet
//----------------------------------------------------------------------------------------------
import 'react-leaflet-markercluster/dist/styles.min.css';
import './../../react-leaflet.css';
import L, { Icon } from 'leaflet';
//----------------------------------------------------------------------------------------------

export const icon = new Icon({
  iconUrl: "./icon/location.svg",
  iconSize: [25, 25]
});

notification.config({
  placement: "topRight",
  bottom: 50,
  duration: 1.5,
});

const AcceptAgreement = () => {

  useEffect(() => {
    delete L.Icon.Default.prototype._getIconUrl;
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
      iconUrl: require("leaflet/dist/images/marker-icon.png"),
      shadowUrl: require("leaflet/dist/images/marker-shadow.png")
    });
  }, []);

  const formItemLayout = {};
  const [global] = useGlobal();
  const initialFormState = {
    mno: "",
  };
  const [loading, setLoading] = useState(false);
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const [acceptAgreement, setacceptAgreement] = useState(initialFormState);

  const onClick = () => {
    const value = global.value;
    acceptAgreement.mno = value;
    setacceptAgreement(prevValue => ({ ...prevValue, tokenid: value }));
  }

  function handleChange(event) {
    const value = event.target.value;
    setacceptAgreement({
      ...acceptAgreement,
      [event.target.name]: value
    });
  }

  //Create a news products
  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)
    const jwtToken = localStorage.getItem("token");
    //Set POST request
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/acceptAgreementInitiation`, { acceptAgreement, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
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

  const openNotificationWithIcon = (type, title, description) => {
    notification[type]({
      message: title,
      description: description,
    });
  };

  return (
    <section className="CommentsWrapper">
      <h2> ACCEPT ROAMING AGREEMENT</h2>
      <Row gutter={[16, 16]} type="flex">
        <Col xl={24} lg={24} md={24}>
          <Form
            {...formItemLayout}
            name="basic"
            initialvalues={{
              remember: true
            }}
            onSubmit={handleSubmit}
          >
            <Spin spinning={loading}>
              <Col lg={8} md={24}>

                <Form.Item
                  label="ROAMING AGREEMENT CLOSED WITH"
                  name="mno"
                  rules={[
                    {
                      required: true,
                      message: "ROAMING AGREEMENT CLOSED WITH"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder={"E.g.: TELEFONICA"}
                    suffix={
                      <Clipboard onClick={onClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Paste MNO Name">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="mno"
                    value={acceptAgreement.mno}
                    onChange={handleChange}
                    style={{ width: '100%' }}
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
                    style={{ width: '100%' }}
                  >
                    ACCEPT ROAMING AGREEMENT
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

export default AcceptAgreement;