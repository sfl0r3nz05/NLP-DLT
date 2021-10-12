import React from "react";
import List from "../pages/List";
import AddBatch from "../pages/AddBatch";
import UserData from "../pages/UserData";
import DelBatch from "../pages/DeleteBatch";
import DelToken from "../pages/DeleteToken";
import CreateToken from "../pages/CreateToken";
import UpdBatch from "../pages/UpdateBatch";
import TransferToken from "../pages/TransferToken";
import TransferValue from "../pages/TransferValue";
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
    path: "/create-token",
    exact: true,
    text: "CREAR TOKEN",
    icon: "plus",
    menuShow: true,
    body: () => <CreateToken />,
  },
  {
    key: "3",
    path: "/add-batch",
    exact: true,
    text: "AGREGAR LOTE",
    icon: "plus",
    menuShow: true,
    body: () => <AddBatch />,
  },
  {
    key: "4",
    path: "/update-batch",
    exact: true,
    text: "ACTUALIZAR LOTE",
    icon: "edit",
    menuShow: true,
    body: () => <UpdBatch />,
  },
  {
    key: "5",
    path: "/delete-batch",
    exact: true,
    text: "BORRAR LOTE",
    icon: "minus",
    menuShow: true,
    body: () => <DelBatch />,
  },
  {
    key: "6",
    path: "/delete-token",
    exact: true,
    text: "BORRAR TOKEN",
    icon: "minus",
    menuShow: true,
    body: () => <DelToken />,
  },
  {
    key: "7",
    path: "/transfer-token",
    exact: true,
    text: "TRANSFERIR TOKEN",
    icon: "shopping",
    menuShow: true,
    body: () => <TransferToken />,
  },
  {
    key: "8",
    path: "/transfer-value",
    exact: true,
    text: "COMPRAR TOKEN",
    icon: "euro",
    menuShow: true,
    body: () => <TransferValue />,
  },
  {
    key: "9",
    path: "/list",
    exact: true,
    text: "MARKETPLACE",
    icon: "unordered-list",
    menuShow: true,
    body: () => <List />,
  },
  {
    key: "10",
    path: "/user-data",
    exact: true,
    text: "DETALLES DEL USUARIO",
    icon: "user",
    menuShow: false,
    body: () => <UserData />,
  },
];
