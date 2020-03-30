const moment = require('moment');
const axios = require('axios');

module.exports = (data) => {
    console.log(moment().format('MMMM Do YYYY, h:mm:ss a'));

    data.foo = 1;

    return data;
};

