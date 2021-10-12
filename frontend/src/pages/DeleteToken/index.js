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
  Tooltip,
} from "antd";
import axios from "axios";
import { Icon as NewIco } from "antd";
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

const INITIAL_STATE = {
  lat: 43.2612,
  lng: -1.779,
  zoom: 12
}

notification.config({
  placement: "topRight",
  bottom: 50,
  duration: 1.5,
});

const DelToken = () => {

  const initialFormState = {
    tokenid: "",
    //idParticipant: "",
    location: "",
  };

  const [deltoken, setDeltoken] = useState(initialFormState);
  const [mapState, setMapState] = useState(INITIAL_STATE); //Estado por defecto
  const [loading, setLoading] = useState(false);
  const [global] = useGlobal();

  const onChange = (values) => {
    //console.log("Success:", values);
  };

  const onClick = () => {
    const value = global.value;
    deltoken.tokenid = value;
    setDeltoken(prevValue => ({ ...prevValue, tokenid: value }));
    //console.log(deltoken);
  }

  function handleChange(event) {
    const value = event.target.value;
    setDeltoken({
      ...deltoken,
      [event.target.name]: value
    });
  }

  //Create a news products
  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)
    onChange();

    const jwtToken = localStorage.getItem("token");
    axios
      .post(`https://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/delToken`, { deltoken }, { headers: { "Authorization": `Bearer ${jwtToken}` } })   //Set POST request
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "TOKEN ELIMINADO SATISFACTORIAMENTE"
          );
        }
        if (res.status === 201) {
          openNotificationWithIcon(
            "error",
            "FALTAN VALORES PARA ELIMINAR EL TOKEN"
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
            "LA LOCALIZACIÓN DEBE COINCIDIR CON LA COMPAÑIA"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "NO SE HA PODIDO ELIMINAR EL TOKEN",
        )
      )
      .finally(() => setLoading(false));
  };

  const formItemLayout = {};

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

  return (
    <section className="CommentsWrapper">
      <h2> ELIMINAR TOKEN</h2>
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
                    placeholder={"Nombre del Token a eliminar"}
                    id="textId"
                    suffix={
                      <Clipboard onClick={onClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Pegar el nombre del Token">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="tokenid"
                    value={deltoken.tokenid}
                    onChange={handleChange}
                    style={{ width: '95%' }}
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
                    style={{ width: '95%' }}
                  >
                    ELIMINAR TOKEN
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
                    style={{ height: 500 }}
                  >
                    <TileLayer
                      noWrap={true}
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
                          setDeltoken(prevValue => ({ ...prevValue, location: comp.geometry.coordinates }))
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
                            setDeltoken(prevValue => ({ ...prevValue, location: location.geometry.coordinates }))
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

export default DelToken;