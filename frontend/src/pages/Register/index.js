import React, { useState, useEffect } from "react";
import {
  AutoComplete,
  Row,
  Col,
  Form,
  Input,
  Button,
  Card,
  Spin,
  notification,
} from "antd";
import "./../../App.css";
import axios from "axios";
import * as lerData from "./../../data/COUNTRY.json";
//Leaflet
//----------------------------------------------------------------------------------------------
import 'react-leaflet-markercluster/dist/styles.min.css';
import 'leaflet/dist/leaflet.css';
import './../../react-leaflet.css';
import L, { Icon } from 'leaflet';
//----------------------------------------------------------------------------------------------
const { TextArea } = Input;
const { Meta } = Card;

export const icon = new Icon({
  iconUrl: "./icon/location.svg",
  iconSize: [25, 25]
});

notification.config({
  placement: "topRight",
  bottom: 50,
  duration: 1.5,
});

const INITIAL_STATE = {
  lat: 43.2612,
  lng: -1.779,
  zoom: 12
}

const initialFormState = {
  mno_name: "",
  mno_country: "",
  mno_network: "",
};

const Register = () => {

  useEffect(() => {
    delete L.Icon.Default.prototype._getIconUrl;
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
      iconUrl: require("leaflet/dist/images/marker-icon.png"),
      shadowUrl: require("leaflet/dist/images/marker-shadow.png")
    });
  }, []);

  const formItemLayout = {};

  const [loading, setLoading] = useState(false);
  const [mapState, setMapState] = useState(INITIAL_STATE); //Estado por defecto
  const [feature, setFeature] = useState(initialFormState);
  let userDetails = JSON.parse(localStorage.getItem('user'));

  function onChange(value) {
    const index = lerData.LER.findIndex(data => data.name === value)
    const ler = lerData.LER[index].id
    setFeature(prevValue => ({ ...prevValue, mno_country: value })) //console.log(token);
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
              <Col lg={10} md={24}>

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
                    style={{ width: '90%' }}
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
                    style={{ width: '90%' }}
                    onSelect={(data) => data}
                    onChange={onChange}
                  />
                </Form.Item>

                <Form.Item
                  label="NETWORK OF THE MNO:"
                  name="mno_network"
                  rules={[
                    {
                      required: true,
                      message: "NETWORK OF THE MNO"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder="E.g.: MOVISTAR"
                    type="text"
                    name="mno_network"
                    style={{ width: '90%' }}
                    value={feature.mno_network}
                    onChange={handleChange}
                  />
                </Form.Item>
              </Col>

              <Col lg={24} md={24}>
                <Form.Item>
                  <Button
                    size="large"
                    type="primary"
                    htmlType="submit"
                    block
                    onClick={e => {
                      handleSubmit(e)
                    }}
                    style={{ width: '37%' }}
                  >
                    REGISTER MNO
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

export default Register;