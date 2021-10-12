import React, { useState, useEffect, useRef } from "react";
import {
  AutoComplete,
  Row,
  Col,
  Form,
  Input,
  Button,
  Spin,
  notification,
  Tooltip
} from "antd";
import axios from "axios";
import { Icon as NewIco } from "antd";
import Control from 'react-leaflet-control';
import * as lerData from "./../../data/LER.json";
import Clipboard from 'react-clipboard.js';
import { useGlobal } from "reactn";

//Leaflet
//----------------------------------------------------------------------------------------------
import MarkerClusterGroup from 'react-leaflet-markercluster';
import { Map, TileLayer, Marker, Popup } from 'react-leaflet'
import * as locData from "./../../data/companies.json";
import 'react-leaflet-markercluster/dist/styles.min.css';
import 'leaflet/dist/leaflet.css';
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

const CreateToken = () => {

  const initialFormState = {
    //idParticipant: "",
    tokenName: "",
    tokenTTL: "",
    tokenLer: "",
    tokenPrice: "",
    location: "",
  };

  useEffect(() => {
    delete L.Icon.Default.prototype._getIconUrl;
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
      iconUrl: require("leaflet/dist/images/marker-icon.png"),
      shadowUrl: require("leaflet/dist/images/marker-shadow.png")
    });
  }, []);

  const [codLer, setCodLer] = useState(initialFormState.tokenLer);
  const [mapState, setMapState] = useState(INITIAL_STATE); //Estado por defecto
  const [token, setToken] = useState(initialFormState);
  const [loading, setLoading] = useState(false);
  const [global] = useGlobal();
  const markerRef = useRef([]);
  const mapRef = useRef();

  function handleChange(event) {
    const value = event.target.value;
    setToken({
      ...token,
      [event.target.name]: value
    });
  }

  function onChange(value) {
    const index = lerData.LER.findIndex(data => data.name === value)
    const ler = lerData.LER[index].id
    token.tokenLer = ler;
    setCodLer(prevValue => ({ ...prevValue, codLer: ler })) //console.log(token);
  }

  const handleClick = () => {
    const value = global.value;
    token.tokenName = value;
    setToken(prevValue => ({ ...prevValue, tokenName: value })) //console.log(token);
  }

  const onClick = () => {
    const value = global.value;
    token.tokenPrice = value;
    setToken(prevValue => ({ ...prevValue, tokenPrice: value })) //console.log(token);
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)

    const jwtToken = localStorage.getItem("token");
    axios
      .post(`https://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/createToken`, { token }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "TOKEN CREADO SATISFACTORIAMENTE"
          );
        }
        if (res.status === 201) {
          openNotificationWithIcon(
            "error",
            "FALTAN VALORES PARA CREAR EL TOKEN"
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
            "NOMBRE DE TOKEN YA EXISTENTE"
          );
        }
        if (res.status === 205) {
          openNotificationWithIcon(
            "error",
            "LA LOCALIZACIÓN DEBE COINCIDIR CON LA COMPAÑIA"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "NO SE HA PODIDO CREAR EL TOKEN",
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
      <h2> CREAR TOKEN</h2>
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
                  label="NOMBRE DEL TOKEN:"
                  name="tokenName"
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
                    placeholder={"Definir nombre del Token"}
                    suffix={
                      <Clipboard onClick={handleClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Pegar el nombre del Token">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="tokenName"
                    value={token.tokenName}
                    onChange={handleChange}
                    style={{ width: '95%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="DESCRIPCIÓN DE CÓDIGO LISTA EUROPEA DE RESIDUOS:"
                  name="tokenLer"
                  rules={[
                    {
                      required: true,
                      message: "Introducir el código LER asociado al Token"
                    }
                  ]}
                >
                  <AutoComplete
                    size="large"
                    dataSource={lerData.LER.map(data => data.name)}
                    placeholder={"Descripción del código LER"}
                    style={{ width: '95%' }}
                    onSelect={(data) => data}
                    onChange={onChange}
                  >
                  </AutoComplete>
                </Form.Item>

                <Form.Item
                  label="CÓDIGO LISTA EUROPEA DE RESIDUOS (LER):"
                  name="tokenLer"
                  rules={[
                    {
                      required: true,
                      message: "Introducir el código LER asociado al Token"
                    }
                  ]}
                >

                  <Input
                    size="large"
                    placeholder={"Código LER"}
                    name="tokenLer"
                    value={codLer.codLer}
                    disabled={true}
                    style={{
                      width: '35%',
                      backgroundColor: "white",
                      color: "black"
                    }}
                  />
                </Form.Item>

                <Form.Item
                  label="TIEMPO DE VIDA DEL TOKEN (TTL):"
                  name="tokenTTL"
                  rules={[
                    {
                      required: true,
                      message: "Introducir tiempo de vida (TTL) del Token"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder={"Token TTL"}
                    type="text"
                    name="tokenTTL"
                    value={token.tokenTTL}
                    onChange={handleChange}
                    style={{ width: '35%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="PRECIO DEL TOKEN EN Wei:"
                  name="tokenPrice"
                  rules={[
                    {
                      required: true,
                      message: "Introducir el precio del Token"
                    }
                  ]}
                  hasFeedback
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
                    name="tokenPrice"
                    value={token.tokenPrice}
                    onChange={handleChange}
                    style={{ width: '35%' }}
                  />
                </Form.Item>
                <br />
                <br />
                <br />
                <Form.Item >
                  <Button
                    type="primary"
                    htmlType="submit"
                    size="large"
                    block
                    style={{ width: '95%' }}
                  >
                    CREAR TOKEN
                      </Button>
                </Form.Item>
              </Col>
              <Col lg={16} md={24}>
                <Form.Item
                  label="LOCALIZACIÓN DEL TOKEN:"
                  name="location"
                >
                  <Map
                    ref={mapRef}
                    center={[mapState.lat, mapState.lng]}
                    zoom={mapState.zoom}
                    style={{ height: 585 }}
                  >
                    <TileLayer
                      url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                      attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                    />
                    <Control position="topright" >
                      <AutoComplete
                        style={{ width: 200 }}
                        size="large"
                        dataSource={locData.companies.map(data => data.name)}
                        placeholder="Nombre empresa"
                        onSelect={(data) => {
                          const comp = locData.companies.find(company => company.name === data)
                          setToken(prevValue => ({ ...prevValue, location: comp.geometry.coordinates }))
                          setMapState({ ...mapState, lat: comp.geometry.coordinates[1], lng: comp.geometry.coordinates[0], zoom: 14 });
                          //console.log(markerRef.current[comp.id].leafletElement);
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
                          ref={el => markerRef.current[location.id] = el}
                          onClick={() => {
                            setToken(prevValue => ({ ...prevValue, location: location.geometry.coordinates }))
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

export default CreateToken;
