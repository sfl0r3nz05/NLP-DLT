const express = require("express");

const { querySingleArticle } = require("../controller/querySingleArticle");

const route = express.Router();

route.get("/", querySingleArticle);

module.exports = route;