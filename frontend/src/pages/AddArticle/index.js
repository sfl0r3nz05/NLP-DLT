import React, { useState } from "react";
import {
  Row,
  Col,
  Form,
  Input,
  Button,
  Spin,
  notification,
  Select,
  Switch
} from "antd";
import "./../../App.css";
import axios from "axios";
import outputNLP from "./../../data/outputNLP.json";

const AddArticle = () => {

  const formItemLayout = {};
  const { TextArea } = Input;
  const { Option } = Select;

  const initialFormState = {
    raname: "",
    articleNo: "",
    articleName: "",
  };
  const [addArticle, setAddArticle] = useState(initialFormState);

  const [selectedArticles, setselectedArticles] = useState({ articles: [] });
  function handleChange(e) {
    addArticle.raname = e;
    setAddArticle(prevValue => ({ ...prevValue, raname: e }));
    let new_e = outputNLP.find(item => {
      return item.hint === e;
    })
    setselectedArticles(new_e)
  }

  const [formVariables, setFormVariables] = useState([{ key: "", value: "" }])
  const handleVariablesChange = (i, e, v) => {
    let newFormVariables = [...formVariables];
    newFormVariables[i] = { key: v, value: e.target.value }
    setFormVariables(newFormVariables);
  }

  const handleVariationsChange = (i, e, v) => {
    console.log(e.target.value);
  }

  const [selectedArticleVar, setSelectedArticleVar] = useState({ variables: [] });
  const [selectedArticleSub, setSelectedArticleSub] = useState({ subarticles: [] });
  const [selectedArticlesStdClause, setSelectedArticlesStdClause] = useState({ subarticles: [] });
  const [selectedArticlesVariation, setSelectedArticlesVariation] = useState({ subarticles: [] });
  const onChange = (value) => {
    setAddArticle({ ...addArticle, articleName: value })
    addArticle.articleName = value;
    outputNLP.map(item => {
      item.articles.map(data => {
        if (data.article === value) {
          setAddArticle({ ...addArticle, articleNo: data.id })
          addArticle.articleNo = data.id;
          setSelectedArticleVar(data)
          setSelectedArticleSub(data)
          if ((selectedArticleSub.subarticles.map(data => data.type)).find(datax => datax === 'stdClause')) {
            console.log("here");
            setSelectedArticlesStdClause(data)
          } else if ((selectedArticleSub.subarticles.map(data => data.type)).find(datax => datax === 'variation')) {
            console.log("here2");
            setSelectedArticlesVariation(data)
          }
        }
      })
    })
  };

  const [formCustomText, setFormCustomText] = useState([{ value: "" }])
  function handleChangeRA(e) {
    const value = e.target.value;
    setAddArticle({
      ...addArticle,
      [e.target.name]: value
    });

    formCustomText[0] = { value: e.target.value }
    setFormCustomText(formCustomText)
  }

  const [input, setInput] = useState({ value: true });
  const handleInput = (e) => {
    input.value = !e;
    setInput(prevValue => ({ ...prevValue, value: !e }));
  }

  const openNotificationWithIcon = (type, title, description) => {
    notification[type]({
      message: title,
      description: description,
    });
  };

  const [loading, setLoading] = useState(false);
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)
    const jwtToken = localStorage.getItem("token");
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/proposeAddArticle`, { addArticle, selectedArticlesStdClause, formVariables, selectedArticlesVariation, formCustomText, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        if (res.status === 200) {
          openNotificationWithIcon(
            "success",
            "ARTICLE SUCCESSFULLY REGISTERED"
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
            "THIS MNO DOES NOT BELONG TO THIS ROAMING AGREEMENT"
          );
        }
        if (res.status === 203) {
          openNotificationWithIcon(
            "error",
            "THIS ARTICLE HAS ALREADY BEEN ADDED"
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
                  <Row>
                    <Form.Item label="NAME OF THE ROAMING AGREEMENT">
                      <Select
                        size="large"
                        placeholder={"E.g.: RA001"}
                        name="raname"
                        onChange={e => handleChange(e)}
                        style={{ width: '35%' }}
                      >
                        {outputNLP.map((item) => (
                          <Option
                            key={item.uid}
                            value={item.hint}
                          >
                            {item.hint}
                          </Option>
                        ))}
                      </Select>
                    </Form.Item>

                    <Form.Item label="SELECT ARTICLE NAME AND ID">
                      <Row >
                        <Col span={9}>
                          <Select
                            size="large"
                            placeholder={"Title of the Article"}
                            style={{ width: '93%' }}
                            onChange={onChange}
                          >
                            {selectedArticles.articles.map((item) => (
                              <Option
                                key={item.id}
                                value={item.article}
                              >
                                {item.article}
                              </Option>
                            ))}
                          </Select>
                        </Col>
                        <Col span={2} >
                          <Input
                            size="large"
                            style={{ width: '100%' }}
                            placeholder={"ID"}
                            value={addArticle.articleNo}
                          />
                        </Col>
                        <Col span={13} ></Col>
                      </Row>
                    </Form.Item>

                    <Form.Item label="STANDARD CLAUSES DEFINED">
                      {selectedArticlesStdClause.subarticles.map((item, index) => (
                        <Row >
                          <Col span={24}>
                            <TextArea
                              size="large"
                              name="key"
                              rows={4}
                              style={{ width: '100%' }}
                              placeholder={"Key"}
                              defaultValue={item.content}
                              disabled
                            />
                          </Col>
                        </Row>
                      ))}
                    </Form.Item>

                    <Form.Item label="SELECT VARIABLES">
                      {selectedArticleVar.variables.map((item, index) => (
                        <Row >
                          <Col span={4}>
                            <Input
                              size="large"
                              name="key"
                              style={{ width: '85%' }}
                              placeholder={"Key"}
                              defaultValue={item.verify}
                              disabled
                            />
                          </Col>
                          <Col span={5}>
                            <Input
                              name="value"
                              size="large"
                              placeholder={"Value"}
                              style={{ width: '88%' }}
                              type="text"
                              onChange={e => handleVariablesChange(index, e, item.verify)}
                            />
                          </Col>
                          <Col span={15}></Col>
                        </Row>
                      ))}
                    </Form.Item>

                    <Form.Item label="SELECT VARIATIONS">
                      {selectedArticlesVariation.subarticles.map((item, index) => (
                        <Row >
                          <Col span={24}>
                            <Input
                              size="large"
                              type='checkbox'
                              name="key"
                              style={{ width: '100%' }}
                              value={item.content}
                              checked
                              onChange={e => handleVariationsChange(index, e)}
                            />
                          </Col>
                        </Row>
                      ))}
                    </Form.Item>

                    <Form.Item label="ENABLE CUSTOM TEXTS">
                      <Col span={24}>
                        <Switch
                          size="large"
                          onChange={e => handleInput(e)}
                        />
                        <TextArea
                          placeholder={"Name of the Roaming Agreement"}
                          name="customText"
                          size="large"
                          rows={4}
                          style={{ width: '100%' }}
                          disabled={input.value}
                          type="text"
                          onChange={handleChangeRA}
                        />
                      </Col>
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