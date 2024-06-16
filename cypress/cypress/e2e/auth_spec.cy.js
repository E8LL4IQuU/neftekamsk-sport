// cypress/e2e/auth_spec.js

describe('API Auth Register', () => {
    it('should return 401 Unauthorized if env is not set to "dev"', () => {
      // Check if the environment variable 'env' is not set to 'dev'
      // Here's some fricking normalization, so vars can only start with CYPRESS_, CYPRESS_ENV is actually ENV for some reason, but it still won't work as it's reserved or something also it's camelCase otherwise
      if (Cypress.env('ENVIRONMENT') !== 'dev') {
        // Make a call to /api/auth/register
        cy.request({
          method: 'POST',
          url: '/api/auth/register',
          failOnStatusCode: false, // Prevent Cypress from failing the test on status code other than 2xx or 3xx
        }).then((response) => {
          // Confirm that status is 401 Unauthorized
          expect(response.status).to.eq(401);
        });
      } else {
        cy.log('Skipping test since env is set to "dev"');
      }
    });
  });
  