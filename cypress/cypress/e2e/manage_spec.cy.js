if (Cypress.env('ENVIRONMENT') === 'dev') {
  let userID;
  const randomUsername = self.crypto.randomUUID();
  const randomPassword = self.crypto.randomUUID();

  const loginViaAPI = () => {
    return cy.request({
      method: 'POST',
      url: '/api/auth/login',
      body: {
        email: randomUsername,
        password: randomPassword
      },
    });
  };

  it('registers a new user in dev mode', () => {
    cy.request('POST', '/api/auth/register', {
      username: 'cypress',
      email: randomUsername,
      password: randomPassword
    }).then((response) => {
      userID = response.body.id;
    });
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

  describe('Manage routes', () => {
    // Ensure user is logged out before each test
    beforeEach(() => {
      cy.clearCookies();
    });

    it('Redirects to login page if not authenticated', () => {
      cy.visit('/manage');
      cy.url().should('include', '/login');
    });

    it('Logs in and accesses /manage route', () => {
      loginViaAPI().then(() => {
        cy.visit('/manage');
        cy.url().should('include', '/manage');
        cy.get('body').contains('Мероприятия'); // Check for specific element in Manage view
      });
    });

    it('Access /manage/events and check elements', () => {
      loginViaAPI().then(() => {
        cy.visit('/manage/events');
        cy.url().should('include', '/manage/events');
        cy.get('h1').contains('Мероприятия');
      });
    });

    it('Access /manage/news and check elements', () => {
      loginViaAPI().then(() => {
        cy.visit('/manage/news');
        cy.url().should('include', '/manage/news');
        cy.get('h1').contains('Новости');
      });
    });
  });

  describe('Manage Events', () => {
    const newEvent = {
      name: 'Test Event',
      description: 'This is a test event.',
      date: '2024-07-01',
    };

    const updatedEvent = {
      name: 'Updated Test Event',
      description: 'This is an updated test event.',
    };

    it('should create a new event', () => {
      loginViaAPI().then(() => {
        cy.visit('/manage/events');
        cy.contains('button', 'Создать').click();
        cy.url().should('include', '/manage/events/create');

        cy.get('input[placeholder="Название мероприятия"]').type(newEvent.name); // Adjust the placeholder if needed
        cy.get('textarea[placeholder="Начните писать описание..."]').type(newEvent.description);

        // // Fill the date field
        // cy.get('input[type="date"]').type(newEvent.date);

        // Handle file upload
        cy.get('input[type="file"]')
          .attachFile('test-image.png');

        cy.contains('button', 'Применить').click();

        cy.get('input[name="eventName"]').should('have.value', newEvent.name);
        cy.get('input[name="eventDescription"]').should('have.value', newEvent.description);
      })
    });

  });

  describe('Post-testing cleanup', () => {
    it('delete user that we created', () => {
      loginViaAPI().then(() => {
        cy.request({
          method: 'DELETE',
          url: `/api/users/${userID}`
        });
      })
    })
  })

}