const express = require("express");

// Controllers
const { proposeUpdateArticle } = require("../controller/proposeUpdateArticle");

const route = express.Router();

route.post("/", proposeUpdateArticle);

module.exports = route;