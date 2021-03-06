var fs = require('fs')
const readBuffer = require("../buffer/readBuffer");

module.exports = async function updatePROD(valueToUpdate) {
    try {
        value = JSON.parse(valueToUpdate);
        console.log(value);

        fs.readFile(__dirname + "/../../data/listOfMNOs.json", function (err, data) {
            var objs = JSON.parse(data)
            var output = objs.filter(function (obj) { return obj.ra_id == value.raid; })
            if (value.rastatus != "") { output[0].ra_status = value.rastatus }
            output[0].timestamp = value.timestamp
            var articles = output[0].articles
            var article = (articles).filter(function (article) { return article.articleId == value.articleno })

            if (value.articlestatus == "added_article") {
                var variablesParsed = (Buffer.from(value.variables, 'base64')).toString('utf-8')
                var variationsParsed = (Buffer.from(value.variations, 'base64')).toString('utf-8')
                var stdclausesParsed = (Buffer.from(value.stdclauses, 'base64')).toString('utf-8')
                var customtextsParsed = (Buffer.from(value.customtexts, 'base64')).toString('utf-8')

                var base = {
                    "articleId": value.articleno.replace(/['"]+/g, ''),
                    "articleName": value.articlename.replace(/['"]+/g, ''),
                    "proposedBy": value.mno1,
                    "articleStatus": value.articlestatus,
                    "variables": JSON.parse(variablesParsed),
                    "variations": JSON.parse(variationsParsed),
                    "stdclauses": JSON.parse(stdclausesParsed),
                    "customtexts": JSON.parse(customtextsParsed),
                }
                output[0].articles.push(base)
            }
            if (value.articlestatus == "accepted_changes") {
                article[0].articleStatus = value.articlestatus
                article[0].proposedBy = value.mno1
            }
            if (value.articlestatus == "rejected_changes") {
                article[0].articleStatus = value.articlestatus
                article[0].proposedBy = value.mno1
            }
            if (value.articlestatus == "proposed_changes") {
                var variablesParsed = (Buffer.from(value.variables, 'base64')).toString('utf-8')
                var variationsParsed = (Buffer.from(value.variations, 'base64')).toString('utf-8')
                var stdclausesParsed = (Buffer.from(value.stdclauses, 'base64')).toString('utf-8')
                var customtextsParsed = (Buffer.from(value.customtexts, 'base64')).toString('utf-8')

                article[0].proposedBy = value.mno1
                article[0].articleStatus = value.articlestatus
                article[0].variables = JSON.parse(variablesParsed)
                article[0].variations = JSON.parse(variationsParsed)
                article[0].stdclauses = JSON.parse(stdclausesParsed)
                article[0].customtexts = JSON.parse(customtextsParsed)
            }

            if (value.rastatus == "accepted_ra") {
                output[0].ra_status = value.rastatus
                output[0].acceptRAproposedBy = value.mno1
                articles[0].proposedBy = value.mno1
            }

            if (value.rastatus == "confirmation_accepted_ra") {
                output[0].ra_status = value.rastatus
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