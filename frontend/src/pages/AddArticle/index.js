import React, { useState, useEffect } from "react";
import {
  Row,
  Col,
  Form,
  Input,
  InputNumber,
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

const AddArticle = () => {

  useEffect(() => {
    delete L.Icon.Default.prototype._getIconUrl;
    L.Icon.Default.mergeOptions({
      iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
      iconUrl: require("leaflet/dist/images/marker-icon.png"),
      shadowUrl: require("leaflet/dist/images/marker-shadow.png")
    });
  }, []);

  const { TextArea } = Input;
  const formItemLayout = {};
  const [global] = useGlobal();
  const initialFormState = {
    raname: "",
    articleNo: "",
    variables: "",
    variations: "",
    customTexts: "",
    standardClauses: "",
  };
  const [loading, setLoading] = useState(false);
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const [addArticle, setAddArticle] = useState(initialFormState);

  const onClick = () => {
    const value = global.value;
    addArticle.raname = value;
    setAddArticle(prevValue => ({ ...prevValue, tokenid: value }));
  }

  const onChange = (value) => {
    setAddArticle({ ...addArticle, articleNo: value })
    addArticle.articleNo = value;
  };

  function handleChange(event) {
    const value = event.target.value;
    setAddArticle({
      ...addArticle,
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
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/proposeAddArticle`, { addArticle, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
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
      <h2> PROPOSE AN ARTICLE FOR A ROAMING AGREEMENT</h2>
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
              <Col lg={16} md={24}>

                <Form.Item
                  label="NAME OF THE ROAMING AGREEMENT"
                  name="raname"
                  rules={[
                    {
                      required: true,
                      message: "NAME OF THE ROAMING AGREEMENT"
                    }
                  ]}
                >
                  <Input
                    size="large"
                    placeholder={"Paste the Name of the Roaming Agreement"}
                    suffix={
                      <Clipboard onClick={onClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                        <Tooltip title="Paste raname Name">
                          <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                        </Tooltip>
                      </Clipboard>
                    }
                    type="text"
                    name="raname"
                    value={addArticle.raname}
                    onChange={handleChange}
                    style={{ width: '50%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="ARTICLE NUMBER:"
                  name="articleNo"
                  rules={[
                    {
                      required: true,
                      message: "Introduce the article number"
                    }
                  ]}
                >
                  <InputNumber
                    size="large"
                    placeholder={"1"}
                    min={1}
                    max={100000}
                    name="articleNo"
                    value={addArticle.articleNo}
                    onChange={onChange}
                    style={{ width: '10%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="LIST OF VARIABLES"
                  name="variables"
                  rules={[
                    {
                      required: true,
                      message: "LIST OF VARIABLES"
                    }
                  ]}
                >
                  <TextArea
                    size="large"
                    rows={2}
                    placeholder={"Include a list of variables"}
                    type="text"
                    name="variables"
                    value={addArticle.variables}
                    onChange={handleChange}
                    style={{ width: '100%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="LIST OF VARIATIONS"
                  name="variations"
                  rules={[
                    {
                      required: true,
                      message: "LIST OF VARIATIONS"
                    }
                  ]}
                >
                  <TextArea
                    size="large"
                    rows={2}
                    placeholder={"Include a list of variations"}
                    type="text"
                    name="variations"
                    value={addArticle.variations}
                    onChange={handleChange}
                    style={{ width: '100%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="LIST OF CUSTOM TEXTS"
                  name="customTexts"
                  rules={[
                    {
                      required: true,
                      message: "LIST OF CUSTOM TEXTS"
                    }
                  ]}
                >
                  <TextArea
                    size="large"
                    rows={2}
                    placeholder={"Include a list of custom texts"}
                    type="text"
                    name="customTexts"
                    value={addArticle.customTexts}
                    onChange={handleChange}
                    style={{ width: '100%' }}
                  />
                </Form.Item>

                <Form.Item
                  label="LIST OF STANDARD CLAUSES"
                  name="standardClauses"
                  rules={[
                    {
                      required: true,
                      message: "LIST OF STANDARD CLAUSES"
                    }
                  ]}
                >
                  <TextArea
                    size="large"
                    rows={2}
                    placeholder={"Include a list of standard clauses"}
                    type="text"
                    name="standardClauses"
                    value={addArticle.standardClauses}
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
                    style={{ width: '50%' }}
                  >
                    PROPOSE ARTICLE
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

export default AddArticle;