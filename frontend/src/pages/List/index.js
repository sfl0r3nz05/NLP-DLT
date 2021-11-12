import React, { useState, useEffect } from "react";
import ReactCountryFlag from "react-country-flag"
import { useGlobal } from "reactn";
import moment from "moment";
import axios from "axios";
import { Input, Button, Col, Form, Icon, notification, Modal, Row, Table, Tag, Tooltip } from "antd";
import Search from "../../components/table/search";
import SearchDates from "../../components/table/searchDates";
import { CopyToClipboard } from 'react-copy-to-clipboard';
import TextArea from "antd/lib/input/TextArea";

notification.config({
  placement: "topRight",
  bottom: 50,
  duration: 1.5,
});

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
    articles: [{
      articleId: "",
      articleName: "",
      articleStatus: "",
      variables: [],
      variations: [],
      stdclauses: [],
      customtexts: []
    }]
  };
  const [list, setList] = useState([initialFormState]);

  const copyState = {
    value: '',
    copied: false,
  };
  const [copy, setCopy] = useState([copyState]);
  const [, setGlobal] = useGlobal();
  const onCopy = (value) => {
    copy.value = value;
    setCopy({ copied: true })
    setGlobal({ value: value })
    setGlobal({ copied: true })
  }

  const openNotificationWithIcon = (type, title, description) => {
    notification[type]({
      message: title,
      description: description,
    });
  };

  const [selectedRow, setSelectedRow] = useState({ variables: [], variations: [], stdclauses: [], customtexts: [] });
  const [isModalVisible, setIsModalVisible] = useState(false);
  const showModal = (v) => {
    setSelectedRow(v)
    setIsModalVisible(true);
  };
  const handleCancel = () => {
    setIsModalVisible(false);
  };

  useEffect(() => {
    axios
      .get(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/list`, {
        params: { ID: userDetails },
        headers: { 'Content-Type': 'application/json' }
      })
      .then(res => {
        setList(res.data);
        console.log(res.data);
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

  const handleSubmit = (e, mno1) => {
    e.preventDefault();
    const jwtToken = localStorage.getItem("token");
    //Set POST request
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/acceptAgreementInitiation`, { list, mno1, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
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
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "THIS MNO CANNOT ACCEPT THE CHANGES PROPOSED BY ITSELF"
          );
        }
        if (res.status === 204) {
          openNotificationWithIcon(
            "success",
            "SENT PROPOSAL TO REACH AGREEMENT"
          );
        }
        if (res.status === 205) {
          openNotificationWithIcon(
            "success",
            "CONGRATULATIONS: AGREEMENT REACHED !!!"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "UNREGISTERED ROAMING AGREEMENT",
        )
      )
  };

  const handleOk = (e) => {
    e.preventDefault();
    const jwtToken = localStorage.getItem("token");
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/acceptProposedChanges`, { list, selectedRow, formVariables, formCustomText, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "SUCCESSFULLY ACCEPTED ARTICLE"
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
            "THIS MNO CANNOT ACCEPT THE CHANGES PROPOSED BY ITSELF"
          );
        }
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "NOT ALLOWED TO MODIFY VARIABLES, VARIATIONS OR CUSTOM TEXTS WHEN CHANGES ARE ACCEPTED"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "UNREGISTERED ROAMING AGREEMENT",
        )
      )
  };

  const handleChange = (e) => {
    e.preventDefault();
    const jwtToken = localStorage.getItem("token");
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/proposeUpdateArticle`, { list, selectedRow, formVariables, formCustomText, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "SUCCESSFULLY UPDATED ARTICLE"
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
            "THIS MNO CANNOT ACCEPT THE CHANGES PROPOSED BY ITSELF"
          );
        }
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "MUST MODIFY VARIABLES, VARIATIONS OR CUSTOM TEXTS WHEN A CHANGE IN THE ARTICLE IS PROPOSED"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "UNREGISTERED ROAMING AGREEMENT",
        )
      )
  };

  const handleReject = (e) => {
    e.preventDefault();
    const jwtToken = localStorage.getItem("token");
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/rejectProposedChanges`, { list, selectedRow, formVariables, formCustomText, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
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
            "THIS MNO CANNOT ACCEPT THE CHANGES PROPOSED BY ITSELF"
          );
        }
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "MUST MODIFY VARIABLES, VARIATIONS OR CUSTOM TEXTS WHEN A CHANGE IN THE ARTICLE IS PROPOSED"
          );
        }
      })
      .catch(() =>
        openNotificationWithIcon(
          "error",
          "UNREGISTERED ROAMING AGREEMENT",
        )
      )
  };

  const [formVariables, setFormVariables] = useState([{ key: "", value: "" }])
  const handleVariablesChange = (i, e, v) => {
    let newFormVariables = [...formVariables];
    newFormVariables[i] = { key: v, value: e.target.value }
    setFormVariables(newFormVariables);
  }

  const [formCustomText, setFormCustomText] = useState([{ value: "" }])
  function handleCustomText(e) {
    formCustomText[0] = { value: e.target.value }
    setFormCustomText(formCustomText)
  }

  const clicRow = (e) => {
    console.log(e);
    setSelectedRow(e)
  }

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
      title: "RA Status", dataIndex: "ra_status", key: "ra_status", align: 'center',
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
    {
      title: "Accept RA", dataIndex: "mno1", align: 'center', render: (mno1) => (
        <span>
          <a
            onClick={e => handleSubmit(e, mno1)}
          >
            <Icon type={list.map((rank, i, arr) => {
              if (arr.length - 1 === i) {
                return (rank.ra_status).toString()
              }
            }) == 'started_ra' ? 'unlock' : 'lock'} />
          </a>
        </span>
      )
    },
  ];

  return (
    <section className="CommentsWrapper">
      <h2>MOBILE NETWORK OPERATORS IN ROAMING AGREEMENTS</h2>
      <Table columns={columns} dataSource={list} expandedRowRender={(record) => {
        const columns = [
          {
            title: 'Article Number', dataIndex: 'articleId', key: 'articleId', align: 'center',
            render: (articleId) => (
              <Row>
                <Col>
                  {articleId}
                </Col>
              </Row>
            )
          },
          {
            title: 'Article Name', dataIndex: 'articleName', key: 'articleName', align: 'center', render: (articleName) => (
              <Row>
                <Col>
                  {articleName}
                </Col>
              </Row>
            )
          },
          {
            title: 'Article Status', dataIndex: 'articleStatus', key: 'articleStatus', align: 'center', render(articleStatus) {
              return {
                children: <Tag color={(articleStatus) === "INIT" ? 'green' : 'volcano'}>{articleStatus}</Tag>
              };
            }
          },
          {
            title: 'Article in details', align: 'center', render() {
              return {
                children: <>
                  <Button type="primary" onClick={() => showModal(selectedRow)}>
                    View
                  </Button>
                </>
              };
            }
          },
        ];
        return (
          <Table
            bordered
            columns={columns}
            dataSource={record.articles}
            pagination={false}
            onRowClick={clicRow}
          />
        );
      }} />
      {<Modal title="Article in detail" visible={isModalVisible} width={800} onCancel={handleCancel} footer={[
        <Button key="accept" type="primary" size="large" onClick={handleOk}>
          Accept
        </Button>,
        <Button key="changes" type="normal" size="large" onClick={handleChange}>
          Changes
        </Button>,
        <Button key="reject" type="dashed" size="large" onClick={handleReject}>
          Reject
        </Button>,
      ]}>
        <Form.Item label="Variables">
          {selectedRow && selectedRow.variables.map((data, index) =>
            < Row >
              <Col span={6}>
                <Input
                  size="large"
                  name="key"
                  style={{ width: '100%' }}
                  defaultValue={data.key}
                  disabled
                />
              </Col>
              <Col span={7}>
                <Input
                  name="value"
                  size="large"
                  placeholder={"value"}
                  style={{ width: '100%' }}
                  defaultValue={data.value}
                  onChange={e => handleVariablesChange(index, e, data.key)}
                />
              </Col>
              <Col span={11} />
            </Row>
          )}
        </Form.Item>
        <Form.Item label="Standard Clauses">

          < Row >
            <TextArea
              size="large"
              name="value"
              style={{ width: '100%' }}
              defaultValue={selectedRow && selectedRow.stdclauses.map(data => data.value)}
              rows={12}
              disabled
            //onChange={e => handleStdClauses(e)}
            />
          </Row>

        </Form.Item>
        <Form.Item label="Variations">
          {selectedRow && selectedRow.variations.map(data =>
            < Row >
              <TextArea
                size="large"
                name="value"
                style={{ width: '100%' }}
                defaultValue={data.value}
                rows={3}
              //onChange={e => handleVariations(e)}
              />
            </Row>
          )}
        </Form.Item>
        <Form.Item label="Custom Texts" >
          {selectedRow && selectedRow.customtexts.map(data =>
            < Row >
              <TextArea
                size="large"
                name="value"
                style={{ width: '100%' }}
                defaultValue={data.value}
                rows={6}
                onChange={e => handleCustomText(e)}
              />
            </Row>
          )}
        </Form.Item>
      </Modal>}
    </section >
  );
};

export default RenderList;
