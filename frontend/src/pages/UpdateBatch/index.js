import React, { useState, useEffect } from "react";
import {
  AutoComplete,
  Button,
  Col,
  Form,
  Input,
  InputNumber,
  notification,
  Row,
  Spin,
  Tooltip,
} from "antd";
import { Icon as NewIco } from "antd";
import axios from "axios";
import "./../../App.css";
import "./../../react-leaflet.css";
import Clipboard from 'react-clipboard.js';
import { useGlobal } from "reactn";

//Leaflet
//----------------------------------------------------------------------------------------------
import MarkerClusterGroup from 'react-leaflet-markercluster';
import { Map, TileLayer, Marker, Popup } from 'react-leaflet'
import * as locData from "./../../data/companies.json";
import 'react-leaflet-markercluster/dist/styles.min.css';
import Control from 'react-leaflet-control';
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

const UpdBatch = () => {

  const initialFormState = {
    tokenid: "",
    //idParticipant: "",
    batchId: "",
    batchAttr_gen1: "",
    location: "",
    tokenPrice: "",
  };

  const INITIAL_STATE = {
    lat: 43.2612,
    lng: -1.779,
    zoom: 12
  }

  const [updbatch, setUpdbatch] = useState(initialFormState);
  const [mapState, setMapState] = useState(INITIAL_STATE); //Estado por defecto
  const [loading, setLoading] = useState(false);
  const [global] = useGlobal();

  const onChange = (value) => {
    setUpdbatch({ ...updbatch, batchId: value }) //console.log(addbatch);
    updbatch.batchId = value;
  };

  function handleChange(event) {
    const value = event.target.value;
    setUpdbatch({
      ...updbatch,
      [event.target.name]: value
    });
  }

  const handleClick = () => {
    const value = global.value;
    setUpdbatch({ ...updbatch, tokenid: value })
  }

  const onClick = () => {
    const value = global.value;
    setUpdbatch({ ...updbatch, tokenPrice: value })  //console.log(updbatch);
  }

  //Create a news products
  const handleSubmit = (e) => {
    //clearState();
    e.preventDefault();
    setLoading(true)
    //onChange();

    const jwtToken = localStorage.getItem("token");
    //Set POST request
    axios
      .post(`https://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/updBatch`, { updbatch }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "LOTE ACTUALIZADO SATISFACTORIAMENTE"
          );
        }
        if (res.status === 201) {
          openNotificationWithIcon(
            "error",
            "FALTAN VALORES PARA ACTUALIZAR EL LOTE"
          );
        }
        if (res.status === 202) {
          openNotificationWithIcon(
            "error",
            "IDENTIFICADOR DE LOTE NO EXISTENTE PARA ESTE TOKEN"
          );
        }
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "CONTENEDOR NO DESPLEGADO PARA ESTA ENTIDAD"
          );
        }
        if (res.status === 204) {
          openNotificationWithIcon(
            "error",
            "ERROR DE INTERACCIÓN CON LA BLOCKCHAIN, REVISAR LA INFORMACIÓN QUE SE ENVÍA"
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
          "NO SE HA PODIDO ACTUALIZAR EL LOTE",
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

  useEffect(() => {
    delete L.Icon.Default.prototype._getIconUrl;
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
      iconUrl: require("leaflet/dist/images/marker-icon.png"),
      shadowUrl: require("leaflet/dist/images/marker-shadow.png")
    });
  }, []);

  const formItemLayout = {};

  return (
    <section className="CommentsWrapper">
      <h2> ACTUALIZAR LOTE</h2>
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
                  label="NOMBRE DEL TOKEN:"
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
                    placeholder={"Nombre del Token creado"}
                    suffix={
                      <Clipboard onClick={handleClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Pegar el nombre del Token">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="tokenid"
                    value={updbatch.tokenid}
                    onChange={handleChange}
                    style={{ width: '95%' }}
                  />

                </Form.Item>

                <Form.Item
                  label="IDENTIFICADOR ÚNICO DE LOTE:"
                  name="batchId"
                  rules={[
                    {
                      required: true,
                      message: "Introducir identificador único del lote"
                    }
                  ]}
                >
                  <InputNumber
                    size="large"
                    placeholder={"1"}
                    min={1}
                    max={100000}
                    name="batchId"
                    value={updbatch.batchId}
                    onChange={onChange}
                    style={{ width: '30%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="PESO (Kg)"
                  name="batchAttr_gen1"
                  rules={[
                    {
                      required: true,
                      message: "Introducir correctamente la métrica"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder={"1000 Kg"}
                    type="text"
                    name="batchAttr_gen1"
                    value={updbatch.batchAttr_gen1}
                    onChange={handleChange}
                    style={{ width: '30%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="MODIFICAR EL PRECIO DEL TOKEN EN Wei:"
                  name="tokenPrice"
                  rules={[
                    {
                      required: true,
                      message: "Modificar el precio del Token"
                    }
                  ]}
                  hasFeedback
                >
                  <Input
                    size="large"
                    placeholder={"100 Wei"}
                    suffix={
                      <Clipboard onClick={onClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Precio del Token del Marketplace como referencia">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="tokenPrice"
                    value={updbatch.tokenPrice}
                    onChange={handleChange}
                    style={{ width: '30%' }}
                  />
                </Form.Item>
                <br />
                <br />
                <br />
                <br />
                <Form.Item>
                  <Button
                    size="large"
                    type="primary"
                    htmlType="submit"
                    block
                    style={{ width: '95%' }}
                  >
                    ACTUALIZAR LOTE
                       </Button>
                </Form.Item>

              </Col>

              <Col lg={16} md={24}>

                <Form.Item
                  label="LOCALIZACIÓN DEL TOKEN:"
                  name="location"
                >
                  <Map
                    center={[mapState.lat, mapState.lng]}
                    zoom={mapState.zoom}
                    style={{ height: 501 }}>
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
                          setUpdbatch(prevValue => ({ ...prevValue, location: comp.geometry.coordinates }))
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
                            setUpdbatch(prevValue => ({ ...prevValue, location: location.geometry.coordinates }))
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

export default UpdBatch;