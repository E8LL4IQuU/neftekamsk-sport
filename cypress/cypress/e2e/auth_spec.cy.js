// cypress/e2e/dev_spec.cy.js

let userID;
const randomUsername = self.crypto.randomUUID();
const randomPassword = self.crypto.randomUUID();

describe('Register API', () => {
        it('registers successfully', () => {
            cy.request('POST', '/api/auth/register', {
                username: 'cypress',
                email: randomUsername,
                password: randomPassword
            }).then((response) => {
                expect(response.status).to.eq(200);
                userID = response.body.id
            })
        })

    })

    describe('Login UI', () => {

        beforeEach(() => {
            cy.visit('/login');
        });

        it('displays the login form', () => {
            cy.get('h2').contains('Sign in to your account');
            cy.get('input[name="email"]').should('be.visible');
            cy.get('input[name="password"]').should('be.visible');
            cy.get('button[type="submit"]').contains('Sign in');
        });

        it('logs in successfully with valid credentials', () => {
            cy.intercept('POST', '**/api/auth/login', {
                statusCode: 200,
                body: { success: true },
            }).as('loginRequest');

            cy.get('input[name="email"]').type(randomUsername);
            cy.get('input[name="password"]').type(randomPassword);
            cy.get('button[type="submit"]').click();

            cy.wait('@loginRequest').its('response.statusCode').should('eq', 200);
            cy.url().should('include', '/manage');
        });

        it('shows error for invalid credentials', () => {
            cy.intercept('POST', '**/api/auth/login', {
                statusCode: 400,
                body: { error: 'Invalid credentials' },
            }).as('loginRequest');

            cy.get('input[name="email"]').type(randomUsername);
            cy.get('input[name="password"]').type('wrongpassword');
            cy.get('button[type="submit"]').click();

            cy.wait('@loginRequest').its('response.statusCode').should('eq', 400);
            cy.get('.Vue-Toastification__toast').contains('Неверный пароль');
        });

        it('shows error for non-existing user', () => {
            cy.intercept('POST', '**/api/auth/login', {
                statusCode: 404,
                body: { error: 'User not found' },
            }).as('loginRequest');

            cy.get('input[name="email"]').type('nonexistent@example.com');
            cy.get('input[name="password"]').type(randomPassword);
            cy.get('button[type="submit"]').click();

            cy.wait('@loginRequest').its('response.statusCode').should('eq', 404);
            cy.get('.Vue-Toastification__toast').contains('Пользователь не найден');
        });

        it('shows server error', () => {
            cy.intercept('POST', '**/api/auth/login', {
                statusCode: 500,
                body: { error: 'Internal Server Error' },
            }).as('loginRequest');

            cy.get('input[name="email"]').type(randomUsername);
            cy.get('input[name="password"]').type(randomPassword);
            cy.get('button[type="submit"]').click();

            cy.wait('@loginRequest').its('response.statusCode').should('eq', 500);
            cy.get('.Vue-Toastification__toast').contains('Ошибка сервера');
        });
    });

