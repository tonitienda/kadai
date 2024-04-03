const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
    supportFile: false,
    specPattern: "**/*.cy.js",
  },
  env: {
    KADAI_FRONTEND_URL: process.env.KADAI_FRONTEND_URL,
  },
});
