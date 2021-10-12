import React, { useState, useEffect } from "react";

import {
  AutoComplete,
  Button,
  Col,
  Form,
  Input,
  notification,
  Row,
  Spin,
  Tooltip
} from "antd";
import { Icon as NewIco } from "antd";
import Control from 'react-leaflet-control';
import axios from "axios";
import "./../../App.css";
import "./../../react-leaflet.css";
import * as locData from "./../../data/companies.json";
import Clipboard from 'react-clipboard.js';
import { useGlobal } from "reactn";

//Leaflet
//----------------------------------------------------------------------------------------------
import MarkerClusterGroup from 'react-leaflet-markercluster';
import { Map, TileLayer, Marker, Popup } from 'react-leaflet'
import 'react-leaflet-markercluster/dist/styles.min.css';
import 'leaflet/dist/leaflet.css';
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

const TransferToken = () => {

  useEffect(() => {
    delete L.Icon.Default.prototype._getIconUrl;
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
      iconUrl: require("leaflet/dist/images/marker-icon.png"),
      shadowUrl: require("leaflet/dist/images/marker-shadow.png")
    });
  }, []);

  const initialFormState = {
    tokenid: "",
    idParticipant_receiver: "",
    msg_value: "",
    location: "",
  };

  const [mapState, setMapState] = useState(INITIAL_STATE); //Estado por defecto
  const [transftoken, setTransftoken] = useState(initialFormState);
  const [loading, setLoading] = useState(false); //See on sending form data
  const [global] = useGlobal();

  function onChange(value) {
    transftoken.idParticipant_receiver = value;
  }

  function handleChange(event) {
    const value = event.target.value;
    setTransftoken({
      ...transftoken,
      [event.target.name]: value
    });
  }

  const handleClick = () => {
    const value = global.value;
    transftoken.tokenid = value;
    setTransftoken(prevValue => ({ ...prevValue, tokenid: value })) //console.log(transftoken);
  }

  const onClick = () => {
    const value = global.value;
    transftoken.msg_value = value;
    setTransftoken(prevValue => ({ ...prevValue, msg_value: value })) //console.log(transftoken);
  }

  //Create a news products
  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)

    const jwtToken = localStorage.getItem("token");
    axios     //Set POST request
      .post(`https://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/transftoken`, { transftoken }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "VALOR/TOKEN TRANSFERIDO SATISFACTORIAMENTE"
          );
        }
        if (res.status === 201) {
          openNotificationWithIcon(
            "error",
            "FALTAN VALORES PARA TRANSFERIR EL TOKEN"
          );
        }
        if (res.status === 202) {
          openNotificationWithIcon(
            "error",
            "CONTENEDOR NO DESPLEGADO PARA ESTA ENTIDAD"
          );
        }
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "ERROR DE INTERACCIÓN CON LA BLOCKCHAIN, REVISAR LA INFORMACIÓN QUE SE ENVÍA"
          );
        }
        if (res.status === 204) {
          openNotificationWithIcon(
            "error",
            "TRANSFERENCIA REALIZADA A LA MISMA ENTIDAD POSEEDORA DEL TOKEN"
          );
        }
        if (res.status === 205) {
          openNotificationWithIcon(
            "error",
            "ENTIDAD DEFINIDA SÓLO COMO PRODUCTOR O NO DEFINIDA COMO VENDEDORA"
          );
        }
        if (res.status === 206) {
          openNotificationWithIcon(
            "error",
            "LA LOCALIZACIÓN DEBE COINCIDIR CON LA COMPAÑIA"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "No se ha podido transferir el token",
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
      <h2> TRANSFERIR EL TOKEN</h2>
      <Row gutter={[16, 16]}>
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
                  label="NOMBRE DEL TOKEN A TRANSFERIR:"
                  name="tokenid"
                  rules={[
                    {
                      required: true,
                      message: "Introducir el nombre del Token"
                    }
                  ]}
                  hasFeedback
                >
                  <Input
                    size="large"
                    placeholder={"Nombre del Token a transferir"}
                    suffix={
                      <Clipboard onClick={handleClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Pegar el nombre del Token">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="tokenid"
                    style={{ width: '95%' }}
                    value={transftoken.tokenid}
                    onChange={handleChange}
                  />

                </Form.Item>
                <Form.Item
                  label="NOMBRE DE LA ENTIDAD A TRANSFERIR EL TOKEN:"
                  name="idParticipant_receiver"
                  rules={[
                    {
                      required: true,
                      message: "Introducir la cantidad de ethers en wei"
                    }
                  ]}
                >
                  <AutoComplete
                    size="large"
                    dataSource={locData.companies.map(data => data.name)}
                    placeholder={"Seleccionar la compañía"}
                    style={{ width: '95%' }}
                    onSelect={(data) => data}
                    onChange={onChange}
                  >
                  </AutoComplete>
                </Form.Item>

                <Form.Item
                  label="CANTIDAD DE ETHERS (WEI) A PAGAR POR TOKEN:"
                  name="msg_value"
                  rules={[
                    {
                      required: true,
                      message: "Introducir la cantidad de ethers en wei"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder={"100 Wei"}
                    suffix={
                      <Clipboard onClick={onClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Pegar el precio del Token">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="msg_value"
                    value={transftoken.msg_value}
                    onChange={handleChange}
                    style={{ width: '30%' }}
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
                    style={{ width: '95%' }}
                  >
                    TRANSFERIR TOKEN
                       </Button>
                </Form.Item>
              </Col>

              <Col lg={16} md={24}>

                <Form.Item
                  label="LOCALIZAR ENTIDAD DONDE SE ENVIARÁ EL TOKEN:"
                  name="location"
                >
                  <Map
                    center={[mapState.lat, mapState.lng]}
                    zoom={mapState.zoom}
                    style={{ height: 500 }}
                  >
                    <TileLayer
                      attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                      url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                    />
                    <Control position="topright" >
                      <AutoComplete
                        style={{ width: 200 }}
                        size="large"
                        dataSource={locData.companies.map(data => data.name)}
                        placeholder="Nombre empresa"
                        onSelect={(data) => {
                          const comp = locData.companies.find(company => company.name === data)
                          setTransftoken(prevValue => ({ ...prevValue, location: comp.geometry.coordinates }))
                          setMapState({ ...mapState, lat: comp.geometry.coordinates[1], lng: comp.geometry.coordinates[0], zoom: 14 });
                        }}
                      />
                    </Control>
                    <MarkerClusterGroup>
                      {locData.companies.map(location => (
                        <Marker
                          key={location.id}
                          position={[
                            location.geometry.coordinates[1],
                            location.geometry.coordinates[0]
                          ]}
                          onClick={() => {
                            setTransftoken(prevValue => ({ ...prevValue, location: location.geometry.coordinates }))
                          }}
                        >
                          <Popup>
                            <div>
                              <h2><a target="popup" href={'https://www.google.com/maps/search/?api=1&query=' + location.geometry.coordinates[1] + "," + location.geometry.coordinates[0]}>{location.name}</a></h2>
                              <a style={{ fontSize: 14 }} target="popup" href={'http://' + location.web}>{location.web}</a>
                              <p style={{ fontSize: 13 }}>{location.description}</p>
                            </div>
                          </Popup>
                        </Marker>
                      ))}
                    </MarkerClusterGroup>
                  </Map>
                </Form.Item>
              </Col>

            </Spin>
          </Form>
        </Col>
      </Row>
    </section>
  );
};

export default TransferToken;