// cypress/e2e/production_spec.cy.js

if (Cypress.env('ENVIRONMENT') !== 'dev') {

  Cypress.Commands.add('apiRequest', (method, url, body, headers = {}) => {
    return cy.request({
      method,
      url,
      body,
      headers,
      failOnStatusCode: false, // To handle and assert on failed requests
    });
  });

  describe('Production environment test', () => {
    // TODO: add tests
    const routes = [
      {
        name: 'disallow registering in production',
        method: 'POST',
        url: '/api/auth/register',
        status: 401
      },
      {
        name: 'disallow authorized state to unauthorized clients',
        method: 'GET',
        url: '/api/auth/user',
        status: 401
      },
      {
        name: 'disallow deleting users to unauthorized clients',
        method: 'DELETE',
        url: '/api/auth/user/1',
        status: 401
      },

      {
        name: 'disallow creating events to unauthorized clients',
        method: 'POST',
        url: '/api/events',
        status: 401
      },
      {
        name: 'disallow updating events to unauthorized clients',
        method: 'PUT',
        url: '/api/events/1',
        status: 401,
      },
      {
        name: 'disallow deleting events to unauthorized clients',
        method: "DELETE",
        url: '/api/events/1'
      },

      {
        name: 'get events',
        method: 'GET',
        url: '/api/events',
        body: null,
        headers: {},
        status: 200,
      },
    ];

    routes.forEach((route) => {
      it(`should ${route.name}`, () => {
        cy.request({
          method: route.method,
          url: `${route.url}`,
          body: route.body,
          headers: route.headers,
          failOnStatusCode: false
        }).then((response) => {
          expect(response.status).to.eq(route.status);
          if (route.assertions) {
            route.assertions(response);
          }
        });
      });
    });
  });
}