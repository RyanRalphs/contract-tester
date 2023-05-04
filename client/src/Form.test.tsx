import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

test('The API URL input is present', () => {
  render(<App />);
  const linkElement = screen.getByText(/API URL/i);
  expect(linkElement).toBeInTheDocument();
});

test('The Expected Outcome input is present', () => {
  render(<App />);
  const linkElement = screen.getByText(/Expected Outcome/i);
  expect(linkElement).toBeInTheDocument();
});

test('The Method input is present', () => {
  render(<App />);
  const linkElement = screen.queryAllByText(/Method/i)[0];
  expect(linkElement).toBeInTheDocument();
});