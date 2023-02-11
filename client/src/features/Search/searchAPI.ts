// A mock function to mimic making an async request for data
export function fetchSearch(input = "") {
  return new Promise<{ data: Array<any> }>((resolve) =>
    setTimeout(() => resolve({ data: [input] }), 500)
  );
}
