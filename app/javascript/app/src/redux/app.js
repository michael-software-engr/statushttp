const DUMMY_ACTION = '/dummy/action';

export default function reducer(state = {}, action) {
  const { type, payload } = action;

  switch (type) {
    case DUMMY_ACTION: {
      return { ...state, dummy: payload };
    }

    default:
      return state;
  }
}
