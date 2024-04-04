const { defineConfig } = require("cypress");
require("dotenv").config();

module.exports = defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
    supportFile: false,
    specPattern: "**/*.cy.js",
  },
  env: {
    FRONTEND_BASE_URL: process.env.FRONTEND_BASE_URL,
    AUTH0_ISSUER_BASE_URL: process.env.AUTH0_ISSUER_BASE_URL,
    AUTH0_USERNAME: process.env.AUTH0_USERNAME,
    AUTH0_PASSWORD: process.env.AUTH0_PASSWORD,
  },
});
