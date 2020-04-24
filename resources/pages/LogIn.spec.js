import { render, fireEvent } from '@testing-library/svelte';
import { get } from 'svelte/store';

import LogIn from './LogIn.svelte';
import { jwt, notifications } from '../stores';

describe('Log in page', () => {
  it('has a Yuzu heading', () => {
    const { getByText } = render(LogIn);
    const heading = getByText('Yuzu');
    expect(heading).toBeInstanceOf(HTMLHeadingElement);
  });
  it('sends the log in form', async () => {
    // Mock the fetch response
    global.fetch = jest.fn().mockImplementation(async () => ({
      ok: true,
      json: async () => ({ token: 'test-token' }),
    }));
    // Render and interact with DOM
    const { getByLabelText, getByText } = render(LogIn);
    const userInput = getByLabelText('User');
    const passwordInput = getByLabelText('Password');
    await Promise.all([
      fireEvent.input(userInput, { target: { value: 'my-user' } }),
      fireEvent.input(passwordInput, { target: { value: 'my-password' } }),
    ]);
    const button = getByText('Log in');
    await fireEvent.click(button);
    // Assert fetch call
    expect(global.fetch).toHaveBeenCalledTimes(1);
    expect(global.fetch).toHaveBeenCalledWith('/api/v1/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        user: 'my-user',
        password: 'my-password',
      }),
    });
    // Check if jwt in the state was updated
    expect(get(jwt)).toBe('test-token');
    // Check if the user was notified
    expect(
      get(notifications).find((n) => n.message === 'Could not log in'),
    ).not.toBeNull();
  });
  it('notifies on login fail', async () => {
    global.fetch = jest.fn().mockImplementation(async () => ({ ok: false }));
    const { getByText } = render(LogIn);
    const button = getByText('Log in');
    await fireEvent.click(button);
    expect(
      get(notifications).find((n) => n.message === 'Could not log in'),
    ).not.toBeNull();
  });
});
