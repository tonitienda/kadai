describe("Kadai Basic", () => {
  it("should render the home page", () => {
    cy.visit(Cypress.env("FRONTEND_BASE_URL") || "http://localhost:3000");
    cy.get("div").contains("Kadai");
  });
});
