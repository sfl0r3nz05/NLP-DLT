var fs = require('fs').promises;
const query = require("../query/query");
//const updateARTICLE = require("../../utils/data/updateARTICLE");

module.exports = async function queryAllArticles(user) {
  try {
    let data = await fs.readFile(__dirname + "/../../data/listOfMNOs.json")
    var obj = JSON.parse(data)

    for (let index = 0; index < obj.length; index++) {
      let method = "queryAllArticles";
      let noArg = 1
      value = await query(method, noArg, obj[index].ra_id, user);
      var obj2 = JSON.parse(value)
      obj[index].articles = obj2.articles
      //await updateARTICLE(obj[index]);
    }

  } catch (error) {
    console.error(error);
    return error
  }
};