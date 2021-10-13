const express = require("express");

// Controllers
const { proposeAddArticle } = require("../controller/proposeAddArticle");

const route = express.Router();

route.post("/", proposeAddArticle);

module.exports = route;