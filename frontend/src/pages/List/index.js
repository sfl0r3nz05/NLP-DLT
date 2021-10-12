import React, { useState, useEffect } from "react";
import { useGlobal } from "reactn";
import moment from "moment";
import axios from "axios";
import { Icon, Row, Col, Table, Tag, Tooltip } from "antd";
import Search from "../../components/table/search";
import SearchDates from "../../components/table/searchDates";
//---------------------------------------------------------------------------------------------
import { CopyToClipboard } from 'react-copy-to-clipboard';
//---------------------------------------------------------------------------------------------

const RenderList = () => {

  const initialFormState = {
    token_name: "",
    token_sellable: "",
    token_price: "",
    token_state: "",
    token_ler: "",
    participant_name: "",
    location: "",
    timestamp: "",
  };
  const copyState = {
    value: '',
    copied: false,
  };

  const [list, setList] = useState([initialFormState]);
  const [copy, setCopy] = useState([copyState]);
  const [, setGlobal] = useGlobal();

  const onCopy = (value) => {
    copy.value = value;
    setCopy({ copied: true })
    setGlobal({ value: value })
    setGlobal({ copied: true })
  }

  useEffect(() => {
    axios
      .get(`https://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/list`, {
        headers: { 'Content-Type': 'application/json' }
      })
      .then(res => {
        //console.log("res", res);
        setList(res.data);
        if (res.ok) {
          return res.json();
        }
        throw res;
      })
      .then(resJson => {
        setList(resJson);
      })
      .catch(error => { });
  }, []); // Execut some element of the array changue

  const columns = [
    {
      title: "Nombre del Token", dataIndex: "token_name", key: "token_name", ...Search("token_name", "por nombre del Token"), align: 'center',
      render: (token_name) => (
        <Row>
          <Col span={20}>
            {token_name}
          </Col>
          <Col span={4}>
            <CopyToClipboard
              text={token_name}
              onCopy={onCopy}
            >
              <Tooltip title="Copiar el nombre del Token">
                <Icon type="copy" style={{ color: 'black', fontSize: 'large' }} />
              </Tooltip>
            </CopyToClipboard>
          </Col>
        </Row>
      )
    },
    {
      title: "Vendibilidad", dataIndex: "token_sellable", key: "token_sellable",
      filters: [
        { text: "Comprable", value: "comprable" },
        { text: "No vendible", value: "no vendible" },
      ],
      filterMultiple: false,
      onFilter: (value, record) => record.token_sellable.toLowerCase().includes(value),
      render(token_sellable, record) {
        return {
          children: <Tag color={(token_sellable) === "comprable" ? 'green' : 'volcano'}>{token_sellable}</Tag>
        };
      }, align: 'center'
    },
    {
      title: "Precio del Token (ETH)", dataIndex: "token_price", key: "token_price", ...Search("token_price", "por precio del Token"), align: 'center',
      render: (token_price) => (
        <Row>
          <Col span={20}>
            {token_price}
          </Col>
          <Col span={4}>
            <CopyToClipboard
              text={token_price}
              onCopy={onCopy}
            >
              <Tooltip title="Copiar el precio del Token">
                <Icon type="copy" style={{ color: 'black', fontSize: 'large' }} />
              </Tooltip>
            </CopyToClipboard>
          </Col>
        </Row>
      )
    },
    { title: "Estado actual del Token", dataIndex: "token_state", key: "token_state", ...Search("token_state", "por estado del Token"), align: 'center' },
    { title: "LER del Token", dataIndex: "token_ler", key: "token_ler", ...Search("token_ler", "por código LER del Token"), align: 'center' },
    { title: "Dueño del Token", dataIndex: "participant_name", key: "participant_name", ...Search("token_ler", "por dueño del Token"), align: 'center' },
    {
      title: "Localización del Token", dataIndex: "location", key: "location", align: 'center',
      render: location => <a target="popup" href={'https://www.google.com/maps/search/?api=1&query=' + location[1] + "," + location[0]}>
        <Tooltip title="Enlace a Google Maps">
          <Icon type="global" style={{ color: 'sky', fontSize: 'large' }} />
        </Tooltip></a>
    },
    { title: "Fecha de cambio de estado", dataIndex: "timestamp", key: "timestamp", sorter: (a, b) => moment(a.timestamp).unix() - moment(b.timestamp).unix(), defaultSortOrder: "descend", ...SearchDates("timestamp"), render: date => moment(date * 1000).format("DD/MM/YYYY"), align: 'center' },
  ]; //{'user/' + record.name}

  return (
    <section className="CommentsWrapper">
      <h2>TOKENS EN MARKETPLACE</h2>
      <Table bordered rowKey="token_name" columns={columns} dataSource={list} />
    </section>
  );
};

export default RenderList;
