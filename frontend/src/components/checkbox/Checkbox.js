import React from "react";
import { Checkbox } from "antd";

export default ({ disabled, label, value, handleChange }) => (
  <Checkbox
    style={{ display: "block", marginLeft: 0 }}
    disabled={disabled || false}
    label={label}
    checked={value}
    onChange={handleChange}
  >
    {label}
  </Checkbox>
);