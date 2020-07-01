import { atom, selector } from 'recoil';

import axios from 'axios';

export const count = atom({
  key: 'count',
  default: 0,
});

export const getDatabaseHandler = selector({
  key: 'textState',
  get: async ({ get }) => {
    try {
      const context = await axios.get('/api/v1/users');
      return context.data.data;
    } catch (error) {
      console.log(error);
      return [];
    }
  },
});
