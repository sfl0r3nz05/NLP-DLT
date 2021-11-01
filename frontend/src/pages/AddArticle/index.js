import React, { useState } from "react";
import { v4 as uuidv4 } from 'uuid'
import {
  AutoComplete,
  Row,
  Col,
  Form,
  Input,
  Button,
  Spin,
  notification,
  Tooltip,
  Switch
} from "antd";
import "./../../App.css";
import axios from "axios";
import { Icon as NewIco } from "antd";
import Clipboard from 'react-clipboard.js';
import { useGlobal } from "reactn";
import * as outputNLP from "./../../data/outputNLP.json";

const initialFormState = {
  raname: "",
  articleNo: "",
  articleName: "",
  customText: "",
};

const AddArticle = ({ current, path, onChange }) => {

  const formItemLayout = {};
  const { TextArea } = Input;
  const [global] = useGlobal();
  const [loading, setLoading] = useState(false);
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const [addArticle, setAddArticle] = useState(initialFormState);
  const [input, setInput] = useState({ value: true });


  const openNotificationWithIcon = (type, title, description) => {
    notification[type]({
      message: title,
      description: description,
    });
  };

  const onClick = () => {
    const value = global.value;
    addArticle.raname = value;
    setAddArticle(prevValue => ({ ...prevValue, raname: value }));
  }

  function handleChange(event) {
    const value = event.target.value;
    setAddArticle({
      ...addArticle,
      [event.target.name]: value
    });
  }

  const onChange2 = (value) => {
    setAddArticle({ ...addArticle, articleName: value })
    addArticle.articleName = value;
    outputNLP.NLP.map(item => {
      item.articles.map(data => {
        if (data.article == value) {
          setAddArticle({ ...addArticle, articleNo: data.id })
          addArticle.articleNo = data.id;
        }
      })
    })
  };

  const handleInput = (e) => {
    input.value = !e;
    setInput(prevValue => ({ ...prevValue, value: !e }));
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)
    const jwtToken = localStorage.getItem("token");
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/proposeAddArticle`, { addArticle, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        console.log(addArticle);
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

  return (
    <section className="CommentsWrapper">
      <h2> PROPOSE ARTICLES FOR ROAMING AGREEMENT</h2>
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
              <Col lg={1} md={24}></Col>
              <Col lg={23} md={24}>
                <Form.Item hasFeedback>
                  <Form.Item label="NAME OF THE ROAMING AGREEMENT">
                    <Input
                      size="large"
                      placeholder={"E.g.: RA001"}
                      suffix={
                        <Clipboard onClick={onClick} style={{ background: 'white', border: '0px', outline: '0px' }}>
                          <Tooltip title="Paste raname Name">
                            <NewIco type="snippets" style={{ color: 'black', fontSize: 'x-large' }} />
                          </Tooltip>
                        </Clipboard>
                      }
                      type="text"
                      name="raname"
                      onChange={handleChange}
                      style={{ width: '40.5%' }}
                    />
                  </Form.Item>
                  {outputNLP.NLP.map((item, index, arr) => (
                    <Row>
                      <Form.Item label="SELECT ARTICLE NAME AND ID">
                        <Col span={11}>
                          <AutoComplete
                            size="large"
                            dataSource={item.articles.map(data => data.article)}
                            placeholder={"Name of the Article"}
                            style={{ width: '89%' }}
                            onChange={onChange2}
                          />
                        </Col>
                        <Col span={2} >
                          <Input
                            size="large"
                            style={{ width: '100%' }}
                            placeholder={"ID"}
                            value={addArticle.articleNo}
                          />
                        </Col>
                        <Col span={11} ></Col>
                      </Form.Item>
                      <Form.Item label="ENABLE CUSTOM TEXTS">
                        <Col span={20}>
                          <Switch
                            size="large"
                            onChange={e => handleInput(e)}
                          />
                          <TextArea
                            placeholder={"Name of the Roaming Agreement"}
                            name="customText"
                            size="large"
                            rows={4}
                            style={{ width: '99%' }}
                            disabled={input.value}
                            type="text"
                            onChange={handleChange}
                          />
                        </Col>
                        <Col span={4}></Col>
                      </Form.Item>
                      <Form.Item>
                        <br />
                        <Button
                          size="large"
                          type="primary"
                          htmlType="submit"
                          block
                          style={{ width: '40.5%' }}
                        >
                          PROPOSE ARTICLE
                        </Button>
                      </Form.Item>
                    </Row>
                  ))}
                </Form.Item>
              </Col>
            </Spin>
          </Form>
        </Col>
      </Row>
    </section >
  );
};

export default AddArticle;