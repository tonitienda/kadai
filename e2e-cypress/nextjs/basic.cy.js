describe("Next.js Basic", () => {
  it("should render the home page", () => {
    cy.visit(Cypress.env("KADAI_FRONTEND_URL") || "http://localhost:3000");
    cy.get("div").contains("Kadai");
  });
});
