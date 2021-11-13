import React, { createContext, useReducer, useEffect } from "react";
import { withRouter } from "react-router-dom";

// Router
import Router from "./router";
import Login from "./pages/Login";

import "./App.less";

import { ConfigProvider } from 'antd';
import locale from 'antd/lib/locale/es_ES';

require('dotenv').config()

export const AuthContext = createContext();

const initialState = {
  isAuthenticated: false,
  user: null,
  token: null
};

const reducer = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      localStorage.setItem("user", JSON.stringify(action.payload.user));
      localStorage.setItem("token", JSON.stringify(action.payload.token));
      return {
        ...state,
        isAuthenticated: true,
        user: action.payload.user,
        token: action.payload.token
      };
    case "LOGOUT":
      localStorage.clear();
      return {
        ...state,
        isAuthenticated: false,
        user: null
      };
    default:
      return state;
  }
};

function App({ history }) {
  const [state, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    const user = JSON.parse(localStorage.getItem("user") || null);

    if (user) {
      dispatch({
        type: "LOGIN",
        payload: {
          user
        }
      });
    }
  }, []);

  return (
    <AuthContext.Provider
      value={{
        state,
        dispatch
      }}
    >
      <ConfigProvider locale={locale}>
        <div className="App">{!state.isAuthenticated ? <Login /> : <Router history={history} />}</div>
      </ConfigProvider>
    </AuthContext.Provider>
  );
}

export default withRouter(App);

