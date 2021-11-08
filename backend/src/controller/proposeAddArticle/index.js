const recoverMNO = require("../../utils/recoverMNO");
const updatePROD = require("../../utils/data/updatePROD");
const readBuffer = require("../../utils/buffer/readBuffer");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const proposeAddArticle = async (req, res) => {
    try {
        let data = req.body; // params from POST
        let value = data.addArticle;

        let value4 = JSON.stringify(data.formVariables)
        let arg4 = Buffer.from(value4).toString('base64');

        let value5 = JSON.stringify(data.selectedArticlesVariation);
        //let arg5 = Buffer.from(value5).toString('base64');
        arg5 = ""

        let value6 = JSON.stringify(data.selectedArticlesStdClause)
        //let arg6 = Buffer.from(value6).toString('base64');
        arg6 = ""

        let value7 = JSON.stringify(data.formCustomText)
        let arg7 = Buffer.from(value7).toString('base64');

        let user = data.userDetails;

        if (!user || !value.raname || !value.articleNo) {
            res.sendStatus(201);
            res.end("201");
            return
        }
        let mno = await recoverMNO(user.username)
        const selectEnv = 1;
        const objs = await readBuffer(selectEnv);
        var output = objs.filter(function (obj) { return ((obj.mno1 == mno || obj.mno2 == mno) && obj.ra_name == value.raname) })
        if (!output[0]) {
            res.sendStatus(202);
            res.end("202");
            return
        }
        noArgs = 7
        let arg1 = output[0].ra_id
        let arg2 = value.articleNo
        let arg3 = value.articleName
        let method = "proposeAddArticle";
        let event_name = "proposed_add_article"

        eventHf = await invokeEvents(method, event_name, noArgs, arg1, arg2, arg3, arg4, arg5, arg6, arg7, user.username);
        if (!eventHf[0]) {
            res.sendStatus(403);
            res.end("403");
            return
        }
        await updatePROD(eventHf[1])
        res.sendStatus(200);
    } catch (error) {
        console.error(error);
        res.sendStatus(400);
    }
};

module.exports = { proposeAddArticle };
