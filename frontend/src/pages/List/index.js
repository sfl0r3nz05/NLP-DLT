import React, { useState, useEffect } from "react";
import ReactCountryFlag from "react-country-flag"
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

  let userDetails = JSON.parse(localStorage.getItem('user'));

  const initialFormState = {
    mno1: "",
    country_mno1: "",
    mno2: "",
    country_mno2: "",
    ra_name: "",
    ra_status: "",
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
      .get(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/list`, {
        params: { ID: userDetails },
        headers: { 'Content-Type': 'application/json' }
      })
      .then(res => {
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
      title: "MNO1 Name", dataIndex: "mno1", key: "mno1", ...Search("mno1", "MNO Name"), align: 'center',
      render: (mno1) => (
        <Row>
          <Col span={20}>
            {mno1}
          </Col>
          <Col span={4}>
            <CopyToClipboard
              text={mno1}
              onCopy={onCopy}
            >
              <Tooltip title="Copy the name of MNO1">
                <Icon type="copy" style={{ color: 'black', fontSize: 'large' }} />
              </Tooltip>
            </CopyToClipboard>
          </Col>
        </Row>
      )
    },
    {
      title: "MNO Country", dataIndex: "country_mno1", key: "country_mno1", ...Search("country_mno1", "por precio del Token"), align: 'center',
      render: (country_mno1) => (
        <Row>
          <Col span={20}>
            <ReactCountryFlag countryCode={country_mno1} svg style={{
              width: '2em',
              height: '2em',
            }} />
          </Col>
        </Row>
      )
    },
    {
      title: "MNO2 Name", dataIndex: "mno2", key: "mno2", ...Search("mno2", "MNO Name"), align: 'center',
      render: (mno2) => (
        <Row>
          <Col span={20}>
            {mno2}
          </Col>
          <Col span={4}>
            <CopyToClipboard
              text={mno2}
              onCopy={onCopy}
            >
              <Tooltip title="Copy the name of MNO2">
                <Icon type="copy" style={{ color: 'black', fontSize: 'large' }} />
              </Tooltip>
            </CopyToClipboard>
          </Col>
        </Row>
      )
    },
    {
      title: "MNO Country", dataIndex: "country_mno2", key: "country_mno2", ...Search("country_mno2", "por precio del Token"), align: 'center',
      render: (country_mno2) => (
        <Row>
          <Col span={20}>
            <ReactCountryFlag countryCode={country_mno2} svg style={{
              width: '2em',
              height: '2em',
            }} />
          </Col>
        </Row>
      )
    },
    {
      title: "RA name", dataIndex: "ra_name", key: "ra_name", ...Search("ra_name", "por precio del Token"), align: 'center',
      render: (ra_name) => (
        <Row>
          <Col span={20}>
            {ra_name}
          </Col>
          <Col span={4}>
            <CopyToClipboard
              text={ra_name}
              onCopy={onCopy}
            >
              <Tooltip title="Copy name of Roaming Agreement">
                <Icon type="copy" style={{ color: 'black', fontSize: 'large' }} />
              </Tooltip>
            </CopyToClipboard>
          </Col>
        </Row>
      )
    },
    {
      title: "RA Status", dataIndex: "ra_status", key: "ra_status",
      filters: [
        { text: "Comprable", value: "comprable" },
        { text: "No vendible", value: "no vendible" },
      ],
      filterMultiple: false,
      onFilter: (value, record) => record.ra_status.toLowerCase().includes(value),
      render(ra_status, record) {
        return {
          children: <Tag color={(ra_status) === "INIT" ? 'green' : 'volcano'}>{ra_status}</Tag>
        };
      }, align: 'center'
    },
    { title: "Date of status change", dataIndex: "timestamp", key: "timestamp", sorter: (a, b) => moment(a.timestamp).unix() - moment(b.timestamp).unix(), defaultSortOrder: "descend", ...SearchDates("timestamp"), render: date => moment(date * 1000).format("DD/MM/YYYY"), align: 'center' },
  ];

  return (
    <section className="CommentsWrapper">
      <h2>MOBILE NETWORK OPERATORS IN ROAMING AGREEMENTS</h2>
      <Table bordered rowKey="mno1" columns={columns} dataSource={list} />
    </section>
  );
};

export default RenderList;
