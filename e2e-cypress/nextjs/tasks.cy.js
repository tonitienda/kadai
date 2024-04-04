const addTasks = (cy, numTasks) => {
  for (let i = 0; i < numTasks; i++) {
    cy.get("button").contains("Add Task").click();
    cy.get("input#task-title").type(`Task ${i + 1}`);
    cy.get("input#task-description").type(
      `This is the description of task ${i + 1}`
    );
    cy.get("button").contains("Save").click();
  }

  cy.get("ul#task-list")
    .find('li[id^="task-"]')
    .should("have.length", numTasks);
};

const deleteAllTasks = (cy) => {
  cy.get("ul#task-list")
    .find('li[id^="task-"]')
    .each((task) => {
      task.find("button").click();
    });
};

describe("Managing tasks", () => {
  it("should login", () => {
    cy.visit(Cypress.env("FRONTEND_BASE_URL") || "http://localhost:3000");

    cy.get("div").contains("Login").click();

    cy.origin(Cypress.env("AUTH0_ISSUER_BASE_URL"), () => {
      cy.get("input#1-email").type(Cypress.env("AUTH0_USERNAME"));
      cy.get("input#1-password").type(Cypress.env("AUTH0_PASSWORD"));

      cy.get("button#1-submit").click();
    });

    deleteAllTasks(cy);
    addTasks(cy, 3);
  });
});
