const express = require("express");

const { queryAllArticles } = require("../controller/queryAllArticles");

const route = express.Router();

route.get("/", queryAllArticles);

module.exports = route;