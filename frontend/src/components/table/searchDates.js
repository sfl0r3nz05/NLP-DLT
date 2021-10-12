import React, { useState } from "react";

import { Button, Icon, DatePicker } from "antd";
import moment from "moment";

const { RangePicker } = DatePicker;
const defaultDatesSearch = [
  moment().set({ year: 2019, month: 9, date: 1, hour: 0, minute: 0, second: 0 }),
  moment().set({ hour: 23, minute: 59, second: 59 })
];

const getRanges = () => {
  return {
    "Este mes": [moment().startOf("month"), moment().endOf("month")],
    "Esta semana": [moment().startOf("week"), moment().endOf("week")],
    Ayer: [
      moment()
        .set({ hour: 0, minute: 0, second: 0 })
        .add(-1, "day"),
      moment()
        .set({ hour: 23, minute: 59, second: 59 })
        .add(-1, "day")
    ],
    "Últimas 24 horas": [moment().subtract(24, "h"), moment()],
    "Últimas 12 horas": [moment().subtract(12, "h"), moment()],
    "Últimas 6 horas": [moment().subtract(6, "h"), moment()],
    "Últimas 3 horas": [moment().subtract(3, "h"), moment()],
    "Última hora": [moment().subtract(1, "h"), moment()],
    "Últimos 30 minutos": [moment().subtract(30, "minutes"), moment()],
    "Últimos 15 minutos": [moment().subtract(15, "minutes"), moment()],
    "Últimos 5 minutos": [moment().subtract(5, "minutes"), moment()]
  };
};

const TableSearchDates = dataIndex => {
  const [dates, setDates] = useState(defaultDatesSearch);

  const getColumnSearchProps = dataIndex => ({
    filterDropdown: ({ setSelectedKeys, selectedKeys, confirm, clearFilters }) => (
      <div style={{ padding: 8 }}>
        <RangePicker
          ranges={getRanges()}
          showTime
          format="DD/MM/YYYY"
          defaultValue={dates}
          onChange={values => setSelectedKeys(values)}
          onPressEnter={() => handleSearch(selectedKeys, confirm, dataIndex)}
          style={{ width: 215, marginBottom: 8, display: "block"}}
        />
        <Button
          type="primary"
          onClick={() => handleSearch(selectedKeys, confirm, dataIndex)}
          icon="search"
          size="small"
          style={{ width: 90, marginRight: 8 }}
        >
          Buscar
        </Button>
        <Button onClick={() => handleReset(clearFilters)} size="small" style={{ width: 90 }}>
          Reset
        </Button>
      </div>
    ),
    filterIcon: filtered => <Icon type="search" style={{ color: filtered ? "#1890ff" : undefined }} />,
    onFilter: (_, record) =>{ let d = record[dataIndex]*1000; return moment(d).isBetween(dates[0], dates[1])},
    render: text => text
  });

  const handleSearch = (selectedKeys, confirm, dataIndex) => {
    confirm();
    setDates(selectedKeys);
  };

  const handleReset = clearFilters => {
    clearFilters();
    setDates(defaultDatesSearch);
  };

  return getColumnSearchProps(dataIndex);
};

export default TableSearchDates;
