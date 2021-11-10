const recoverMNO = require("../../utils/recoverMNO");
const updatePROD = require("../../utils/data/updatePROD");
const readBuffer = require("../../utils/buffer/readBuffer");
const invokeEvents = require("../../utils/invoke/invokeEvents");

const proposeDeleteArticle = async (req, res) => {
    try {
        let data = req.body; // params from POST
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
        noArgs = 3
        let arg1 = output[0].ra_id
        let arg2 = value.articleNo
        let arg3 = value.articleName
        let arg4 = "";
        let arg5 = "";
        let arg6 = "";
        let arg7 = "";
        let method = "proposeDeleteArticle";
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

module.exports = { proposeDeleteArticle };

