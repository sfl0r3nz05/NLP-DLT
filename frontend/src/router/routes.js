import React from "react";
import List from "../pages/List";
import UserData from "../pages/UserData";
import Register from "../pages/Register";
import Agreement from "../pages/Agreement";
import AcceptAgreement from "../pages/AcceptAgreement";

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
    text: "ROAMING AGREEMENT",
    icon: "plus",
    menuShow: true,
    body: () => <Agreement />,
  },
  {
    key: "3",
    path: "/acceptagreement",
    exact: true,
    text: "ACCEPT RA",
    icon: "plus",
    menuShow: true,
    body: () => <AcceptAgreement />,
  },
  {
    key: "4",
    path: "/list",
    exact: true,
    text: "RAs AVAILABLE",
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
