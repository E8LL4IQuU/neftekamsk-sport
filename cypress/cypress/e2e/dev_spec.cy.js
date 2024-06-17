// cypress/e2e/dev_spec.cy.js

if (Cypress.env('ENVIRONMENT') === 'dev') {

    Cypress.Commands.add('apiRequest', (method, url, body, headers = {}) => {
        return cy.request({
            method,
            url,
            body,
            headers,
            failOnStatusCode: false, // To handle and assert on failed requests
        });
    });

    describe('Dev environment tests', () => {
        let userID;
        const randomUsername = self.crypto.randomUUID();
        const randomPassword = self.crypto.randomUUID();

        before(() => {
            cy.request('POST', '/api/auth/register', {
                username: 'cypress',
                email: randomUsername,
                password: randomPassword
            }).then((response) => {
                expect(response.status).to.eq(200);
                userID = response.body.id
            });
        });

        after(() => {
            cy.request('DELETE', `/api/users/${userID}`).then((response) => {
                expect(response.status).to.eq(204);
            })
        })

        // TODO: add tests
        const routes = [
            {
                name: 'Login user',
                method: 'POST',
                url: '/api/auth/login',
                body: {
                    email: randomUsername,
                    password: randomPassword
                },
                headers: {},
                status: 200,
                assertions: (response) => {
                    expect(response.status).to.eq(200);
                    expect(response.body).to.have.property('id', userID);
                    expect(response.body).to.have.property('name', 'cypress');
                    expect(response.body).to.have.property('email', randomUsername);
                }
            },
            // {
            //     name: 'Get event by ID',
            //     method: 'GET',
            //     url: '/api/events/1', // Assuming there's an event with ID 1
            //     body: null,
            //     headers: {},
            //     status: 200,
            //     assertions: (response) => {
            //         expect(response.status).to.eq(200);
            //         // Add your assertions here
            //     }
            // },
            // {
            // {
            //     name: 'Delete event',
            //     method: 'DELETE',
            //     url: '/api/events/1', // Assuming there's an event with ID 1
            //     body: null,
            //     headers: {},
            //     status: 200,
            //     assertions: (response) => {
            //         expect(response.status).to.eq(200);
            //         expect(response.body).to.have.property('message', 'Record deleted successfully');
            //     }
            // },
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
}