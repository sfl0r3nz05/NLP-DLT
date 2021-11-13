import React from "react";
import { Slider as SliderAntd, InputNumber, Row, Col } from "antd";

import { FormItem, Label } from "./styles";

const Slider = ({ name, min, max, value, disabled, callback, unit }) => {
  const handleChangue = val => callback(val);

  //const handleAfterChange = val => callback(val);
  return (
    <FormItem>
      <Label>{name + " (" + unit + ")"}</Label>
      <Row type="flex" align="middle" style={{ margin: 0 }}>
        <Col span={16}>
          <SliderAntd
            disabled={disabled}
            min={min}
            max={max}
            tooltipVisible={false}
            onChange={handleChangue}
            value={typeof value === "number" ? value : 0}
          />
        </Col>
        <Col span={8}>
          <InputNumber
            style={{ width: "100%" }}
            min={min}
            max={max}
            value={typeof value === "number" ? value : 0}
            onChange={handleChangue}
            step={1}
            disabled={disabled}
            /*formatter={value => `${value} mm`}*/
          />
        </Col>
      </Row>
    </FormItem>
  );
};

export default Slider;
