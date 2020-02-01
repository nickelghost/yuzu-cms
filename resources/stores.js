import { writable } from 'svelte/store';

export const jwt = writable(
  localStorage.getItem('jwt') || '',
);
