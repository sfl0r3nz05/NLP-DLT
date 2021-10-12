import React, { useState, useEffect } from "react";
import { Table } from "antd";

const RenderTable = props => {
  //const { data, columns, onClick, expandedData } = props;

  const [data, setData] = useState([]);
  const [columns, setColumns] = useState([]);
  const [onClick, setOnClick] = useState(null);
  const [expandedData, setExpandedData] = useState(null);

  useEffect(() => {
    props.data && setData(props.data);
    props.columns && setColumns(props.columns);
    props.onClick && setOnClick(props.onClick);
    props.expandedData && setExpandedData(props.expandedData);
  }, [props]);

  return expandedData ? (
    <Table
      bordered
      dataSource={data}
      columns={columns}
      onRow={element => ({
        onClick: e => (onClick === undefined || onClick === null ? e.preventDefault() : onClick(element)) // Changue only of we have callback
      })}
      expandedRowRender={record => expandedData !== undefined && expandedData !== null && expandedData(record)}
      scroll={{ x: 1100 }}
      bodyStyle={{ overflowX: "overlay" }}
    />
  ) : (
    <Table //tabla sin expand
      bordered
      dataSource={data}
      columns={columns}
      onRow={element => ({
        onClick: e => (onClick === undefined || onClick === null ? e.preventDefault() : onClick(element)) // Changue only of we have callback
      })}
      scroll={{ x: 1100 }}
    />
  );
};

export default RenderTable;
