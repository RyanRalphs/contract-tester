import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('Container and main header are present', () => {
  render(<App />);
  const linkElement = screen.getByText(/API Form/i);
  expect(linkElement).toBeInTheDocument();
});
