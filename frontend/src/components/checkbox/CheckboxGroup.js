import React from "react";
import Checkbox from "./Checkbox";

export default ({ options, ...props }) => (
  <div
    className="ant-checkbox-group"
    style={{ display: "inline-block", marginRight: 10 }}
  >
    {options.map(label => (
      <Checkbox
        key={label}
        label={label}
        disabled={props.disabled}
        handleChange={props.handleChange}
        value={props[label]}
      />
    ))}
  </div>
);