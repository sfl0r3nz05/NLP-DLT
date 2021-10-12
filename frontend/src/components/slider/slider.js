import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { Slider as SliderAntd, InputNumber, Row, Col, Form } from "antd";

const SliderWrapper = styled(SliderAntd)`
  display: flex;
  flex-direction: column;
  width: 100%;
`;

const Slider = ({ min, max, callback, data, onChange }) => {
  const [value, setvalue] = useState(null);

  useEffect(() => {
    if (!value) setvalue(data);
  }, [data, value]);

  const handleChangue = val => {
    setvalue(val);
    onChange(val);
  };

  const handleAfterChange = val => callback(val);
  return (
    <Row>
      <Form.Item label="Nº etiquetas">
        <Col span={12}>
          <SliderAntd
            min={min}
            max={max}
            style={{ margin: 0 }}
            tooltipVisible={false}
            onChange={handleChangue}
            onAfterChange={handleAfterChange}
            value={typeof value === "number" ? value : 0}
          />
        </Col>
        <Col span={4}>
          <InputNumber
            min={1}
            max={20}
            style={{ marginLeft: 16 }}
            value={typeof value === "number" ? value : 0}
            onChange={handleChangue}
            onAfterChange={handleAfterChange}
            step={0.01}
          />
        </Col>
      </Form.Item>
    </Row>
  );
};

export default Slider;
