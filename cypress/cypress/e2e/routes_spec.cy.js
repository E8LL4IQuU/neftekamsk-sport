// cypress/e2e/routes_spec.cy.js

Cypress.Commands.add('apiRequest', (method, url, body, headers = {}) => {
  return cy.request({
    method,
    url,
    body,
    headers,
    failOnStatusCode: false, // To handle and assert on failed requests
  });
});

describe('API Endpoints', () => {
  const routes = [
    {
      name: 'Should not register in production',
      method: 'POST',
      url: '/api/auth/register',
      status: Cypress.env('ENVIRONMENT') === 'dev' ? 400 : 401
    },
    // {
    //   name: 'Register user in dev mode',
    //   method: 'POST',
    //   url: '/api/auth/register',
    //   body: {
    //     // Your registration data
    //   },
    //   headers: {},
    //   status: Cypress.env('ENVIRONMENT') === 'dev' ? 200 : 404,
    //   assertions: (response) => {
    //     if (Cypress.env('ENVIRONMENT') === 'dev') {
    //       expect(response.status).to.eq(200);
    //       // Add your assertions here
    //     } else {
    //       expect(response.status).to.eq(404);
    //       // Add your assertions here for 404
    //     }
    //   }
    // },
    // {
    //   name: 'Login user',
    //   method: 'POST',
    //   url: '/api/auth/login',
    //   body: {
    //     email: 'test@example.com',
    //     password: 'password123'
    //   },
    //   headers: {},
    //   status: 200,
    //   assertions: (response) => {
    //     expect(response.status).to.eq(200);
    //     expect(response.body).to.have.property('token');
    //   }
    // },
    {
      name: 'Should not give user info to unauthorized clients',
      method: 'GET',
      url: '/api/auth/user',
      body: null,
      headers: {
        Authorization: 'Bearer YOUR_TOKEN_HERE'
      },
      status: Cypress.env('ENVIRONMENT') === 'dev' ? 200 : 401
    },
    // {
    //   name: 'Logout user',
    //   method: 'POST',
    //   url: '/api/auth/logout',
    //   body: null,
    //   headers: {},
    //   status: 200,
    //   assertions: (response) => {
    //     expect(response.status).to.eq(200);
    //     expect(response.body).to.have.property('message', 'Logged out successfully');
    //   }
    // },
    {
      name: 'Get all events',
      method: 'GET',
      url: '/api/events',
      body: null,
      headers: {},
      status: 200,
      assertions: (response) => {
        expect(response.status).to.eq(200);
        // Add your assertions here
      }
    },
    // {
    //   name: 'Get event by ID',
    //   method: 'GET',
    //   url: '/api/events/1', // Assuming there's an event with ID 1
    //   body: null,
    //   headers: {},
    //   status: 200,
    //   assertions: (response) => {
    //     expect(response.status).to.eq(200);
    //     // Add your assertions here
    //   }
    // },
    // {
    //   name: 'Should not allow creating events to unauthorized clients',
    //   method: 'POST',
    //   url: '/api/events',
    //   body: {
    //     // Your event data
    //   },
    //   headers: {},
    //   status: 401
    // },
    // {
    //   name: 'Should not allow updating events to unauthorized clients',
    //   method: 'PUT',
    //   url: '/api/events/1', // Assuming there's an event with ID 1
    //   body: {
    //     // Your updated event data
    //   },
    //   headers: {},
    //   status: 401,
    //   assertions: (response) => {
    //     expect(response.status).to.eq(200);
    //     // Add your assertions here
    //   }
    // },
    // {
    //   name: 'Delete event',
    //   method: 'DELETE',
    //   url: '/api/events/1', // Assuming there's an event with ID 1
    //   body: null,
    //   headers: {},
    //   status: 200,
    //   assertions: (response) => {
    //     expect(response.status).to.eq(200);
    //     expect(response.body).to.have.property('message', 'Record deleted successfully');
    //   }
    // },
    // TODO: Repeat the above structure for other routes
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
