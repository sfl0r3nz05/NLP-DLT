import React from "react";
import List from "../pages/List";
import Agreement from "../pages/Agreement";
import UserData from "../pages/UserData";
import Register from "../pages/Register";

export const routes = [
  {
    key: "1",
    path: "/",
    exact: true,
    text: "REGISTER MNO",
    icon: "plus",
    menuShow: true,
    body: () => <Register />,
  },
  {
    key: "2",
    path: "/agreement",
    exact: true,
    text: "MANAGE AGREEMENT",
    icon: "plus",
    menuShow: true,
    body: () => <Agreement />,
  },
  {
    key: "3",
    path: "/list",
    exact: true,
    text: "MARKETPLACE",
    icon: "unordered-list",
    menuShow: true,
    body: () => <List />,
  },
  {
    key: "4",
    path: "/user-data",
    exact: true,
    text: "DETALLES DEL USUARIO",
    icon: "user",
    menuShow: false,
    body: () => <UserData />,
  },
];
