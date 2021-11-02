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
  Switch,
  Select
} from "antd";
import "./../../App.css";
import axios from "axios";
import { useGlobal } from "reactn";
import outputNLP from "./../../data/outputNLP.json";

const AddArticle = () => {

  const initialFormState = {
    raname: "",
    articleNo: "",
    articleName: "",
    customText: "",
  };

  const formItemLayout = {};
  const { TextArea } = Input;
  const { Option } = Select;
  const [loading, setLoading] = useState(false);
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const [addArticle, setAddArticle] = useState(initialFormState);
  const [input, setInput] = useState({ value: true });
  const [itemmap, setItemmap] = useState([{ uid: "", articles: [{}] }]);
  //const [itemmap, setItemmap] = useState({ uid: "", type: "", verify: "", hint: "", articles: [{ id: "", article: "", uuid: "" }] });

  const openNotificationWithIcon = (type, title, description) => {
    notification[type]({
      message: title,
      description: description,
    });
  };

  function handleChange(e) {
    //console.log(e);
    addArticle.raname = e;
    setAddArticle(prevValue => ({ ...prevValue, raname: e.hint }));
    let new_e = outputNLP.find(item => {
      return (item.hint === e)
    })
    let new_e_e = outputNLP.map(item => {
      if (item.hint === e) {
        return item.articles.map(data => data.article)
      }
    })

    if (new_e_e[0] === undefined) {
      itemmap.articles = new_e_e[0];
      setItemmap(prevValue => ({ ...prevValue, articles: new_e_e[0] }));
      console.log("herex", itemmap);
    } else {
      itemmap.articles = new_e_e[1];
      setItemmap(prevValue => ({ ...prevValue, articles: new_e_e[0] }));
      console.log("herex2", itemmap);
    }
  }

  function handleChange2(event) {
    const value = event.target.value;
    setAddArticle({
      ...addArticle,
      [event.target.name]: value
    });
  }

  const onChange2 = (value) => {
    console.log(value);
    setAddArticle({ ...addArticle, articleName: value })
    addArticle.articleName = value;
    outputNLP.map(item => {
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
                  <Row>
                    <Form.Item label="NAME OF THE ROAMING AGREEMENT">
                      <Select
                        size="large"
                        name="raname"
                        onChange={e => handleChange(e)}
                        style={{ width: '40.5%' }}
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
                      <Col span={11}>
                        <Select
                          size="large"
                          placeholder={"Name of the Article"}
                          style={{ width: '89%' }}
                          onChange={onChange2}
                        >
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
                          onChange={handleChange2}
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