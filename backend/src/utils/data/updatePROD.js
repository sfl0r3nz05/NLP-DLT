var fs = require('fs')
const readBuffer = require("../buffer/readBuffer");

module.exports = async function updatePROD(valueToUpdate) {
    try {
        value = JSON.parse(valueToUpdate);
        console.log(value.articleno.replace(/['"]+/g, ''));
        console.log(value.articlename.replace(/['"]+/g, ''));

        fs.readFile(__dirname + "/../../data/listOfMNOs.json", function (err, data) {
            var objs = JSON.parse(data)
            var output = objs.filter(function (obj) { return obj.ra_id == value.raid; })
            output[0].ra_status = value.rastatus
            output[0].timestamp = value.timestamp
            if (value.articleno) {
                var base = {
                    "articleId": value.articleno.replace(/['"]+/g, ''),
                    "articleName": value.articlename.replace(/['"]+/g, ''),
                    "articleStatus": value.articlestatus,
                }
                output[0].articles.push(base)
            }
            var json = JSON.stringify(objs)
            fs.writeFile(__dirname + "/../../data/listOfMNOs.json", json, function (err) {
                if (err) throw err;
            })
        })
    } catch (error) {

    }
}