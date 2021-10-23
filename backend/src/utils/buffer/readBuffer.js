
const fs = require("fs");
const { promisify } = require("util");
const readFile = promisify(fs.readFile);

module.exports = async function readBuffer(selectEnv) {
  try {
    let BUFFERPATH;
    if (selectEnv === 1) {
      BUFFERPATH = process.env.PATH_TO_BUFFER_PRODUCTS;
    }
    if (selectEnv === 2) {
      BUFFERPATH = process.env.PATH_TO_BUFFER_COMPANIES;
    }
    if (selectEnv === 3) {
      BUFFERPATH = process.env.PATH_TO_BUFFER_USERS;
    }
    if (selectEnv === 4) {
      BUFFERPATH = process.env.PATH_TO_BUFFER_COUNTRY;
    }
    const data = await readFile(
      __dirname + `${BUFFERPATH}`,
      "utf8"
    );
    const dataParsed = JSON.parse(data);
    return dataParsed;
  } catch (error) {
    console.error(error);
  }
};