import React, { useState } from "react";
import {
  AutoComplete,
  Row,
  Col,
  Form,
  Input,
  InputNumber,
  Button,
  Spin,
  notification,
  Table,
  Tooltip,
  Select
} from "antd";
import "./../../App.css";
import axios from "axios";
import { Icon as NewIco } from "antd";
import Clipboard from 'react-clipboard.js';
import { useGlobal } from "reactn";

const style = { background: '#ffffff', padding: '0px 0' };

const dataSource = ['Variable', 'Variation', 'Standard Clause', 'Custom Text'];

const _defaultCosts = [
  {
    feature: "",
    name: "",
    price: 0,
  }
];

const initialFormState = {
  raname: "",
  articleNo: "",
};

const AddArticle = () => {

  const formItemLayout = {};
  const [global] = useGlobal();
  const [loading, setLoading] = useState(false);
  const [costs, setCosts] = useState(_defaultCosts);
  let userDetails = JSON.parse(localStorage.getItem('user'));
  const [addArticle, setAddArticle] = useState(initialFormState);

  const handleCostsChange = event => {
    const _tempCosts = [...costs];
    _tempCosts[event.target.dataset.id][event.target.name] = event.target.value;
    setCosts(_tempCosts);
  };

  const addNewCost = () => {
    setCosts(prevCosts => [...prevCosts, { name: "", price: 0 }]);
  };

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

  const onChange = (value) => {
    setAddArticle({ ...addArticle, articleNo: value })
    addArticle.articleNo = value;
  };

  function handleChange(event) {
    const value = event.target.value;
    console.log(value);
    setAddArticle({
      ...addArticle,
      [event.target.name]: value
    });
  }


  const handleSubmit = (e) => {
    e.preventDefault();
    setLoading(true)
    const jwtToken = localStorage.getItem("token");
    //Set POST request
    axios
      .post(`http://${process.env.REACT_APP_GATEWAY_HOST}:${process.env.REACT_APP_GATEWAY_PORT}/proposeAddArticle`, { addArticle, costs, userDetails }, { headers: { "Authorization": `Bearer ${jwtToken}` } })
      .then((res) => {
        console.log(addArticle);
        console.log(costs);
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
              <Col lg={24} md={24}>

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
                    style={{ width: '35%' }}
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
                <Form.Item hasFeedback>
                  <Row type="flex" justify="start" gutter={16}>
                    <Col xs={4} sm={4} xl={4} >
                      <div style={style}> FEATURE </div>
                    </Col>
                    <Col xs={2} sm={2} xl={2}>
                      <div style={style}> ID </div>
                    </Col>
                    <Col xs={5} sm={5} xl={5}>
                      <div style={style}> VALUE </div>
                    </Col>
                  </Row>
                  {costs.map((item, index) => (
                    <Row type="flex" justify="start" gutter={16} key={index}>
                      <Col xs={4} sm={4} xl={4} >
                        <Input
                          name="feature"
                          size="large"
                          data-id={index}
                          style={{ width: '100%' }}
                          type="text"
                          value={item.feature}
                          onChange={handleCostsChange}
                        />
                      </Col>
                      <Col xs={2} sm={2} xl={2}>
                        <Input
                          name="price"
                          size="large"
                          placeholder={"ID"}
                          style={{ width: '100%' }}
                          data-id={index}
                          type="number"
                          value={item.price}
                          onChange={handleCostsChange}
                        />
                      </Col>
                      <Col xs={5} sm={5} xl={5}>
                        <Input
                          name="name"
                          size="large"
                          data-id={index}
                          style={{ width: '100%' }}
                          type="text"
                          value={item.name}
                          onChange={handleCostsChange}
                        />
                      </Col>
                    </Row>
                  ))}
                </Form.Item>
                <Form.Item>
                  <Button
                    type="primary"
                    icon="plus"
                    style={{ width: '2%' }}
                    onClick={addNewCost}>
                  </Button>
                  <br />
                  <br />
                  <Button
                    size="large"
                    type="primary"
                    htmlType="submit"
                    block
                    style={{ width: '35%' }}
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