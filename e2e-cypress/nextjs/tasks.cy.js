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
  cy.get("ul#task-list").then((taskList) => {
    if (taskList.find('li[id^="task-"]').length === 0) {
      return;
    }

    taskList.find('li[id^="task-"]').each((task) => {
      task.find("button").click();
    });
  });
};

const deleteFirstTask = (cy) => {
  cy.get("ul#task-list").then((taskList) => {
    if (taskList.find('li[id^="task-"]').length === 0) {
      return;
    }

    taskList.find('li[id^="task-"]').find("button").click();
  });
};

const undo = (cy) => {
  cy.get("button").contains("Undo").click();
};

const assertNumTasks = (cy, numTasks) => {
  cy.get("ul#task-list")
    .find('li[id^="task-"]')
    .should("have.length", numTasks);
};

describe("Managing tasks", () => {
  it("should get task list", () => {
    cy.log("should log in");

    cy.visit(Cypress.env("FRONTEND_BASE_URL") || "http://localhost:3000");

    cy.log("Logging in");
    cy.get("div").contains("Login").click();

    cy.log("Authenticating as Alice");
    cy.origin(Cypress.env("AUTH0_ISSUER_BASE_URL"), () => {
      cy.get("input#1-email").type(Cypress.env("ALICE_USERNAME"));
      cy.get("input#1-password").type(Cypress.env("ALICE_PASSWORD"));

      cy.log("Submitting login form");
      cy.get("button#1-submit").click();
    });

    cy.get("h3").contains("Tasks").should("exist");
    cy.get("ul#task-list").should("exist");
  });

  it("should add and delete tasks", () => {
    cy.log("should log in");

    cy.visit(Cypress.env("FRONTEND_BASE_URL") || "http://localhost:3000");

    cy.log("Logging in");
    cy.get("div").contains("Login").click();

    cy.log("Authenticating as Alice");
    cy.origin(Cypress.env("AUTH0_ISSUER_BASE_URL"), () => {
      cy.get("input#1-email").type(Cypress.env("ALICE_USERNAME"));
      cy.get("input#1-password").type(Cypress.env("ALICE_PASSWORD"));

      cy.log("Submitting login form");
      cy.get("button#1-submit").click();
    });

    deleteAllTasks(cy);
    addTasks(cy, 3);
    assertNumTasks(cy, 3);
    deleteFirstTask(cy);
    assertNumTasks(cy, 2);
    undo(cy);
    assertNumTasks(cy, 3);
  });
});
