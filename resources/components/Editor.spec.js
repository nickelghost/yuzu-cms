import { render, fireEvent } from '@testing-library/svelte';

import Editor from './Editor.svelte';

describe('Editor', () => {
  it('should render a textarea', () => {
    const { container } = render(Editor);
    expect(container.querySelector('textarea')).not.toBeNull();
  });
  it('should output a preview', async () => {
    const { container } = render(Editor);
    const textarea = container.querySelector('textarea');
    await fireEvent.input(textarea, { target: { value: '# Random heading' } });
    expect(container.querySelector('h1')).toHaveTextContent('Random heading');
  });
});
