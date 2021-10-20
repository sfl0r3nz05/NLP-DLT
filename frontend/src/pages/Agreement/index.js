import React, { useState, useEffect } from "react";
import {
  AutoComplete,
  Row,
  Col,
  Form,
  Input,
  Button,
  Spin,
  notification
} from "antd";
import { Icon as NewIco } from "antd";
import "./../../App.css";
import axios from "axios";
import Control from 'react-leaflet-control';
import Clipboard from 'react-clipboard.js';
import { useGlobal } from "reactn";
import * as lerData from "./../../data/LER.json";
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

const INITIAL_STATE = {
  lat: 43.2612,
  lng: -1.779,
  zoom: 12
}

const Agreement = () => {

  let userDetails = JSON.parse(localStorage.getItem('user'));

  useEffect(() => {
    delete L.Icon.Default.prototype._getIconUrl;
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
      iconUrl: require("leaflet/dist/images/marker-icon.png"),
      shadowUrl: require("leaflet/dist/images/marker-shadow.png")
    });
  }, []);

  const formItemLayout = {};

  const initialFormState = {
    mno1: "",
    mno2: "",
    nameRA: "",
  };

  const [createAgreement, setcreateAgreement] = useState(initialFormState);
  const [mapState, setMapState] = useState(INITIAL_STATE); //Estado por defecto
  const [loading, setLoading] = useState(false);
  const [global] = useGlobal();

  function onChange(value) {
    const index = lerData.LER.findIndex(data => data.name === value)
    const ler = lerData.LER[index].id
    setcreateAgreement(prevValue => ({ ...prevValue, mno2: value })) //console.log(token);
  }

  const onClick = () => {
    const value = global.value;
    createAgreement.mno1 = value;
    setcreateAgreement(prevValue => ({ ...prevValue, mno1: value }))
  }

  function handleChange(event) {
    const value = event.target.value;
    setcreateAgreement({
      ...createAgreement,
      [event.target.name]: value
    });
    //console.log(createAgreement);
  }

  //Create a news products
  const handleSubmit = (e) => {
    console.log(e);
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
            "BLOCKCHAIN INTERACTION ERROR"
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
      <h2> ROAMING AGREEMENT CREATION</h2>
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
                    placeholder={"MNO to propose the Roaming Agreement"}
                    style={{ width: '100%' }}
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
                    placeholder={"E.g.: Proximus Roaming Agreement"}
                    type="text"
                    name="nameRA"
                    value={createAgreement.nameRA}
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