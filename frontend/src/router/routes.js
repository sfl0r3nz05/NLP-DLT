import React from "react";
import List from "../pages/List";
import UserData from "../pages/UserData";
import RegisterMNO from "../pages/RegisterMNO";
import Agreement from "../pages/Agreement";
import AddArticle from "../pages/AddArticle";

export const routes = [
  {
    key: "1",
    path: "/",
    exact: true,
    text: "REGISTER MNO",
    icon: "plus",
    menuShow: true,
    body: () => <RegisterMNO />,
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
    path: "/addarticle",
    exact: true,
    text: "PROPOSE ARTICLE",
    icon: "plus",
    menuShow: true,
    body: () => <AddArticle />,
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
    key: "5",
    path: "/user-data",
    exact: true,
    text: "DETALLES DEL USUARIO",
    icon: "user",
    menuShow: false,
    body: () => <UserData />,
  },
];
