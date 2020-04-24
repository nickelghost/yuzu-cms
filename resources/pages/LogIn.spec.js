import { render } from '@testing-library/svelte';

import LogIn from './LogIn.svelte';

describe('Log in page', () => {
  it('has a Yuzu heading', () => {
    const { getByText } = render(LogIn);
    const heading = getByText('Yuzu');
    expect(heading).toBeInstanceOf(HTMLHeadingElement);
  });
  it('has user and password inputs', () => {
    const { getByLabelText } = render(LogIn);
    expect(getByLabelText('User')).toBeInstanceOf(HTMLInputElement);
    expect(getByLabelText('Password')).toBeInstanceOf(HTMLInputElement);
  });
  it('has a log in button', () => {
    const { getByText } = render(LogIn);
    expect(getByText('Log in')).toBeInstanceOf(HTMLButtonElement);
  });
});
