import React, { useState } from "react";
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

const dataSource = ['Variable', 'Variation', 'Standard Clause', 'Custom Text'];

const AddArticle = () => {

  const formItemLayout = {};
  const { TextArea } = Input;
  const [global] = useGlobal();
  const [loading, setLoading] = useState(false);
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const [addArticle, setAddArticle] = useState(initialFormState);
  const [formVariables, setFormVariables] = useState([{ id: "", key: "", value: "" }])
  const [formVariations, setFormVariations] = useState([{ value: "" }])
  const [formStdClauses, setFormStdClauses] = useState([{ value: "" }])
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
    console.log(value);
    setAddArticle({
      ...addArticle,
      [event.target.name]: value
    });
  }

  const handleVariablesChange = (i, e) => {
    let newFormVariables = [...formVariables];
    newFormVariables[i][e.target.name] = e.target.value;
    setFormVariables(newFormVariables);
  }

  const handleVariationsChange = (i, e) => {
    let newFormVariations = [...formVariations];
    newFormVariations[i][e.target.name] = e.target.value;
    setFormVariations(newFormVariations);
  }

  const handleStdClausesChange = (i, e) => {
    let newFormStdClauses = [...formStdClauses];
    newFormStdClauses[i][e.target.name] = e.target.value;
    setFormStdClauses(newFormStdClauses);
  }

  const onChange = (value) => {
    setAddArticle({ ...addArticle, articleNo: value })
    addArticle.articleNo = value;
    outputNLP.NLP.map(item => {
      item.articles.map(data => {
        if (data.id == value) {
          setAddArticle({ ...addArticle, articleName: data.article })
          addArticle.articleName = data.article;
        }
      })
    })
  };

  let addVariablesFields = () => {
    setFormVariables([...formVariables, { id: "", key: "", value: "" }])
  }

  let removeVariablesFields = (i) => {
    let newFormVariables = [...formVariables];
    newFormVariables.splice(i, 1);
    setFormVariables(newFormVariables)
  }

  let addVariationsFields = () => {
    setFormVariations([...formVariations, { value: "" }])
  }

  let removeVariationsFields = (i) => {
    let newFormVariations = [...formVariations];
    newFormVariations.splice(i, 1);
    setFormVariations(newFormVariations)
  }

  let addStdClausesFields = () => {
    setFormStdClauses([...formStdClauses, { value: "" }])
  }

  let removeStdClausesFields = (i) => {
    let newFormStdClauses = [...formStdClauses];
    newFormStdClauses.splice(i, 1);
    setFormStdClauses(newFormStdClauses)
  }

  const handleInput = (e) => {
    input.value = !e;
    setInput(prevValue => ({ ...prevValue, value: !e }));
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)
    const jwtToken = localStorage.getItem("token");
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/proposeAddArticle`, { addArticle, formVariables, formVariations, formStdClauses, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
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
                    value={addArticle.raname || ""}
                    onChange={handleChange}
                    style={{ width: '40.5%' }}
                  />
                </Form.Item>

                <Form.Item hasFeedback>
                  {outputNLP.NLP.map((item, index) => (
                    <Row>
                      <Form.Item
                        label="SELECT ARTICLE ID AND NAME"
                      >
                        <Col span={2} >
                          <AutoComplete
                            size="large"
                            dataSource={item.articles.map(data => data.id)}
                            placeholder={"ID"}
                            onSelect={(data) => data}
                            onChange={onChange}
                            style={{ width: '75%' }}
                          />
                        </Col>
                        <Col span={14}>
                          <Input
                            size="large"
                            style={{ width: '55%' }}
                            placeholder={"Name of the Article"}
                            defaultValue={addArticle.articleName}
                            value={addArticle.articleName}
                          />
                        </Col>
                        <Col span={8} ></Col>
                      </Form.Item>

                      <Form.Item label="SELECT VARIABLES">
                        {formVariables.map((element, index) => (
                          <Row >
                            <Col span={2}>
                              <Input
                                name="id"
                                size="large"
                                placeholder={"ID"}
                                style={{ width: '82%' }}
                                type="number"
                                value={element.id}
                                onChange={e => handleVariablesChange(index, e)}
                              />
                            </Col>
                            <Col span={3} >
                              <Input
                                name="key"
                                size="large"
                                placeholder={"Key"}
                                style={{ width: '85%' }}
                                type="text"
                                value={element.key}
                                //dataSource={dataSource}
                                onChange={e => handleVariablesChange(index, e)}
                              >
                              </Input>

                            </Col>
                            <Col span={5}>
                              <Input
                                name="value"
                                size="large"
                                placeholder={"Value"}
                                style={{ width: '94%' }}
                                type="text"
                                value={element.value || ""}
                                onChange={e => handleVariablesChange(index, e)}
                              />
                            </Col>
                            <Col span={1}>
                              {
                                index ?
                                  <Button
                                    size="large"
                                    icon="minus"
                                    onClick={() => removeVariablesFields(index)}
                                    style={{ background: "gray", width: '80%' }}
                                  >
                                  </Button>
                                  : null
                              }
                            </Col>
                            <Col span={1}>
                              <Button
                                size="large"
                                type="primary"
                                icon="plus"
                                style={{ width: '80%' }}
                                onClick={addVariablesFields}
                              >
                              </Button>
                            </Col>
                            <Col span={8}></Col>
                          </Row>
                        ))}
                      </Form.Item>

                      <Form.Item label="SELECT VARIATIONS">
                        {formVariations.map((element, index) => (
                          <Row >
                            <Col span={20}>
                              <Input
                                name="value"
                                size="large"
                                placeholder={"Variation"}
                                style={{ width: '99%' }}
                                type="text"
                                value={element.value}
                                onChange={e => handleVariationsChange(index, e)}
                              />
                            </Col>
                            <Col span={1}>
                              {
                                index ?
                                  <Button
                                    size="large"
                                    icon="minus"
                                    onClick={() => removeVariationsFields(index)}
                                    style={{ background: "gray", width: '80%' }}
                                  >
                                  </Button>
                                  : null
                              }
                            </Col>
                            <Col span={1}>
                              <Button
                                size="large"
                                type="primary"
                                icon="plus"
                                style={{ width: '80%' }}
                                onClick={addVariationsFields}
                              >
                              </Button>
                            </Col>
                            <Col span={2}></Col>
                          </Row>
                        ))}
                      </Form.Item>

                      <Form.Item label="SELECT STANDARD CLAUSES">
                        {formStdClauses.map((element, index) => (
                          <Row >
                            <Col span={20}>
                              <Input
                                name="value"
                                size="large"
                                placeholder={"Standard Clause"}
                                style={{ width: '99%' }}
                                type="text"
                                value={element.value}
                                onChange={e => handleStdClausesChange(index, e)}
                              />
                            </Col>
                            <Col span={1}>
                              {
                                index ?
                                  <Button
                                    size="large"
                                    icon="minus"
                                    onClick={() => removeStdClausesFields(index)}
                                    style={{ background: "gray", width: '80%' }}
                                  >
                                  </Button>
                                  : null
                              }
                            </Col>
                            <Col span={1}>
                              <Button
                                size="large"
                                type="primary"
                                icon="plus"
                                style={{ width: '80%' }}
                                onClick={addStdClausesFields}
                              >
                              </Button>
                            </Col>
                            <Col span={2}></Col>
                          </Row>
                        ))}
                      </Form.Item>

                      <Form.Item label="ENABLE CUSTOM TEXTS">
                        <Row>
                          <Col span={20}>
                            <Switch
                              onChange={e => handleInput(e)}>
                            </Switch>
                            <TextArea
                              placeholder={"Name of the Roaming Agreement"}
                              name="customText"
                              size="large"
                              rows={4}
                              style={{ width: '99%' }}
                              disabled={input.value}
                              type="text"
                              value={addArticle.customText}
                              onChange={handleChange}
                            />
                          </Col>
                          <Col span={4}></Col>
                        </Row>
                      </Form.Item>

                    </Row>
                  ))}
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
              </Col>
            </Spin>
          </Form>
        </Col>
      </Row>
    </section >
  );
};

export default AddArticle;