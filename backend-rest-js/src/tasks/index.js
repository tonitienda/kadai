const makeTaskRouter = require("./router");

const inMemoryDatasource = require("./db-in-memory");

module.exports = makeTaskRouter(inMemoryDatasource);
