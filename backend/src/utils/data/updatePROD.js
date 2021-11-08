var fs = require('fs')
const readBuffer = require("../buffer/readBuffer");

module.exports = async function updatePROD(valueToUpdate) {
    try {
        value = JSON.parse(valueToUpdate);
        console.log(value);

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
                    "variables": (Buffer.from(value.variables, 'base64')).toString('utf-8'),
                    "variations": (Buffer.from(value.variations, 'base64')).toString('utf-8'),
                    "stdclauses": (Buffer.from(value.stdclauses, 'base64')).toString('utf-8'),
                    "customtexts": (Buffer.from(value.customtexts, 'base64')).toString('utf-8'),
                }
                output[0].articles.push(base)
            }
            var json = JSON.stringify(objs)
            fs.writeFile(__dirname + "/../../data/listOfMNOs.json", json, function (err) {
                if (err) throw err;
            })
        })
    } catch (error) {
        console.log(error);
    }
}