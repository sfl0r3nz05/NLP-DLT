var fs = require('fs')
const readBuffer = require("../buffer/readBuffer");

module.exports = async function updatePROD(valueToUpdate) {
    try {
        value = JSON.parse(valueToUpdate);
        selectEnv = 1; //Detect flag
        objs = await readBuffer(selectEnv);
        var output = objs.filter(function (obj) { return ((obj.mno1 == value.mno1 && obj.mno2 == value.mno2) || (obj.mno1 == value.mno2 && obj.mno2 == value.mno1)); })

        fs.readFile(__dirname + "/../../data/listOfMNOs.json", function (err, data) {
            var obj = JSON.parse(data)
            obj.push({
                "mno1": value.mno1,
                "country_mno1": country1.flag,
                "mno2": value.mno2,
                "country_mno2": country2.flag,
                "ra_name": value.raname,
                "ra_id": value.raid,
                "ra_status": value.rastatus,
                "timestamp": value.timestamp
            })
            var json = JSON.stringify(obj)
            fs.writeFile(__dirname + "/../../data/listOfMNOs.json", json, function (err) {
                if (err) throw err;
            })
        })
        console.log(ler);
    } catch (error) {

    }
}